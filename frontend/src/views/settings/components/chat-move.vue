<template>
  <a-modal :width="400" v-model:visible="props.visible"
           :title="$t('settings.chat.catmove')"
           :cancel-text="$t('common.cancel')"
           :ok-text="$t('common.submit')"
           @cancel="handleCancel"
           @before-ok="handleBeforeOk"
           @ok="handleOk">
    <a-form ref="formRef" :model="form" layout="vertical">
      <a-form-item field="cate_id" :label="$t('settings.chat.label.catename')"
                   :rules="[
                       {required:true, message:$t('home.menuCate.settings.addCate.errTip1')},
                   ]"
                   :validate-trigger="['change','input']">
        <a-select v-model="form.cate_id" :style="{width:'100%'}" placeholder="Please select ...">
          <a-option v-for="(item, index) in cateList" :key="index" :value="item.id" :label="item.name"/>
        </a-select>
      </a-form-item>
    </a-form>
  </a-modal>
</template>

<script setup>
import {onMounted, reactive, ref, watch} from "vue";
import {GetAllCateList} from "../../../../wailsjs/go/chat/Service.js";
import {useI18n} from "vue-i18n";

const {t} = useI18n();
const props = defineProps({
  visible: {
    type: Boolean,
    default: false
  },
  formData: {
    type: Object,
    default: null
  }
});
const form = reactive({
  id: null,
  cate_id: null,
});
const cateList = reactive([]);
const emits = defineEmits(["cancel", "ok"])
const formRef = ref(null)
watch(() => props.formData, () => {
  if (props.formData) {
    form.id = props.formData.id;
    form.cate_id = props.formData.cate_id;
  } else {
    formRef.value.resetFields();
  }
})
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
const initCateList = () => {
  GetAllCateList().then(res => {
    if (res.code !== 0) {
      return;
    }
    res.data.push({
      id: 0,
      name: t('common.nocate')
    })
    cateList.splice(0, cateList.length);
    cateList.push(...res.data);
  })
}
onMounted(() => {
  initCateList();
})
</script>

<style scoped>
.color-space-div {
  position: relative;
}

.color-space-div .selected {
  position: absolute;
  top: 0;
  left: 0;
  height: 100%;
  width: 100%;
  text-align: center;
}
</style>