<template>
  <div class="layout-container">
    <a-layout style="height: 100vh;">
      <a-layout>
        <!-- menu cate sider -->
        <a-layout-sider class="layout-sider-cate-container">
          <menu-cate @select="handleCateList" :cateid="headerInfo.cateId"></menu-cate>
        </a-layout-sider>
        <!-- menu list sider -->
        <a-layout-sider class="layout-sider-chat-container">
          <menu-list :cateid="headerInfo.cateId" :chatid="headerInfo.chatId" @new:chat="addNewChat"
                     @select:chat="handleSelectChat"
                     @header:info="handleHeaderInfo"></menu-list>
        </a-layout-sider>
        <!-- chat content -->
        <a-layout-content class="chat-list-container">
          <a-layout>
            <a-layout-header>
              <content-header :info="headerInfo"></content-header>
            </a-layout-header>
            <a-layout-content class="scrollbar">
              <chat-list ref="chatListRef" @add:chat="handleChat" :chatid="headerInfo.chatId"
                         :cateid="headerInfo.cateId" @finish="handleFinished"
                         @model:info="handleHeaderInfo"></chat-list>
            </a-layout-content>
            <menu-tips @add:prompt="handlePromptToChat"></menu-tips>
            <!-- prompt input footer -->
            <a-layout-footer class="prompt-input-container">
              <prompt-input :value="prompt" @add:chat="handleChat" :checkoff="checkOffFlag"
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
import {onMounted, ref} from "vue";
import {useI18n} from "vue-i18n";

const {t} = useI18n();
const router = useRouter();
const prompt = ref('');
const chatListRef = ref(null);
// ----------------------------------------------------------------
const headerInfo = ref({
  cateName: t('common.nocate'),
  chatName: t('common.newchat'),
  cateId: 0,
  chatId: 0,
  modelName: 'loading',
  msgNum: 0,
  tokenNum: 0
})
const sendLoading = ref(false)
const checkOffFlag = ref(false)

const handleFinished = (value) => {
  sendLoading.value = value;
  checkOffFlag.value = value;
}
const handleHeaderInfo = (data) => {
  if (data.cateName !== undefined) {
    headerInfo.value.cateName = data.cateName;
  }
  if (data.chatName !== undefined) {
    headerInfo.value.chatName = data.chatName;
  }
  if (data.cateId !== undefined) {
    headerInfo.value.cateId = data.cateId;
  }
  if (data.chatId !== undefined) {
    headerInfo.value.chatId = data.chatId;
  }
  if (data.modelName !== undefined) {
    headerInfo.value.modelName = data.modelName;
  }
  if (data.tokenNum !== undefined) {
    headerInfo.value.tokenNum = data.tokenNum;
  }
  if (data.msgNum !== undefined) {
    headerInfo.value.msgNum = data.msgNum;
  }
}
// ----------------------------------------------------------------
onMounted(() => {
  const promptValue = router.currentRoute.value.query.prompt || '';
  const chatid = parseInt(router.currentRoute.value.query.chatid || 0);
  const cateid = parseInt(router.currentRoute.value.query.cateid || 0);
  prompt.value = promptValue;
  if (chatid !== 0) {
    headerInfo.value.chatId = chatid;
  }
  if (cateid !== 0) {
    headerInfo.value.cateId = cateid;
  }
})
// ----------------------------------------------------------------
const handleCateList = (item) => {
  headerInfo.value.cateId = item.id;
  headerInfo.value.chatId = 0;
  headerInfo.value.cateName = item.name;
}
const handlePromptToChat = (row) => {
  prompt.value = row.prompt
}
const addNewChat = () => {
  chatListRef.value.addNewChat();
}
const handleChat = (list, prompt, loading) => {
  sendLoading.value = loading;
  checkOffFlag.value = false;
  chatListRef.value.handleChat(list, prompt);
}
const handleSelectChat = (row) => {
  headerInfo.value.chatName = row.name;
  sendLoading.value = false;
  headerInfo.value.chatId = row.id;
}
// ----------------------------------------------------------------
</script>

<style scoped>
.layout-container :deep(.arco-space-item:nth-child(1)) {
  width: 100%;
}

.layout-container :deep(.arco-layout-sider-light) {
  box-shadow: none !important;
}

.layout-container .chat-list-container {
  background-color: var(--color-bg-2);
  position: relative;
  width: 100%;
  height: 100vh;
}

.layout-container .layout-sider-cate-container {
  width: 70px !important;
  background: var(--color-neutral-2);
  text-align: center;
}

.layout-container .layout-sider-chat-container {
  width: 240px !important;
  position: relative;
  border-right: 1px solid var(--color-neutral-3);
  /*background: var(--color-neutral-3);*/
}

.prompt-input-container {
  position: absolute;
  z-index: 9;
  /*padding: 0 20px;*/
  bottom: 0px;
  width: -webkit-fill-available;
}
</style>