<template>
  <div class="chat-container" ref="chatRecordRef">
    <a-space class="chat-space-container flash"
             v-if="chatList.length === 0 && !recordLoading"
             align="start" :size="10">
      <a-avatar
          class="chat-avatar"
          :size="32">
        <img alt="avatar" :src="ChatGPTLogo"/>
      </a-avatar>
      <a-space direction="vertical">
        <a-card class="chat-card-item">
          <a-typography-paragraph class="example-assistant-p">
            <div class="assistant-p-title">{{ $t('common.example.tips') }}</div>
            <div style="line-height: 30px;">
              <p style="font-weight: bold; color: #fff;">{{ $t('common.example.tips1') }}</p>
              <p v-for="(item, index) in exampleList" :key="index" style="color: #fff">
                {{ index + 1 }}.
                <a-link @click="handleExampleClick(item)" :style="{color: '#fff', background: 'none'}">
                  {{ item }}
                </a-link>
              </p>
            </div>
          </a-typography-paragraph>
        </a-card>
      </a-space>
    </a-space>

    <!-- chat list -->
    <a-space class="chat-space-container" direction="vertical" :size="20">
      <template v-for="(item, index) in chatList" :key="index">
        <a-space v-if="item.role !== 'system'" class="flash" align="start">
          <!-- avatar -->
          <a-avatar
              class="chat-avatar"
              :style="{backgroundColor: item.role === 'assistant' ? '#fff' : 'rgb(var(--purple-5))', overflow: 'hidden'}"
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
            <a-card class="chat-card-item" :style="{
                backgroundColor: item.role === 'assistant' ? 'var(--color-border-1)' : 'rgb(var(--purple-5))'
              }">
              <a-typography-paragraph copyable>
                <template v-if="item.role === 'assistant' && item.content === $t('common.generate.start')">
                  <a-spin size="6" dot/>
                </template>
                <div v-else class="chat-div scrollbar"
                     :style="{color: item.role === 'assistant' ? 'rgb(var(--gray-10))' : '#fff' }"
                     v-html="item.content"></div>
                <template #copy-tooltip>
                  {{ $t('common.copy') }}
                </template>
                <template #copy-icon={copied} v-if="regFlag && !currLoading">
                  <icon-copy v-if="!copied"
                             :style="{color: item.role === 'assistant' ? 'rgb(var(--gray-10))' : '#fff' }"/>
                  <icon-check-circle-fill v-else
                                          :style="{color: item.role === 'assistant' ? 'rgb(var(--gray-10))' : '#fff' }"/>
                </template>
              </a-typography-paragraph>

              <template #actions>
                <div v-if="item.reg_flag && regFlag" style="position: absolute; left: 10px;"
                     @click="handleRegenerate(item, index)">
                  <a-button type="text" size="mini" :style="{color: 'rgb(var(--purple-5))'}" :loading="item.loading">
                    <template #icon>
                      <icon-refresh/>
                    </template>
                    {{ $t('common.regenerate') }}
                  </a-button>
                </div>
                <a-popconfirm
                    v-if="regFlag && !currLoading"
                    :cancel-text="$t('common.cancel')"
                    :ok-text="$t('common.ok')"
                    :content="$t('common.confirmDel')"
                    @ok="handleDelete(item, index)">
                  <icon-delete :style="{color: item.role === 'assistant' ? 'rgb(var(--gray-10))' : '#fff' }"/>
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
  BreakOffChatStream,
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
const chatRecordRef = ref(null);

// ----------------------------------------------------------------
const initChatRecordList = (chatid, page) => {
  recordLoading.value = true;
  // no page 100 limit Todo
  GetChatRecordList(chatid, page).then(res => {
    if (res.code === 0) {
      regFlag.value = true;
      res.data.list.forEach((item, index) => {
        item.loading = false;
        item.reg_flag = false;
        if (item.role === 'user') {
          item.content = marked.parse(item.content);
        }
        if (item.role === 'assistant') {
          item.reg_flag = true;
          if (index === 0) {
            item.reg_flag = false;
          } else if (res.data.list[index - 1].role === 'assistant') {
            item.reg_flag = false;
          }
        }
      })
      chatList.splice(0, chatList.length, ...res.data.list);
    }
  }).finally(() => {
    nextTick(() => {
      recordLoading.value = false;
      handleScrollToBottom();
      tokenNumFromMessage(settingInfo.value, chatList);
      copyTextListener('pre code');
    })
  })
}
const handleScrollToBottom = () => {
  if (currIndex.value === 0) {
    setTimeout(() => {
      chatRecordRef.value.scrollTo({
        top: chatRecordRef.value.scrollHeight,
        behavior: "smooth",
      });
    }, 200);
  }
}
watch(() => props.cateid, () => {
  cateId.value = props.cateid;
  resetFlag(0);
})
watch(() => props.chatid, () => {
  chatId.value = props.chatid;
  resetFlag(0);
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
      Message.error("Websocket connect close, Please close and reopen.")
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
      handleScrollToBottom();
      switch (data.data.code) {
        case 0: // streaming
          tempContent.value += data.data.data;
          if (chatList.length > 0) {
            chatList[chatIndex].content = marked.parse(tempContent.value);
          }
          emits('finish', true);
          break;
        case -1: // stream error
          tempContent.value += data.data.data;
          if (tempContent.value.indexOf('response body closed') < 0) {
            chatList[chatIndex].content = tempContent.value;
          }
          resetFlag(chatIndex);
          emits('finish', false);
          break;
        case 1: // stream finished
          if (chatList.length > 0) {
            reqPromptList.push(chatList[chatIndex]);
          }
          resetFlag(chatIndex);
          emits('finish', false);
          tokenNumFromMessage(settingInfo.value, reqPromptList);
          if (chatList.length > 0) {
            handleSaveChat(chatList[chatIndex]);
          }
          break;
      }
  }
}
const resetFlag = (chatIndex) => {
  BreakOffChatStream()
  regFlag.value = true;
  currLoading.value = false;
  currIndex.value = 0;
  if (chatList.length > 0) {
    chatList[chatIndex].loading = false;
    chatList[chatIndex].reg_flag = true;
  }
}
onMounted(() => {
  if (window.go !== undefined) {
    initSettingInfo();
    initWs();
    copyTextListener('pre code');
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
  if (chatRecordRef.value) {
    chatRecordRef.value.scrollTop = chatRecordRef.value.scrollHeight;
  }
  currIndex.value = 0;
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
  chatPromptList.forEach((item, index) => {
    item.loading = false;
    item.reg_flag = false;
    if (item.role === 'user') {
      item.content = marked.parse(item.content);
    }
    if (item.role === 'assistant') {
      item.reg_flag = true;
      if (index === 0) {
        item.reg_flag = false;
      } else if (chatPromptList[index - 1].role === 'assistant') {
        item.reg_flag = false;
      }
    }
  });
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
    reply_content: '',
    reg_flag: false,
    loading: false,
  }]
  emits('add:chat', promptList, {
    type: 2,
  }, true)
}
const handleDelete = (row, index) => {
  if (currLoading.value) return;
  currIndex.value = index;
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
  let userChatInfo = chatList[index - 1];
  if (userChatInfo.role !== 'user' || userChatInfo.content === "") {
    Message.error(t('common.chat.nopre'))
    return;
  }
  tempContent.value = "";
  currIndex.value = index;
  row.loading = true;
  currLoading.value = true;

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
.chat-card-item {
  width: 510px;
  border-top-right-radius: 10px;
  border-bottom-right-radius: 10px;
  border-bottom-left-radius: 10px;
  background-color: rgb(var(--purple-5));
  border: none;
}

.chat-card-item .example-assistant-p {
  margin-bottom: 0px;
}

.chat-card-item .example-assistant-p .assistant-p-title {
  font-size: 18px;
  font-weight: bold;
  color: #fff;
}

.chat-card-item .example-assistant-p .arco-link:hover {
  border-bottom: 1px solid #fff;
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
  animation: slide-down 1s ease-out forwards;
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
  color: #fff;
  background: none;
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
  background: var(--color-border-1);
}

@keyframes scroll-down {
  0% {
    opacity: 0;
    transform: translateY(-20px);
  }
  50% {
    opacity: 1;
  }
  100% {
    opacity: 0;
    transform: translateY(20px);
  }
}
</style>