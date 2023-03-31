<template>
  <a-modal v-model:visible="props.visible" title="添加分类" @cancel="handleCancel" @before-ok="handleBeforeOk">
    <a-form :model="form" layout="vertical" >
      <a-form-item field="name" label="名称">
        <a-input v-model="form.name" placeholder="请输入分类名称 如：笔记"/>
      </a-form-item>
      <a-form-item field="desc" label="备注">
        <a-textarea v-model="form.desc" placeholder="请填写备注说明" allow-clear/>
      </a-form-item>
    </a-form>
  </a-modal>
</template>

<script setup>
import {reactive} from "vue";

const props = defineProps({
  visible: {
    type: Boolean,
    default: false
  }
});
const form = reactive({
  name: '',
  desc: ''
});
const emits = defineEmits(["cancel"])
const handleCancel = () => {
  emits("cancel", false)
}
const handleBeforeOk = (done) => {
  console.log(form)
  window.setTimeout(() => {
    done()
  }, 3000)
}

</script>

<style scoped>

</style>