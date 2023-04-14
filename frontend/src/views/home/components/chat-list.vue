<template>
  <div class="chat-container" style="padding: 20px;">
    <a-space direction="vertical" :size="20">
      <a-space class="chat-space-container" align="start" :size="10">
        <a-avatar
            :style="{backgroundColor: '#fff', overflow: 'hidden'}"
            :size="32">
          <img alt="avatar" style="height: 30px; width: 30px" :src="ChatGPTLogo"/>
        </a-avatar>
        <!-- 聊天框 -->
        <a-space direction="vertical">
          <a-card hoverable class="chat-card-item" :style="{ backgroundColor: '#fff' }">
            <a-typography-paragraph :style="{ marginTop: '1em' }">
              <div style="font-size: 18px; font-weight: bold;">你好！ 今天我能为您做点什么？</div>
              <div style="line-height: 30px;">
                <p style="font-weight: bold;">您可以这么问：</p>
                <p>1.
                  <a-link @click="handleExampleClick('帮我制定一个一周的健身计划')">帮我制定一个一周的健身计划</a-link>
                </p>
                <p>2.
                  <a-link @click="handleExampleClick('为我撰写有关AI对人类发展的5个吸引眼球的文章标题')">为我撰写有关AI对人类发展的5个吸引眼球的文章标题
                  </a-link>
                </p>
                <p>3.
                  <a-link @click="handleExampleClick('给我一份制作红烧肉的食谱')">给我一份制作红烧肉的食谱</a-link>
                </p>
              </div>
            </a-typography-paragraph>
          </a-card>
        </a-space>
      </a-space>
      <!-- 对话列表 -->
      <a-space class="chat-space-container" align="start" :size="10" v-for="(item, index) in chatList" :key="index">
        <!-- 头像 -->
        <a-avatar
            :style="{backgroundColor: item.role === 'assistant' ? '#fff' : '#165DFF', overflow: 'hidden'}"
            :size="32">
          <img v-if="item.role === 'assistant'"
               alt="avatar"
               style="height: 30px; width: 30px"
               :src="ChatGPTLogo"
          />
          <template v-else>
            <icon-robot/>
          </template>
        </a-avatar>
        <!-- 聊天框 -->
        <a-space direction="vertical">
          <a-card hoverable class="chat-card-item" :style="{
                borderColor: item.role === 'assistant' ? '' : '#c7d7fa',
                backgroundColor: item.role === 'assistant' ? '#fff': '#e8f3ff'
              }">
            <a-typography-paragraph copyable>
              <div class="chat-div" v-html="markedContent(item)"></div>
            </a-typography-paragraph>
            <template #actions>
              <a-dropdown @select="handleSelect" position="bl">
                <IconMore/>
                <template #content>
                  <a-doption>{{ $t('common.share') }}</a-doption>
                  <a-doption>{{ $t('common.save') }}</a-doption>
                </template>
              </a-dropdown>
            </template>
          </a-card>
          <a-button v-if="item.role === 'assistant' && sendLoading">
            <template #icon>
              <icon-record-stop/>
            </template>
            {{ $t('common.breakoff') }}
          </a-button>
        </a-space>
      </a-space>
    </a-space>
  </div>
</template>
<script>
</script>
<script setup>
import ChatGPTLogo from '@assets/images/chatgpt_black_logo.svg';
import {onMounted, reactive, ref, watch} from "vue";
import marked from "@plugins/markdown/marked.js";
import "highlight.js/styles/default.css";

const props = defineProps({
  list: {
    type: Array,
    default: []
  },
  loading: {
    type: Boolean,
    default: false
  }
});
let chatList = reactive(props.list)
const sendLoading = ref(props.loading);
const emits = defineEmits(['add:chat']);

watch(() => props.list, () => {
  chatList = props.list
})
watch(() => props.loading, () => {
  sendLoading.value = props.loading
})
// 渲染markdown
const markedContent = (item) => {
  if (item.role === 'assistant') {
    return marked.parse(item.content);
  } else {
    return item.content
  }
}
// const handleCopy = (e) => {
//   const textContent = e.target.parentNode.parentNode.parentNode.parentNode.firstElementChild.querySelector('.hljs').textContent;
// }

const handleExampleClick = (content) => {
  emits('add:chat', content, true)
}

const handleSelect = (e) => {
  console.log(e)
}
</script>

<style scoped>
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
  /*background: #e5e6ec;*/
}

.chat-container :deep(.arco-card-body p) {
  margin: 0px !important;
}

.chat-container :deep(.chat-div p code) {
  background: #f7f7f7;
  color: #800;
}

.chat-container :deep(.arco-typography-operation-copy), .chat-container :deep(.arco-typography-operation-copied) {
  position: absolute;
  bottom: 8px;
  right: 40px;
}
</style>