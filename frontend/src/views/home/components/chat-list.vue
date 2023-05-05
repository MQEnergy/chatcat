<template>
  <div class="chat-container">
    <a-space v-if="chatList.length === 0 && !recordLoading" class="chat-space-container flash" align="start" :size="10">
      <a-avatar
          class="chat-avatar"
          :size="32">
        <img alt="avatar" :src="ChatGPTLogo"/>
      </a-avatar>
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
      <template v-for="(item, index) in chatList" :key="index">
        <a-space v-if="item.role !== 'system'" class="flash" align="start">
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
                <div v-if="item.role === 'assistant' && regFlag" style="position: absolute; left: 10px;"
                     @click="handleRegenerate(item, index)">
                  <a-button type="text" size="mini" :loading="item.loading">
                    <template #icon>
                      <icon-refresh/>
                    </template>
                    {{ $t('common.regenerate') }}
                  </a-button>
                </div>
                <a-popconfirm
                    :cancel-text="$t('common.cancel')"
                    :ok-text="$t('common.ok')"
                    :content="$t('common.confirmDel')"
                    @ok="handleDelete(item, index)">
                  <icon-delete/>
                </a-popconfirm>
              </template>
            </a-card>
          </a-space>
        </a-space>
      </template>
    </a-space>
  </div>
</template>
<script>
</script>
<script setup>
import ChatGPTLogo from '@assets/images/chatgpt_black_logo.svg';
import {nextTick, onMounted, reactive, ref, toRaw, watch} from "vue";
import marked from "@plugins/markdown/marked.js";
import "highlight.js/styles/default.css";
import {
  ChatCompletionStream,
  DeleteChatRecord,
  GetChatRecordList,
  GetTokensNumFromMessages,
  GetWsUrl, SetChatRecordData
} from "../../../../wailsjs/go/chat/Service.js";
import {useI18n} from "vue-i18n";
import {GetGeneralInfo} from "../../../../wailsjs/go/setting/Service.js";
import {copyTextListener} from "@plugins/copy/copy.js";
import {assembleReqChatList, assembleShowChatList} from "@plugins/assemble/prompt.js";
import {Message} from "@arco-design/web-vue";

const {t} = useI18n();

const props = defineProps({
  cateid: {
    type: Number,
    default: 0
  },
  chatid: {
    type: Number,
    default: 0
  }
})
let chatList = reactive([]);
let reqPromptList = reactive([]);
const clientId = ref("")
const tempContent = ref('');
const settingInfo = ref(null);
const exampleList = reactive(['帮我制定一个一周的健身计划', '为我撰写有关AI对人类发展的5个吸引眼球的文章标题', '给我一份制作红烧肉的食谱', '给我写一个python的http请求代码'])
const emits = defineEmits(['add:chat', 'finish', 'model:info']);
const chatId = ref(0);
const cateId = ref(0);
const recordLoading = ref(false);
const curPage = ref(1);
const currIndex = ref(0);
const regFlag = ref(false);
const currLoading = ref(false);

// ----------------------------------------------------------------
const initChatRecordList = (chatid, page) => {
  recordLoading.value = true;
  // no page Todo
  GetChatRecordList(chatid, page).then(res => {
    if (res.code === 0) {
      regFlag.value = true;
      res.data.list.forEach((item) => {
        item.loading = false;
      })
      chatList.splice(0, chatList.length, ...res.data.list);
    }
  }).finally(() => {
    recordLoading.value = false;
    nextTick(() => {
      tokenNumFromMessage(settingInfo.value, chatList);
      copyTextListener('pre code');
    })
  })
}
watch(() => props.cateid, () => {
  cateId.value = props.cateid;
})
watch(() => props.chatid, () => {
  chatId.value = props.chatid;
  initChatRecordList(chatId.value, curPage.value);
})
// ----------------------------------------------------------------
const initSettingInfo = () => {
  GetGeneralInfo().then(res => {
    if (res.code !== 0) {
      return
    }
    settingInfo.value = res.data;
    emits('model:info', {
      modelName: res.data.chat_model,
    })
  })
}
const tokenNumFromMessage = (settings, list) => {
  GetTokensNumFromMessages(settings, list).then(res => {
    emits('model:info', {
      tokenNum: res.data,
      msgNum: chatList.length
    })
  })
}
// ----------------------------------------------------------------
const initWs = () => {
  GetWsUrl().then(res => {
    const socket = new WebSocket(res);
    socket.addEventListener('open', (event) => {
    });
    socket.addEventListener('message', (event) => {
      const data = JSON.parse(event.data);
      messageHandler(data);
    });
    socket.addEventListener('close', (event) => {
      Message.error("Websocket connect error, Please close and reopen.")
    });
    socket.addEventListener('error', (event) => {
      Message.error("Websocket connect error, Please close and reopen.")
    });
  })
};
const messageHandler = (data) => {
  switch (data.code) {
    case 0: // connect
      clientId.value = data.data.client_id;
      console.log(`ws connected：${event.data}`);
      break;
    case 10010: // message
      let chatIndex = currIndex.value !== 0 ? currIndex.value : chatList.length - 1;
      switch (data.data.code) {
        case 0: // streaming
          emits('finish', true);
          tempContent.value += data.data.data;
          if (chatList.length > 0) {
            chatList[chatIndex].content = marked.parse(tempContent.value);
          }
          break;
        case -1: // stream error
          tempContent.value += data.data.data;
          if (tempContent.value.indexOf('response body closed') < 0) {
            chatList[chatIndex].content = tempContent.value;
          }
          regFlag.value = true;
          chatList[chatIndex].loading = false;
          emits('finish', false);
          break;
        case 1: // stream finished
          if (chatList.length > 0) {
            reqPromptList.push(chatList[chatIndex]);
          }
          regFlag.value = true;
          chatList[chatIndex].loading = false;
          emits('finish', false);
          tokenNumFromMessage(settingInfo.value, reqPromptList);
          if (chatList.length > 0) {
            handleSaveChat(chatList[chatIndex]);
          }
          break;
      }
  }
}
onMounted(() => {
  if (window.go !== undefined) {
    initSettingInfo();
    initWs();
  }
})
const handleSaveChat = (chatInfo) => {
  const reqChatInfo = {
    cate_id: cateId.value,
    chat_id: chatId.value,
    name: "",
    content: chatInfo.content,
    reply_content: chatInfo.reply_content || '',
    prefix: chatInfo.prefix || '',
    role: chatInfo.role,
  }
  if (chatInfo.id) {
    reqChatInfo.id = chatInfo.id;
  }
  SetChatRecordData(reqChatInfo).then(res => {
    if (res.code !== 0) {
      Message.error(res.msg);
    }
  })
}
const addNewChat = () => {
  chatList.splice(0, chatList.length);
}
const handleChat = (promptList, prompt) => {
  if (promptList.length === 0) {
    return
  }
  if (chatList.length > 100) {
    Message.error(t('common.chat.limit'));
    nextTick(() => {
      emits('finish', false);
    })
    return;
  }
  if (prompt.type === 2) {
    handleSaveChat(promptList[0]);
  } else {
    if (promptList.length === 2) {
      if (promptList[0].content !== "") {
        handleSaveChat(promptList[0]);
      }
      handleSaveChat(promptList[1]);
    }
  }
  let showChatList = assembleShowChatList(promptList, prompt);
  showChatList.push({
    role: 'assistant',
    content: t('common.generate.start'),
  });
  let chatPromptList = JSON.parse(JSON.stringify(toRaw(showChatList)));
  tempContent.value = "";
  chatList.push(...chatPromptList)
  let reqPromptList;
  if (prompt.type === 2) {
    reqPromptList = assembleReqChatList(promptList, prompt);
  } else {
    reqPromptList = assembleReqChatList(chatList, prompt);
    reqPromptList.pop();
  }
  if (reqPromptList.length === 0) {
    return;
  }
  regFlag.value = false;
  ChatCompletionStream(reqPromptList, clientId.value).then(res => {
    if (res.code === -1) {
      chatList[chatList.length - 1].content = res.msg
      emits('finish', false);
    }
  }).finally(() => {
    copyTextListener('pre code');
  });
}
// ----------------------------------------------------------------
const handleExampleClick = (content) => {
  let promptList = [{
    role: 'user',
    prefix: '',
    content: content,
    reply_content: ''
  }]
  emits('add:chat', promptList, {
    type: 2,
  }, true)
}
const handleDelete = (row, index) => {
  DeleteChatRecord(row.id).then((res) => {
    if (res.code !== 0) {
      Message.error(res.msg)
      return;
    }
    initChatRecordList(chatId.value, curPage.value);
  })
}
const handleRegenerate = (row, index) => {
  if (row.loading || currLoading.value) {
    return
  }
  if (index == 0) {
    Message.error(t('common.chat.nopre'))
    return
  }
  currIndex.value = index;
  row.loading = true;
  currLoading.value = true;

  let userChatInfo = chatList[index - 1];
  if (userChatInfo.role !== 'user' || userChatInfo.content === "") {
    Message.error(t('common.chat.nopre'))
    return;
  }
  row.content = t('common.generate.start');
  let reqPromptList = [{
    role: userChatInfo.role,
    content: userChatInfo.prefix + userChatInfo.content + userChatInfo.reply_content,
  }]
  ChatCompletionStream(reqPromptList, clientId.value).then(res => {
    if (res.code === -1) {
      chatList[index].content = res.msg;
      emits('finish', false);
    }
  }).finally(() => {
    copyTextListener('pre code');
  });
}

defineExpose({
  handleChat,
  addNewChat,
  initChatRecordList
})
</script>

<style scoped>
.chat-div {
  overflow: scroll;
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
  height: 70vh;
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
  bottom: 9px;
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