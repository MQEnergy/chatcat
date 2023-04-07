<template>
  <a-modal v-model:visible="props.visible"
           :title="$t('settings.prompt')"
           :cancel-text="$t('settings.prompt.cate.cancel')"
           :ok-text="$t('settings.prompt.cate.submit')"
           @cancel="handleCancel"
           @before-ok="handleBeforeOk"
           @ok="handleOk">
    <a-form ref="formRef" :model="form" layout="vertical">
      <a-form-item field="name" :label="$t('settings.prompt.cate.name')"
                   :rules="[
                       {required:true, message:$t('settings.prompt.cate.name.errTip1')},
                   ]"
                   :validate-trigger="['change','input']">
        <a-input v-model="form.name" :placeholder="$t('settings.prompt.cate.name.placeholder')"/>
      </a-form-item>
      <a-form-item field="prompt" :label="$t('settings.prompt.cate.prompt')"
                   :rules="[
                       {required:true, message:$t('settings.prompt.cate.prompt.errTip1')},
                   ]"
                   :validate-trigger="['change','input']">
        <a-textarea v-model="form.desc" :placeholder="$t('settings.prompt.cate.prompt.placeholder')" allow-clear/>
      </a-form-item>
      <a-form-item field="desc" :label="$t('settings.prompt.cate.desc')">
        <a-textarea v-model="form.desc" :placeholder="$t('settings.prompt.cate.desc.placeholder')" allow-clear/>
      </a-form-item>
      <a-form-item field="enprompt" :label="$t('settings.prompt.cate.enprompt')">
        <a-textarea v-model="form.desc" :placeholder="$t('settings.prompt.cate.enprompt.placeholder')" allow-clear/>
      </a-form-item>
    </a-form>
  </a-modal>
</template>

<script setup>
import {reactive, ref} from "vue";

const props = defineProps({
  visible: {
    type: Boolean,
    default: false
  }
});
const form = reactive({
  name: '',
  desc: '',
  prompt: '',
  enprompt: '',
});
const emits = defineEmits(["cancel", "ok"])
const formRef = ref(null)
const handleCancel = () => {
  formRef.value.resetFields();
  emits("cancel", false)
}
const handleBeforeOk = (done) => {
  if (formRef.value) {
    formRef.value.validate()
        .then(res => done(!res ? true : false))
        .catch(error => done(false));
  } else {
    done(true);
  }
}
const handleOk = () => {
  emits("ok", form)
}
</script>

<style scoped>

</style>