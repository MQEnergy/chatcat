<template>
  <a-space class="prompt-container" direction="horizontal" style="width: 100%;">
    <a-mention size="large" v-model="promptValue" :data="promptPrefix"
               prefix="/" placeholder="输入 / 获取模板 或直接输入问题" @keydown.enter="handleSend" allow-clear/>
    <a-button size="large" type="primary" @click="handleSend" :loading="sendLoading">send</a-button>
  </a-space>
</template>
<script setup>
import {reactive, ref, watch} from "vue";

const props = defineProps({
  value: {
    type: String,
    default: ""
  },
  loading: {
    type: Boolean,
    default: false
  }
})
const promptValue = ref(props.value);
const promptPrefix = reactive(['问答:', '翻译:', '改写:', '润色:', '总结:', '分析:', '解释:', '解释代码:', '检查代码:']);
const sendLoading = ref(props.loading);

watch(() => props.value, () => {
  promptValue.value = props.value;
})
watch(() => props.loading, () => {
  sendLoading.value = props.loading;
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
</script>

<style scoped>
.prompt-container :deep(.arco-input-wrapper .arco-input) {
  height: 30px !important;
}

.prompt-container :deep(.arco-btn) {
  height: 44px !important;
}
</style>