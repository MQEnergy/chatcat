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
                  <a-link @click="handleExampleClick('为我撰写有关AI对人类发展的5个吸引眼球的文章标题')">
                    为我撰写有关AI对人类发展的5个吸引眼球的文章标题
                  </a-link>
                </p>
                <p>3.
                  <a-link @click="handleExampleClick('给我一份制作红烧肉的食谱')">给我一份制作红烧肉的食谱</a-link>
                </p>
                <p>4.
                  <a-link @click="handleExampleClick('给我写一个python的http请求代码')">给我写一个python的http请求代码
                  </a-link>
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
                  <a-doption>{{ $t('common.del') }}</a-doption>
                </template>
              </a-dropdown>
            </template>
          </a-card>
        </a-space>
      </a-space>
    </a-space>

  </div>
</template>
<script>
</script>
<script setup>
import ChatGPTLogo from '@assets/images/chatgpt_black_logo.svg';
import {onMounted, reactive, ref, watch, toRaw} from "vue";
import marked from "@plugins/markdown/marked.js";
import "highlight.js/styles/default.css";
import {ChatCompletionStream, GetChatRecordList, GetWsUrl} from "../../../../wailsjs/go/chat/Service.js";
import {useI18n} from "vue-i18n";

const {t} = useI18n();
const props = defineProps({
  id: {
    type: Number,
    default: 0
  },
});

let chatList = reactive([]);
let reqPromptList = reactive([]);
const clientId = ref("")
const tempContent = ref('');

const emits = defineEmits(['add:chat', 'finish']);

// ----------------------------------------------------------------
// 初始化ws
const initWs = () => {
  GetWsUrl().then(res => {
    let wsUrl = res + "?group_id=chat"
    const socket = new WebSocket(wsUrl);
    socket.addEventListener('open', (event) => {
    });
    socket.addEventListener('message', (event) => {
      const data = JSON.parse(event.data);
      switch (data.code) {
        case 0:
          clientId.value = data.data.client_id;
          console.log(`ws connected：${event.data}`);
          break;
        case 10010:
          if (data.data.code === 0) {
            emits('finish', true);
            tempContent.value += data.data.data;
            if (chatList.length > 0) {
              chatList[chatList.length - 1].content = tempContent.value;
            }
          } else {
            if (chatList.length > 0) {
              reqPromptList.push(chatList[chatList.length - 1]);
            }
            emits('finish', false);
          }
          break;
      }
    });
    socket.addEventListener('close', (event) => {
      console.log('websocket closed');
    });
    socket.addEventListener('error', (event) => {
      console.error('websocker error');
    });
  })
};
onMounted(() => {
  initWs();
})
const addNewChat = () => {
  chatList.splice(0, chatList.length);
}
// 对话
const handleChat = (value) => {
  value = value.trim();
  let promptList = [{
    role: 'user',
    content: value,
  }];
  tempContent.value = "";
  chatList.push(...promptList)
  let _deepList = JSON.parse(JSON.stringify(toRaw(promptList)));
  _deepList[0].content = value + "," + t('common.markdown');
  chatList.push({
    role: 'assistant',
    content: t('common.generate.start'),
  })
  reqPromptList.push(_deepList[0])
  reqPromptList = reqPromptList.slice(-4);
  ChatCompletionStream(reqPromptList, clientId.value).then(res => {
    if (res.code === -1) {
      chatList[chatList.length - 1].content = res.msg
    }
  });
}
// ----------------------------------------------------------------
defineExpose({
  handleChat,
  addNewChat
})
// 初始化对话列表
const initChatList = (id, page) => {
  GetChatRecordList(id, page).then(res => {
    console.log('handleSelectChat', id, res.data.list);
    if (res.code === 0) {
      chatList.push(...res.data.list);
    }
  });
}
watch(() => props.id, () => {
  initChatList(props.id, 1)
})
const markedContent = (item) => {
  return marked.parse(item.content);
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