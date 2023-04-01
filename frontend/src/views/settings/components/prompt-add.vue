<template>
  <a-modal v-model:visible="props.visible" :title="$t('settings.prompt.cate')"
           :cancel-text="$t('settings.prompt.cate.cancel')"
           :ok-text="$t('settings.prompt.cate.submit')"
           @cancel="handleCancel"
           @before-ok="handleBeforeOk">
    <a-form :model="form" layout="vertical">
      <a-form-item field="name" :label="$t('settings.prompt.cate.name')">
        <a-input v-model="form.name" :placeholder="$t('settings.prompt.cate.name.placeholder')"/>
      </a-form-item>
      <a-form-item field="desc" :label="$t('settings.prompt.cate.desc')">
        <a-textarea v-model="form.desc" :placeholder="$t('settings.prompt.cate.desc.placeholder')" allow-clear/>
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