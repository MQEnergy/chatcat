<template>
  <a-card :title="$t('settings.contact')" :bordered="false" :header-style="{borderColor: 'var(--color-fill-2)'}">
    <template #extra>
      <a-link @click="handleStarRedirect" target="_blank">
        <icon-star-fill/>
        {{ $t('settings.contact.star') }}
      </a-link>
    </template>
    <a-form ref="formRef" :size="form.size" :model="form" layout="vertical" :style="{width:'440px'}"
            @submit="handleSubmit">
      <a-form-item field="feedbackType" :label="$t('settings.contact.feedbackType')">
        <a-select v-model="form.feedbackType" :placeholder="$t('settings.contact.feedbackType.placeholder')">
          <a-option value="1">{{ $t('settings.contact.feedbackType.option1') }}</a-option>
          <a-option value="2">{{ $t('settings.contact.feedbackType.option2') }}</a-option>
        </a-select>
      </a-form-item>
      <a-form-item field="title"
                   :rules="[
                       {required:true,message:$t('settings.contact.title.errorTips1')},
                       {minLength:5,message:$t('settings.contact.title.errorTips2')}
                   ]"
                   :validate-trigger="['change','input']">
        <template #label>
          {{ $t('settings.contact.title') }}
          <div :style="{color: 'var(--color-neutral-6)', fontSize: '12px'}">{{ $t('settings.contact.title.tips') }}</div>
        </template>
        <a-input v-model="form.title" :placeholder="$t('settings.contact.title.placeholder')"/>
      </a-form-item>
      <a-form-item field="desc" :label="$t('settings.contact.desc')"
                   :rules="[
                       {required:true,message:$t('settings.contact.desc.errorTips1')},
                       {minLength:5,message:$t('settings.contact.desc.errorTips2')}
                   ]"
                   :validate-trigger="['change','input']">
        <template #label>
          {{ $t('settings.contact.desc') }}
          <div :style="{color: 'var(--color-neutral-6)', fontSize: '12px'}">{{ $t('settings.contact.desc.tips') }}</div>
        </template>
        <a-textarea style="height: 200px;" v-model="form.desc" :placeholder="$t('settings.contact.desc.placeholder')" :max-length="200"
                    show-word-limit allow-clear/>
      </a-form-item>
      <a-form-item>
        <a-space>
          <a-button type="primary" html-type="submit">{{ $t('settings.contact.addBtn') }}</a-button>
          <a-button @click="$refs.formRef.resetFields()">{{ $t('settings.contact.cancelBtn') }}</a-button>
        </a-space>
      </a-form-item>
    </a-form>
  </a-card>
</template>

<script>
import {reactive} from 'vue';
import {BrowserOpenURL} from "../../../../wailsjs/runtime/runtime.js";
import config from "@config/config.js";
import {GetFeedBackUrl} from "../../../../wailsjs/go/setting/Service.js";

export default {
  setup() {
    const handleSubmit = ({values, errors}) => {
      console.log('values:', values, '\nerrors:', errors)
      if (!errors) {
        GetFeedBackUrl({
          title: values.title,
          body: values.desc,
          issue_type: parseInt(values.feedbackType)
        }).then(res => {
          BrowserOpenURL(res.data)
          console.log('res:', res.data)
        })
      }
    }
    const handleStarRedirect = () => {
      BrowserOpenURL(config.githubUrl)
    }
    const form = reactive({
      size: 'medium',
      title: '',
      feedbackType: '1',
      desc: '',
    });
    return {
      form,
      handleSubmit,
      handleStarRedirect
    }
  },
}
</script>

<style scoped>
</style>