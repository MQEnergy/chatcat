<template>
  <div class="chat-container">
    <a-space class="chat-space-container" align="start" :size="10">
      <a-avatar
          class="chat-avatar"
          :size="32">
        <img alt="avatar" :src="ChatGPTLogo"/>
      </a-avatar>
      <!-- 聊天框 -->
      <a-space direction="vertical">
        <a-card hoverable class="chat-card-item" :style="{ backgroundColor: 'var(--color-bg-5)' }">
          <a-typography-paragraph :style="{ marginBottom: 0 }">
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

    <!-- chat list -->
    <a-space class="chat-space-container" direction="vertical" :size="10">
      <a-space v-for="(item, index) in chatList" align="start" :key="index">
        <template v-if="item.role !== 'system'">
          <!-- avatar -->
          <a-avatar
              class="chat-avatar"
              :style="{backgroundColor: item.role === 'assistant' ? '#fff' : '#165DFF', overflow: 'hidden'}"
              :size="32">
            <img v-if="item.role === 'assistant'"
                 alt="avatar"
                 :src="ChatGPTLogo"
            />
            <template v-else>
              <icon-robot/>
            </template>
          </a-avatar>
          <!-- chat -->
          <a-space direction="vertical">
            <a-card hoverable class="chat-card-item" :style="{
                backgroundColor: item.role === 'assistant' ? 'var(--color-bg-5)': 'var(--color-neutral-2)'
              }">
              <a-typography-paragraph copyable>
                <div class="chat-div" v-html="item.content"></div>
              </a-typography-paragraph>
              <template #actions>
                <a-dropdown @select="handleSelect" position="bl">
                  <IconMore/>
                  <template #content>
                    <a-doption :value="1">{{ $t('common.save') }}</a-doption>
                    <a-doption :value="2">{{ $t('common.del') }}</a-doption>
                  </template>
                </a-dropdown>
              </template>
            </a-card>
          </a-space>
        </template>
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
import {copyText} from "@plugins/copy/copy.js";
import {assembleReqChatList, assembleShowChatList} from "@plugins/assemble/prompt.js";

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
                chatList[chatList.length - 1].content = marked.parse(tempContent.value);
              }
              break;
            case -1: // stream error
              tempContent.value += data.data.data;
              if (tempContent.value.indexOf('response body closed') < 0) {
                chatList[chatList.length - 1].content = tempContent.value;
              }
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
    copyText('.hljs')
  }
})
const addNewChat = () => {
  chatList.splice(0, chatList.length);
}
const handleChat = (promptList, prompt) => {
  if (promptList.length === 0) {
    return
  }
  let showChatList = assembleShowChatList(promptList, prompt);
  showChatList.push({
    role: 'assistant',
    content: t('common.generate.start'),
  });
  let chatPromptList = JSON.parse(JSON.stringify(toRaw(showChatList)));
  tempContent.value = "";
  chatList.push(...chatPromptList)
  let reqPromptList = [];
  if (prompt.type === 2) {
    reqPromptList = assembleReqChatList(promptList, prompt);
  } else {
    reqPromptList = assembleReqChatList(chatList, prompt);
    reqPromptList.pop();
  }
  if (reqPromptList.length === 0) {
    return;
  }
  ChatCompletionStream(reqPromptList, clientId.value).then(res => {
    if (res.code === -1) {
      chatList[chatList.length - 1].content = res.msg
      emits('finish', false);
    }
  }).finally(() => {
    copyText('.hljs');
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
const handleExampleClick = (content) => {
  let promptList = [{
    role: 'user',
    prefix: '',
    content: content
  }]
  console.log("handleExampleClick", promptList);
  emits('add:chat', promptList, {
    type: 2,
  }, true)
}
const handleSelect = (e) => {
  switch (e) {
    case '1': // save
      break;
    case '2': // delete
      break;
  }
}

defineExpose({
  handleChat,
  addNewChat,
  initChatList
})
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
  border-top-right-radius: 6px;
  border-bottom-right-radius: 6px;
  border-bottom-left-radius: 6px;
}

.chat-space-container {
  margin-top: 10px;
}

.chat-container {
  width: 100%;
  padding: 20px;
  overflow-y: scroll;
  margin-top: 55px;
  height: 72vh;
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
  height: 26px;
  width: 26px;
  padding-top: 0.2rem;
}

.chat-container :deep(.hljs-btn-copy) {
  z-index: 90;
  border-radius: 4px;
  cursor: pointer;
  display: none;
  background: var(--color-bg-2);
}
</style>