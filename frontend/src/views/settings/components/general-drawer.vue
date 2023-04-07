<template>
  <a-drawer :width="500"
            :visible="props.visible"
            @ok="handleOk"
            @cancel="handleCancel"
            :hide-cancel="true"
            unmountOnClose>
    <template #title>
      {{ $t('settings.general.alertTips.drawer.title') }}
    </template>
    <a-typography-title :heading="3" style="margin-top: 0px;">
      {{ $t('settings.general.alertTips.drawer.step1.h1') }}
    </a-typography-title>
    <a-typography>
      <a-typography-title :heading="4">
        {{ $t('settings.general.alertTips.drawer.step1.h2') }}
      </a-typography-title>
      <a-typography-paragraph>
        <ol>
          <li>
            <a-link icon>{{ $t('settings.general.alertTips.drawer.step1.li1-1') }}</a-link>
            {{ $t('settings.general.alertTips.drawer.step1.li1-2') }}
            <a-link icon>{{ $t('settings.general.alertTips.drawer.step1.li1-3') }}</a-link>
            {{ $t('settings.general.alertTips.drawer.step1.li1-4') }}
          </li>
          <li>
            <a-link href="https://platform.openai.com/account/api-keys" icon>
              {{ $t('settings.general.alertTips.drawer.step1.li2') }}
            </a-link>
          </li>
          <li>{{ $t('settings.general.alertTips.drawer.step1.li3') }}</li>
        </ol>
      </a-typography-paragraph>
      <a-typography-paragraph>
        <a-space direction="vertical">
          <a-image
              width="100%"
              :src="Img1"
              show-loader
          />
          <a-image
              width="100%"
              :src="Img2"
              show-loader
          />
        </a-space>
      </a-typography-paragraph>
      <a-typography-title :heading="4">
        {{ $t('settings.general.alertTips.drawer.step2.h2') }}
      </a-typography-title>
      <a-space direction="vertical">
        <a-typography-text>
          {{ $t('settings.general.alertTips.drawer.step2.li1') }}
        </a-typography-text>
        <a-typography-text>
          <a-input-password :style="{width:'390px'}" placeholder="sk-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" allow-clear>
            <template #prepend>
              API key
            </template>
          </a-input-password>
          <a-button type="primary">保存</a-button>
        </a-typography-text>
      </a-space>
      <a-typography-title :heading="4">
        <a-space>
          {{ $t('settings.general.alertTips.drawer.step3.h2-1') }}
          <a-link style="font-size: 24px;" href="https://platform.openai.com/account/usage" icon>
            {{ $t('settings.general.alertTips.drawer.step3.h2-2') }}
          </a-link>
          {{ $t('settings.general.alertTips.drawer.step3.h2-3') }}
        </a-space>
      </a-typography-title>
    </a-typography>
  </a-drawer>
</template>

<script setup>
import Img1 from '@assets/images/img.png';
import Img2 from '@assets/images/img_1.png';

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