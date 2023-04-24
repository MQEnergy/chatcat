<template>
  <div class="page-header" style="--wails-draggable: drag;">
    <a-page-header
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
</script>

<style scoped>
.page-header {
  background: var(--color-bg-2);
  --wails-draggable: drag;
  top: 0;
  position: absolute;
  width: 100%;
  border-bottom: 1px solid var(--color-neutral-3);
  z-index: 9
}

.page-header :deep(.arco-page-header) {
  padding: 10px 0 !important;
}
</style>