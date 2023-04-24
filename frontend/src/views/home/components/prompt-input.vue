<template>
  <a-space class="prompt-container" direction="horizontal" style="width: 100%;">
    <div class="prompt-input-div">
      <div class="prompt-extension">
        <a-select v-model="curSysPrompt" :style="{width:'160px'}" placeholder="Please select ..."
                  @change="handleSysPromptChange">
          <a-option v-for="(item, index) in systemPromptList" :key="index" :value="item" :label="item.label"/>
        </a-select>
        {{ curSysPrompt.value }}
        <a-select v-if="curSysPrompt.type === 2" v-model="curSysPrompt.language"
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
  label: "自定义内容:",
  value: "",
  type: 0,
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
  let systemPromptValue = curSysPrompt.value.value;
  if (curSysPrompt.value.type === 2) {
    systemPromptValue += curSysPrompt.value.language;
  }
  let sendPromptList = [
    {
      role: 'user',
      content: promptValue.value.trim()
    }
  ];
  if (curSysPrompt.value.type !== 0) {
    sendPromptList.unshift({
      role: 'system',
      content: systemPromptValue,
    });
  }
  sendLoading.value = true;
  emits('ok', sendPromptList, curSysPrompt.value.type, sendLoading.value)
  promptValue.value = "";
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
  curSysPrompt.value.language = '中文(简体)';
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