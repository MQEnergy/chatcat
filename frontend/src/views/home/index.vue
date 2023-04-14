<template>
  <div class="layout-container">
    <a-layout style="height: 100vh;">
      <a-layout>
        <!-- 侧边分类栏 -->
        <a-layout-sider style="width: 76px; background: #121212; text-align: center; ">
          <menu-cate :list="cateList" @select="handleCateList"></menu-cate>
        </a-layout-sider>
        <!-- 侧边chat内容列表栏 -->
        <a-layout-sider style="width: 240px; position: relative; background: var(--color-neutral-3)">
          <menu-list :list="chatList.list"></menu-list>
        </a-layout-sider>
        <!-- 当前chat内容区 -->
        <a-layout-content style="position: relative; width: 100%; height: 100vh;">
          <a-layout>
            <!-- 头部 -->
            <a-layout-header>
              <content-header :info="headerInfo"></content-header>
            </a-layout-header>
            <!-- 内容区域 -->
            <a-layout-content class="absolute-div scrollbar">
              <chat-list class="chat-list-container" ref="chatListRef" :list="contentList"
                         :loading="sendLoading" @add:chat="handleChat"></chat-list>
            </a-layout-content>
            <!-- 菜单提示 -->
            <menu-tips @add:prompt="handlePromptToChat"></menu-tips>
            <!-- 对话输入框 -->
            <a-layout-footer class="prompt-input-container">
              <prompt-input :value="prompt" @ok="handleChat" :loading="sendLoading"></prompt-input>
            </a-layout-footer>
          </a-layout>
        </a-layout-content>
      </a-layout>
    </a-layout>
  </div>
</template>

<script setup>
import MenuTips from "@views/home/components/menu-tips.vue";
import MenuCate from "@views/home/components/menu-cate.vue";
import MenuList from "@views/home/components/menu-list.vue";
import ContentHeader from "@views/home/components/content-header.vue";
import ChatList from "@views/home/components/chat-list.vue";
import PromptInput from "@views/home/components/prompt-input.vue";
import {useRouter} from "vue-router";
import {onMounted, reactive, ref, toRaw} from "vue";
import {ChatCompletionStream, GetWsUrl} from "../../../wailsjs/go/chat/Service.js";
import {useI18n} from "vue-i18n";

const {t} = useI18n();
const router = useRouter();
const prompt = ref('');

let contentList = reactive([]);
const tempContent = ref('');
let deepList = reactive([]);
// ----------------------------------------------------------------
// 滚动位置
const chatListRef = ref(null);
const chatListHeight = ref(0);
// ----------------------------------------------------------------
// 头部显示
const headerInfo = ref({
  cateName: '这是分类名称',
  chatName: '这是对话语句',
  modelName: 'GPT3.5',
  msgNum: '16',
  tokenNum: '193'
})
// 分类列表
const cateList = reactive([
  {
    id: 1,
    cateName: '前端',
    letter: 'Q',
    color: '#3370ff'
  }, {
    id: 2,
    cateName: '后端',
    letter: 'H',
    color: '#14c9c9'
  }, {
    id: 3,
    cateName: '设计',
    letter: 'S',
    color: '#ff7d00'
  }, {
    id: 4,
    cateName: '服务器',
    letter: 'F',
    color: '#ffc72e'
  }
])
let chatList = reactive({
  list: []
})
const sendLoading = ref(false)
// ws
const initWs = () => {
  GetWsUrl().then(res => {
    let wsUrl = res + "?group_id=ask"
    const socket = new WebSocket(wsUrl);
    socket.addEventListener('open', (event) => {
    });
    socket.addEventListener('message', (event) => {
      const data = JSON.parse(event.data);
      switch (data.code) {
        case 0:
          console.log(`ws连接成功：${event.data}`);
          break;
        case 10010:
          if (data.data.code === 0) {
            tempContent.value += data.data.data;
            contentList[contentList.length - 1].content = tempContent.value;
          } else {
            deepList.push(contentList[contentList.length - 1]);
            sendLoading.value = false;
          }
          break;
      }
    });
    socket.addEventListener('close', (event) => {
      console.log('WebSocket连接已关闭');
    });
    socket.addEventListener('error', (event) => {
      console.error('WebSocket连接出错');
    });
  })
}
// ----------------------------------------------------------------
onMounted(() => {
  const promptValue = router.currentRoute.value.query.prompt || '';
  prompt.value = promptValue;
  initWs();
})
// ----------------------------------------------------------------
const handleCateList = (item) => {
  let _chatList = [];
  switch (item.id) {
    case 1:
      _chatList = [{
        id: 1,
        name: '这是chat对话语句1',
        sort: 50
      }, {
        id: 2,
        name: '这是chat对话语句2',
        sort: 50
      }, {
        id: 3,
        name: '这是chat对话语句3',
        sort: 50
      }, {
        id: 1,
        name: '这是chat对话语句1',
        sort: 50
      }, {
        id: 2,
        name: '这是chat对话语句2',
        sort: 50
      }, {
        id: 3,
        name: '这是chat对话语句3',
        sort: 50
      }, {
        id: 1,
        name: '这是chat对话语句1',
        sort: 50
      }, {
        id: 2,
        name: '这是chat对话语句2',
        sort: 50
      }, {
        id: 3,
        name: '这是chat对话语句3',
        sort: 50
      }, {
        id: 1,
        name: '这是chat对话语句1',
        sort: 50
      }, {
        id: 2,
        name: '这是chat对话语句2',
        sort: 50
      }, {
        id: 3,
        name: '这是chat对话语句3',
        sort: 50
      }, {
        id: 1,
        name: '这是chat对话语句1',
        sort: 50
      }, {
        id: 2,
        name: '这是chat对话语句2',
        sort: 50
      }, {
        id: 3,
        name: '这是chat对话语句3',
        sort: 50
      }, {
        id: 1,
        name: '这是chat对话语句1',
        sort: 50
      }, {
        id: 2,
        name: '这是chat对话语句2',
        sort: 50
      }, {
        id: 3,
        name: '这是chat对话语句3',
        sort: 50
      }, {
        id: 1,
        name: '这是chat对话语句1',
        sort: 50
      }, {
        id: 2,
        name: '这是chat对话语句2',
        sort: 50
      }, {
        id: 3,
        name: '这是chat对话语句3',
        sort: 50
      }, {
        id: 1,
        name: '这是chat对话语句1',
        sort: 50
      }, {
        id: 2,
        name: '这是chat对话语句2',
        sort: 50
      }, {
        id: 3,
        name: '这是chat对话语句3',
        sort: 50
      }]
      break;
    case 2:
      _chatList = [{
        id: 1,
        name: '这是chat对话语句4',
        sort: 50
      }, {
        id: 2,
        name: '这是chat对话语句5',
        sort: 50
      }, {
        id: 3,
        name: '这是chat对话语句6',
        sort: 50
      }]
      break;
  }
  chatList.list = _chatList
}
const handlePromptToChat = (row) => {
  prompt.value = row.prompt
}

const handleChat = (value, loading) => {
  tempContent.value = '';
  sendLoading.value = loading;
  let promptList = [{
    role: 'user',
    content: value,
  }];
  contentList.push(...promptList)
  let _deepList = JSON.parse(JSON.stringify(toRaw(promptList)));
  _deepList[0].content = value + "," + t('common.markdown');
  contentList.push({
    role: 'assistant',
    content: "正在生成中...",
  })
  // chat stream
  deepList.push(_deepList[0])
  // 只保留最后四个
  deepList = deepList.slice(-4);
  console.log("deepList: ", deepList)
  ChatCompletionStream(deepList).then(res => {
    if (res.code === -1) {
      contentList[contentList.length - 1].content = res.msg
    }
  })
}
// ----------------------------------------------------------------
</script>

<style scoped>
.layout-container :deep(.arco-space-item:nth-child(1)) {
  width: 100%;
}

.layout-container :deep(.arco-layout-sider) {
  width: 300px;
}

.layout-container :deep(.arco-layout-content) {
  background-color: var(--color-bg-2);
}

.prompt-input-container {
  position: absolute;
  z-index: 9;
  padding: 0 20px;
  bottom: 20px;
  width: -webkit-fill-available;
}
</style>