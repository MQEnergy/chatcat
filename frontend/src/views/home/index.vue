<template>
  <div class="layout-container">
    <a-layout style="height: 100vh;">
      <a-layout>
        <!-- 侧边分类栏 -->
        <a-layout-sider style="width: 76px; background: #121212; text-align: center; ">
          <menu-cate :list="cateList" @select="handleCateList"></menu-cate>
        </a-layout-sider>
        <!-- 侧边chat内容列表栏 -->
        <a-layout-sider style="width: 240px; position: relative; background: var(--color-neutral-3);">
          <menu-list :cateid="currCateId" @add:chat="addNewChat" @select:chat="handleSelectChat"></menu-list>
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
              <chat-list class="chat-list-container" ref="chatListRef" :id="chatId"
                         @add:chat="handleChat" @finish="handleFinished"></chat-list>
            </a-layout-content>
            <!-- 菜单提示 -->
            <menu-tips @add:prompt="handlePromptToChat"></menu-tips>
            <!-- 对话输入框 -->
            <a-layout-footer class="prompt-input-container">
              <prompt-input :value="prompt" @ok="handleChat" :checkoff="checkOffFlag"
                            :loading="sendLoading"></prompt-input>
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
import {GetChatCateList} from "../../../wailsjs/go/chat/Service.js";
import {useI18n} from "vue-i18n";

const {t} = useI18n();
const router = useRouter();
const prompt = ref('');
const currCateId = ref(0);
// ----------------------------------------------------------------
// 滚动位置 Todo
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
const sendLoading = ref(false)
const checkOffFlag = ref(false)
let cateList = reactive([])
const chatId = ref(0)

const handleFinished = (value) => {
  sendLoading.value = value;
  checkOffFlag.value = value;
}
// 初始化分类列表
const initCateList = () => {
  GetChatCateList().then(res => {
    cateList.splice(0, cateList.length);
    cateList.push(...res.data.list);
  })
}
// ----------------------------------------------------------------
onMounted(() => {
  const promptValue = router.currentRoute.value.query.prompt || '';
  prompt.value = promptValue;
  initCateList();
})
// ----------------------------------------------------------------
const handleCateList = (item) => {
  currCateId.value = item.id;
  headerInfo.value.cateName = item.name;
}
const handlePromptToChat = (row) => {
  prompt.value = row.prompt
}
const addNewChat = () => {
  chatListRef.value.addNewChat();
}
const handleChat = (value, loading) => {
  sendLoading.value = loading;
  checkOffFlag.value = false;
  chatListRef.value.handleChat(value)
}
const handleSelectChat = (row) => {
  chatId.value = row.id;
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