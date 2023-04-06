<template>
  <div class="layout-demo">
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
              <chat-list :list="contentList"></chat-list>
            </a-layout-content>
            <!-- 菜单提示 -->
            <menu-tips></menu-tips>
            <!-- 对话输入框 -->
            <a-layout-footer class="prompt-input-container">
              <prompt-input :value="prompt"></prompt-input>
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
import {onMounted, reactive, ref} from "vue";

const router = useRouter();
const prompt = ref('');

onMounted(() => {
  // 输入框
  const promptValue = router.currentRoute.value.query.prompt || '';
  prompt.value = promptValue;
})

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
// 分类下的chat列表
let chatList = reactive({
  list: []
})

// 聊天内容列表
const contentList = reactive([
  {
    id: 1,
    content: '',
    avatar: '',
  }
])
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
      },{
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
      },{
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
      },{
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
      },{
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
      },{
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
      },{
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
      },{
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
</script>

<style scoped>
.layout-demo :deep(.arco-space-item:nth-child(1)) {
  width: 100%;
}

.layout-demo :deep(.arco-layout-sider) {
  width: 300px;
}

.layout-demo :deep(.arco-layout-content) {
  background-color: var(--color-bg-2);
}

.arco-icon {
  stroke-width: 2;
}

.absolute-div {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  height: auto;
  min-height: 80px;
  max-height: calc(100% - 124px);
  width: 100%;
  background-color: blue;
  overflow-y: scroll;
}

.prompt-input-container {
  position: absolute;
  z-index: 9;
  padding: 0 20px;
  bottom: 20px;
  width: -webkit-fill-available;
}
</style>