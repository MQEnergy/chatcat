<template>
  <a-space class="prompt-container" direction="horizontal" style="width: 100%;">
    <a-mention size="large" v-model="promptValue" :data="promptPrefix"
               prefix="/" placeholder="输入 / 或直接输入对话关键词" @select="handleSelect" @keydown.enter="handleSend" allow-clear/>
    <a-button size="large" type="primary" @click="handleSend">send</a-button>
  </a-space>
</template>
<script setup>
import {reactive, ref, watch} from "vue";

const props = defineProps({
  value: {
    type: String,
    default: ""
  }
})
const promptValue = ref(props.value);
const promptPrefix = reactive(['问答:', '翻译:', '改写:', '润色:', '总结:', '分析:', '解释:', '解释代码:', '检查代码:']);

watch(() => props.value, () => {
  promptValue.value = props.value;
})
const emits = defineEmits(['ok'])
const handleSelect = (value) => {
}
const handleSend = () => {
  if (promptValue.value === "/") {
    return
  }
  const filter = promptPrefix.filter((item) => {
    return promptValue.value == item
  });
  if (filter.length > 0) {
    return
  }
  emits('ok', promptValue.value)
  promptValue.value = "";
}
</script>

<style scoped>
.prompt-container :deep(.arco-input-wrapper .arco-input) {
  height: 30px !important;
}

.prompt-container :deep(.arco-btn) {
  height: 44px !important;
}
</style>