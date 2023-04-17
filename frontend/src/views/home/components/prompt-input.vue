<template>
  <a-space class="prompt-container" direction="horizontal" style="width: 100%;">
    <a-mention size="large" v-model="promptValue" :data="promptPrefix"
               prefix="/" placeholder="输入 / 获取模板 或直接输入问题" @keydown.enter="handleSend" allow-clear/>
    <a-button class="prompt-btn" style="width: 60px;" size="large" type="primary" @click="handleSend"
              :loading="sendLoading">
      <template #icon>
        <icon-send size="18"/>
      </template>
    </a-button>
    <!-- 停止对话stream -->
    <div class="backoff-container" v-if="checkOffFlag" style="position: absolute; bottom: 50px; left: 0px; ">
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
const promptPrefix = reactive(['问答:', '翻译:', '改写:', '润色:', '总结:', '分析:', '解释:', '解释代码:', '检查代码:']);
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
  if (promptValue.value == "" || promptValue.value === "/" || sendLoading.value == true) {
    return
  }
  const filter = promptPrefix.filter((item) => {
    return promptValue.value == "/" + item && promptValue.value.length == item.length + 1
  });
  if (filter.length > 0) {
    return
  }
  sendLoading.value = true;
  emits('ok', promptValue.value, sendLoading.value)
  promptValue.value = "";
}
const handleBreakOffChat = () => {
  BreakOffChatStream()
}
</script>

<style scoped>
.prompt-container {
  position: relative;
}

.prompt-container :deep(.arco-input-wrapper .arco-input) {
  height: 30px !important;
}

.prompt-container :deep(.prompt-btn) {
  height: 44px !important;
}
</style>