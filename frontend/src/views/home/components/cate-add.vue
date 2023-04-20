<template>
  <a-modal :width="530" v-model:visible="props.visible"
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
                       {maxLength:10, message:$t('home.menuCate.settings.addCate.errTip2')}
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
      <a-space class="color-space-div" wrap style="cursor: pointer; margin-bottom: 10px;">
        <a-tag v-for="(item, index) of colorList" :key="index" :checked="item.checked"
               @click="handleColorCheck(item, index)" :color="item.value"
               :style="{width: '30px', position: 'relative'}">
          <div v-if="item.checked" class="selected">
            <icon-check :size="14"/>
          </div>
        </a-tag>
      </a-space>
      <a-form-item field="desc" :label="$t('common.desc')">
        <a-textarea v-model="form.desc" :placeholder="$t('common.desc.placeholder')" allow-clear/>
      </a-form-item>
    </a-form>
  </a-modal>
</template>

<script setup>
import {reactive, ref, watch} from "vue";

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
  name: '',
  desc: '',
  color: '#3370ff',
});
const emits = defineEmits(["cancel", "ok"])
const formRef = ref(null)
watch(() => props.formData, () => {
  form.id = props.formData.id;
  form.name = props.formData.name;
  form.desc = props.formData.desc;
  form.color = props.formData.color;
  console.log("form", form)
})
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
let colorList = reactive([
  {value: '#f9801c', checked: false},
  {value: '#3271fd', checked: false},
  {value: '#1fcabc', checked: false},
  {value: '#53b851', checked: false},
  {value: '#f6a646', checked: false},
  {value: '#f96c16', checked: false},
  {value: '#0b73f5', checked: false},
  {value: '#8189f5', checked: false},
  {value: '#26818a', checked: false},
  {value: '#1e80ff', checked: false},
  {value: '#2848f9', checked: false},
  {value: '#c73d3d', checked: false},
  {value: '#a932f7', checked: false},
]);
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
  console.log(form)
  emits("ok", form)
}
const handleCheckTag = (row, idx) => {
  tagCheckIdx.value = idx;
  form.name = row.name;
}
const handleColorCheck = (row, idx) => {
  colorList.forEach((item) => {
    item.checked = false;
  })
  form.color = row.value;
  colorList[idx].checked = !colorList[idx].checked;
}

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