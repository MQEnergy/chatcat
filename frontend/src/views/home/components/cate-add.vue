<template>
  <a-modal v-model:visible="props.visible"
           :title="$t('home.menuCate.settings.addCate')"
           :cancel-text="$t('common.reset')"
           :ok-text="$t('common.submit')"
           @cancel="handleCancel"
           @before-ok="handleBeforeOk"
           @ok="handleOk">
    <a-form ref="formRef" :model="form" layout="vertical">
      <a-form-item field="name" :label="$t('common.name')"
                   :rules="[
                       {required:true, message:$t('home.menuCate.settings.addCate.errTip1')},
                       {maxLength:5, message:$t('home.menuCate.settings.addCate.errTip2')}
                   ]"
                   :validate-trigger="['change','input']">
        <a-input v-model="form.name" :placeholder="$t('common.name.placeholder')" allow-clear/>
      </a-form-item>
      <a-space style="cursor: pointer; margin-bottom: 10px;">
        <a-tag checkable v-for="(item, index) in tagList" :key="index"
               color="arcoblue" :checked="tagCheckIdx === index"
               @check="handleCheckTag(item, index)">{{ item.name }}
        </a-tag>
      </a-space>
      <a-form-item field="desc" :label="$t('common.desc')">
        <a-textarea v-model="form.desc" :placeholder="$t('common.desc.placeholder')" allow-clear/>
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
  desc: ''
});
const emits = defineEmits(["cancel", "ok"])
const formRef = ref(null)
const tagCheckIdx = ref(null)
const tagList = reactive([
  {
    id: 1,
    name: '技术',
  }, {
    id: 2,
    name: '作业',
  }, {
    id: 3,
    name: '汇报',
  }, {
    id: 4,
    name: '文章',
  }, {
    id: 5,
    name: '论文',
  },
])
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
const handleCheckTag = (row, idx) => {
  tagCheckIdx.value = idx;
  form.name = row.name;
}
</script>

<style scoped>

</style>