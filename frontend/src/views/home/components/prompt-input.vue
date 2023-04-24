<template>
  <a-space class="prompt-container" direction="horizontal" style="width: 100%;">
    <div class="prompt-input-div">
      <div class="prompt-extension">
        <a-select v-model="curSysPrompt.id" :style="{width:'160px'}" placeholder="Please select ..."
                  @change="handleSysPromptChange">
          <a-option v-for="(item, index) in systemPromptList" :key="index" :value="item.id" :label="item.label"/>
        </a-select>
        <span v-if="curSysPrompt.type === 2 && curSysPrompt.id !== 1">{{ curSysPrompt.desc }}</span>
        <a-input v-model="curSysPrompt.desc" size="mini" :style="{width:'320px', marginLeft: '10px'}"
                 v-if="curSysPrompt.type === 1"
                 placeholder="例如：我希望你能充当私人教练。"/>

        <a-select v-if="curSysPrompt.type === 2 && curSysPrompt.id === 3" v-model="curSysPrompt.language"
                  :style="{width:'120px', marginLeft: '10px'}"
                  placeholder="Please select ...">
          <a-option v-for="(item, index) in curSysPrompt.extra" :key="index" :value="item.label" :label="item.label"/>
        </a-select>
      </div>
      <a-textarea v-model="promptValue" class="prompt-textarea" :placeholder="$t('common.prompt.input.placeholder')"
                  @keydown="handleKeyDownSend"/>
      <a-button class="prompt-btn" size="large" type="primary" @click="handleSend"
                :loading="sendLoading">
        <template #icon>
          <icon-send size="18"/>
        </template>
      </a-button>
    </div>

    <!-- 停止对话stream -->
    <div class="backoff-container" v-if="checkOffFlag" style="position: absolute; bottom: 175px; left: 10px; ">
      <a-button size="medium" @click="handleBreakOffChat">
        <template #icon>
          <icon-record-stop/>
        </template>
        {{ $t('common.breakoff') }}
      </a-button>
    </div>
  </a-space>
</template>
<script setup>
import {reactive, ref, watch} from "vue";
import {BreakOffChatStream} from "../../../../wailsjs/go/chat/Service.js";
import SysInputEnums from "../../../config/sys-input.js";

const props = defineProps({
  value: {
    type: String,
    default: ""
  },
  loading: {
    type: Boolean,
    default: false
  },
  checkoff: {
    type: Boolean,
    default: false
  }
})
const promptValue = ref(props.value);
const systemPromptList = reactive(SysInputEnums)
const curSysPrompt = ref({
  id: 2,
  label: "问答:",
  desc: "请输入问答内容:",
  type: 2,
  is_sys: 2,
  extra: [],
  language: '',
});
const sendLoading = ref(props.loading);
const checkOffFlag = ref(props.checkoff);

watch(() => props.value, () => {
  promptValue.value = props.value;
})
watch(() => props.loading, () => {
  sendLoading.value = props.loading;
})
watch(() => props.checkoff, () => {
  checkOffFlag.value = props.checkoff;
})
const emits = defineEmits(['ok'])
const handleSend = () => {
  if (promptValue.value == "" || sendLoading.value == true) {
    return;
  }
  const sendPromptList = assemblePrompt();
  sendLoading.value = true;
  emits('ok', sendPromptList, curSysPrompt.value, sendLoading.value)
  promptValue.value = "";
}
const assemblePrompt = () => {
  let promptList = [];
  switch (curSysPrompt.value.type) {
    case 1: // Todo
      promptList = [{
        role: 'system',
        prefix: '',
        content: curSysPrompt.value.desc
      }, {
        role: 'user',
        prefix: '',
        content: promptValue.value.trim()
      }];
      break;
    case 2:
      let systemPromptValue = curSysPrompt.value.desc;
      if (curSysPrompt.value.type === 2 && curSysPrompt.value.id === 3) {
        systemPromptValue += curSysPrompt.value.language + ":";
      }
      if (curSysPrompt.value.type === 2 && curSysPrompt.value.id === 2) {
        systemPromptValue = "";
      }
      promptList = [{
        role: 'user',
        prefix: systemPromptValue,
        content: promptValue.value.trim()
      }]
      break;
  }
  return promptList
}
const handleKeyDownSend = (event) => {
  if (event.ctrlKey && event.keyCode === 13) {
    handleSend()
  }
}
const handleBreakOffChat = () => {
  BreakOffChatStream()
}
const handleSysPromptChange = (e) => {
  console.log("e", e)
  const promptInfo = systemPromptList.filter((item) => {
    return e === item.id
  });
  if (promptInfo.length > 0) {
    console.log(promptInfo)
    curSysPrompt.value = promptInfo[0];
  }
  if (curSysPrompt.value.type === 2 && curSysPrompt.value.id === 3) {
    curSysPrompt.value.language = '中文(简体)';
  }
}
</script>

<style scoped>
.prompt-container {
  position: relative;
}

.prompt-container .prompt-input-div {
  height: 145px;
  border-top: 1px solid var(--color-neutral-3);
  width: 100%;
  padding: 10px;
  background: var(--color-bg-2);
  position: relative;
}

.prompt-container .prompt-input-div .prompt-textarea {
  background: var(--color-bg-3);
  height: 100px;
  border: 1px solid var(--color-neutral-3);
  margin-top: 10px;
  border-radius: 6px;
}

.prompt-container .prompt-extension {
}

.prompt-container :deep(.arco-input-wrapper .arco-input) {
  height: 30px !important;
}

.prompt-container :deep(.prompt-btn) {
  border-radius: 6px;
  width: 45px;
  position: absolute;
  right: 20px;
  bottom: 20px;
}
</style>