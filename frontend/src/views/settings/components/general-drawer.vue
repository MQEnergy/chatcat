<template>
  <a-drawer :width="500"
            :visible="props.visible"
            @ok="handleOk"
            @cancel="handleCancel"
            :hide-cancel="true"
            unmountOnClose>
    <template #title>
      歡迎使用ChatGPT Sidebar！ 這是您開始的旅程：
    </template>
    <a-typography-title :heading="3" style="margin-top: 0px;">
      如何
      <a-link style="font-size: 28px;" status="danger">免費</a-link>
      獲取 API 密鑰和 18 美元的 OpenAI 積分。
    </a-typography-title>
    <a-typography>
      <a-typography-title :heading="4">
        第 1 步：獲取 API 密鑰
      </a-typography-title>
      <a-typography-paragraph>
        <ol>
          <li>
            <a-space>
              <a-link icon>登錄您的 OpenAI 帳戶</a-link>
              如果您沒有帳戶，請在
              <a-link icon>此處</a-link>
              註冊
            </a-space>
          </li>
          <li>
            <a-link href="https://platform.openai.com/account/api-keys" icon>訪問 OpenAI API 密鑰頁面</a-link>
          </li>
          <li>創建一個新的密鑰並按如下方式複制它</li>
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
        第 2 步：將 API 密鑰粘貼到下面的框中
      </a-typography-title>
      <a-space direction="vertical">
        <a-typography-text>
          不用擔心。 您的 API 密鑰僅在 ChatGPT 擴展中本地使用。 使用 API 比直接使用 ChatGPT 更穩定。
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
          第 3 步：在
          <a-link style="font-size: 24px;" href="https://platform.openai.com/account/usage" icon>此處</a-link>
          查看您18美元的積分
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