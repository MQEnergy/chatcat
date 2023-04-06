<template>
  <div style="--wails-draggable: drag;">
    <a-page-header
        :style="{
          background: 'var(--color-bg-2)',
          '--wails-draggable': 'drag', 'top': 0, 'position': 'absolute',
          'width': '100%', 'height': '30px', 'border-bottom': '1px solid var(--color-neutral-3)',
          'z-index': '9'
      }"
        :title="props.info.cateName"
        :subtitle="props.info.chatName"
        :show-back="false"
    >
      <template #extra>
        <a-radio-group type="button" default-value="large">
          <a-radio value="mini">{{ props.info.modelName }}</a-radio>
          <a-radio value="small">{{ props.info.msgNum }} messages</a-radio>
          <a-radio value="large">{{ props.info.tokenNum }} tokens</a-radio>
        </a-radio-group>
      </template>
    </a-page-header>
  </div>
</template>

<script setup>
const props = defineProps({
  info: {
    type: Object,
    default: {
      cateName: "",
      chatName: "",
      modelName: "",
      msgNum: "",
      tokenNum: ""
    }
  }
})
// let headerInfo = reactive({
//   cateName: props.info.cateName,
//   chatName: props.info.chatName,
//   modelName: props.info.modelName,
//   msgNum: props.info.msgNum,
//   tokenNum: props.info.tokenNum
// })
// watch(() => props.info, () => {
//   headerInfo = props.info
// })
</script>

<style scoped>

</style>