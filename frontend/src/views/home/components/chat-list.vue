<template>
  <div class="chat-container" style="padding: 20px;">
    <a-space direction="vertical" :size="20">
      <a-space class="chat-space-container" align="start" :size="10">
        <a-avatar
            class="chat-avatar"
            :size="32">
          <img alt="avatar" style="height: 28px; width: 28px" :src="ChatGPTLogo"/>
        </a-avatar>
        <!-- 聊天框 -->
        <a-space direction="vertical">
          <a-card hoverable class="chat-card-item" :style="{ backgroundColor: 'var(--color-bg-5)' }">
            <a-typography-paragraph :style="{ marginTop: '1em' }">
              <div style="font-size: 18px; font-weight: bold;">{{ $t('common.example.tips') }}</div>
              <div style="line-height: 30px;">
                <p style="font-weight: bold;">{{ $t('common.example.tips1') }}</p>
                <p v-for="(item, index) in exampleList" :key="index">
                  {{ index + 1 }}.
                  <a-link @click="handleExampleClick(item)">{{ item }}</a-link>
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
            class="chat-avatar"
            :style="{backgroundColor: item.role === 'assistant' ? '#fff' : '#165DFF', overflow: 'hidden'}"
            :size="32">
          <img v-if="item.role === 'assistant'"
               alt="avatar"
               style="height: 28px; width: 28px"
               :src="ChatGPTLogo"
          />
          <template v-else>
            <icon-robot/>
          </template>
        </a-avatar>
        <!-- 聊天框 -->
        <a-space direction="vertical">
          <a-card hoverable class="chat-card-item" :style="{
                borderColor: item.role === 'assistant' ? '' : '',
                backgroundColor: item.role === 'assistant' ? 'var(--color-bg-5)': 'var(--color-neutral-2)'
              }">
            <a-typography-paragraph copyable>
              <div class="chat-div" v-html="markedContent(item)"></div>
            </a-typography-paragraph>
            <template #actions>
              <a-dropdown @select="handleSelect" position="bl">
                <IconMore/>
                <template #content>
                  <!--                  <a-doption>{{ $t('common.share') }}</a-doption>-->
                  <a-doption>{{ $t('common.edit') }}</a-doption>
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
import {onMounted, reactive, ref, toRaw} from "vue";
import marked from "@plugins/markdown/marked.js";
import "highlight.js/styles/default.css";
import {
  ChatCompletionStream,
  GetChatRecordList,
  GetTokensNumFromMessages,
  GetWsUrl
} from "../../../../wailsjs/go/chat/Service.js";
import {useI18n} from "vue-i18n";
import {GetGeneralInfo} from "../../../../wailsjs/go/setting/Service.js";
import {Message} from "@arco-design/web-vue";

const {t} = useI18n();
let chatList = reactive([]);
let reqPromptList = reactive([]);
const clientId = ref("")
const tempContent = ref('');
const settingInfo = ref(null);
const exampleList = reactive(['帮我制定一个一周的健身计划', '为我撰写有关AI对人类发展的5个吸引眼球的文章标题', '给我一份制作红烧肉的食谱', '给我写一个python的http请求代码'])
const emits = defineEmits(['add:chat', 'finish', 'header:info']);

// ----------------------------------------------------------------
const initSettingInfo = () => {
  GetGeneralInfo().then(res => {
    settingInfo.value = res.data;
    emits('header:info', {
      modelName: res.data.chat_model,
    })
  })
}
const tokenNumFromMessage = (settings, list) => {
  GetTokensNumFromMessages(settings, list).then(res => {
    emits('header:info', {
      tokenNum: res.data,
      msgNum: list.length
    })
  })
}
// ----------------------------------------------------------------
// init ws
const initWs = () => {
  GetWsUrl().then(res => {
    let wsUrl = res + "?group_id=chat"
    const socket = new WebSocket(wsUrl);
    socket.addEventListener('open', (event) => {
    });
    socket.addEventListener('message', (event) => {
      const data = JSON.parse(event.data);
      switch (data.code) {
        case 0: // connect
          clientId.value = data.data.client_id;
          console.log(`ws connected：${event.data}`);
          break;
        case 10010: // message
          switch (data.data.code) {
            case 0: // streaming
              emits('finish', true);
              tempContent.value += data.data.data;
              if (chatList.length > 0) {
                chatList[chatList.length - 1].content = tempContent.value;
              }
              break;
            case -1: // stream error
              tempContent.value += data.data.data;
              chatList[chatList.length - 1].content = tempContent.value;
              emits('finish', false);
              break;
            case 1: // stream finished
              if (chatList.length > 0) {
                reqPromptList.push(chatList[chatList.length - 1]);
              }
              emits('finish', false);
              tokenNumFromMessage(settingInfo.value, reqPromptList);
              break;
          }
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
  if (window.go !== undefined) {
    initSettingInfo();
    initWs();
  }
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
// 初始化对话列表
const initChatList = (id, page) => {
  addNewChat()
  GetChatRecordList(id, page).then(res => {
    if (res.code === 0) {
      chatList.push(...res.data.list);
    }
  });
}
defineExpose({
  handleChat,
  addNewChat,
  initChatList
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

.chat-avatar {
  background-color: #fff;
  overflow: hidden;
  text-align: center;
}
.chat-avatar img {
  padding-top: 1px;
}
</style>