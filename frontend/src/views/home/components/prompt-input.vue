<template>
  <a-space direction="horizontal" style="width: 100%;">
    <a-mention size="large" v-model="promptValue" :data="['翻译', '润色', '总结', '分析', '解释', '解释代码', '检查代码']"
               prefix="/" placeholder="输入 / 或直接输入对话关键词"/>
    <a-button size="large" type="primary" @click="handleSend">send</a-button>
  </a-space>
</template>
<script setup>
import {ref, watch} from "vue";

const props = defineProps({
  value: {
    type: String,
    default: ""
  }
})
const promptValue = ref(props.value)
watch(() => props.value, () => {
  promptValue.value = props.value;
})
const emits = defineEmits(['ok'])
const handleSend = () => {
  emits('ok', promptValue.value)
}
</script>

<style scoped>

</style>