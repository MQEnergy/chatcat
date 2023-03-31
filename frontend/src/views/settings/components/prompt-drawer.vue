<template>
  <a-drawer :width="500"
            :visible="props.visible"
            @ok="handleOk"
            @cancel="handleCancel"
            :hide-cancel="true"
            unmountOnClose>
    <template #title>
      prompt提示词列表
    </template>

  </a-drawer>
</template>

<script setup>
const props = defineProps({
  visible: {
    type: Boolean,
    default: false
  }
});

const emits = defineEmits(['cancel']);
const handleOk = () => {
  emits('cancel', false);
}
const handleCancel = () => {
  emits('cancel', false);
}
</script>

<style scoped>

</style>