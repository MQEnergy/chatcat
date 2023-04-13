<template>
  <div class="chat-container" ref="chatListRef" style="padding: 20px;">
    <a-space direction="vertical" :size="30">
      <a-space class="chat-space-container" align="start" :size="10" v-for="(item, index) in chatList" :key="index">
        <!-- 头像 -->
        <a-avatar
            :style="{backgroundColor: item.from === 2 ? '#fff' : '#165DFF', overflow: 'hidden'}"
            :size="32">
          <img v-if="item.from === 2"
               alt="avatar"
               :src="ChatGPTLogo"
          />
          <template v-else>
            <icon-robot/>
          </template>
        </a-avatar>
        <!-- 聊天框 -->
        <a-space direction="vertical">
          <a-card hoverable class="chat-card-item" :style="{
          borderColor: item.from === 2 ? '' : '#c7d7fa',
          backgroundColor: item.from === 2 ? '#fff': '#e8f3ff'
        }">
            <!--            <a-typography-paragraph copyable>-->
            <div class="chat-div" v-html="markedContent(item.content)"
                 :style="{color: item.from === 2 ? '#000': '#000'}">
            </div>
            <!--            </a-typography-paragraph>-->

            <!--            <a-button v-if="item.from === 2" class="copy-btn" type="primary" data-clipboard-text="Just because you can doesn't mean you should — clipboard.js">-->
            <!--              <icon-copy/>-->
            <!--            </a-button>-->
            <template #actions>
              <IconShareInternal/>
              <icon-copy @click="handleCopy($event)"/>
              <a-dropdown @select="handleSelect" position="bl">
                <IconMore/>
                <template #content>
                  <a-doption>{{ $t('common.save') }}</a-doption>
                  <a-doption>{{ $t('common.copy') }}</a-doption>
                </template>
              </a-dropdown>
            </template>
          </a-card>
        </a-space>
      </a-space>
    </a-space>
  </div>
</template>

<script setup>
import ChatGPTLogo from '@assets/images/chatgpt_black_logo.svg';
import {onMounted, reactive, ref, watch, watchEffect} from "vue";
import marked from "@plugins/markdown/marked.js";
import "highlight.js/styles/default.css";

const props = defineProps({
  list: {
    type: Array,
    default: []
  },
});
let chatList = reactive(props.list)

watch(() => props.list, () => {
  chatList = props.list
})

const markedContent = (contents) => {
  return marked.parse(contents);
}
// 初始化滚动条位置
const chatListRef = ref(null);
onMounted(() => {
  chatListRef.value = document.querySelector('.chat-container');
});
// 监听页面高度变化
watchEffect(() => {
  // 确保 chatListRef.value 被正确地渲染
  if (chatListRef.value) {
    // const {scrollHeight} = chatListRef.value;
    // window.scrollTo(0, scrollHeight);
    chatList.scrollTop = chatList.scrollHeight;
  }
});
const handleCopy = (e) => {
  const textContent = e.target.parentNode.parentNode.parentNode.parentNode.firstElementChild.querySelector('.hljs').textContent;

}
const handleSelect = (e) => {
  console.log(e)
}
</script>

<style scoped>
.chat-container :deep(.arco-card-size-medium) {
  /*font-size: 16px;*/
}

.chat-container :deep(.arco-card-size-medium .arco-card-body) {
  padding: 10px;
}

.chat-div {
  overflow: scroll;
}

.chat-space-container :deep(.arco-card-actions) {
  margin-top: 0px !important;
}

.chat-card-item {
  width: 510px;
  border-top-right-radius: 12px;
  border-bottom-right-radius: 12px;
  border-bottom-left-radius: 12px;
}

.chat-space-container {
  width: 100%;
//background: #e5e6ec;
}

.chat-container :deep(.arco-card-body p) {
  margin: 0px !important;
}

.chat-container :deep(.chat-div p code) {
  background: #f7f7f7;
  color: #800;
}
</style>