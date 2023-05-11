<template>
  <div class="page-header">
    <a-space style="width: 100%;" align="center">
      <a-page-header
          style="width: 100%; --wails-draggable: drag;"
          :title="props.info.cateName"
          :subtitle="props.info.chatName"
          :show-back="false"
      >
        <template #extra>
          <a-space>
            <a-tag>{{ props.info.modelName }}</a-tag>
            <a-tag>{{ props.info.msgNum }} messages</a-tag>
            <a-tag>{{ props.info.tokenNum }} tokens</a-tag>
          </a-space>
        </template>
      </a-page-header>
      <a-space v-if="isWin" size="medium" style="margin-right: 20px; font-size: 20px;">
        <icon-minus @click="handleMinize"/>
        <icon-refresh @click="handleRefresh"/>
        <icon-fullscreen v-if="isFullScreen" @click="handleFullscreen"/>
        <icon-fullscreen-exit v-else @click="handleNormalScreen"/>
        <icon-close @click="handleClose"/>
      </a-space>
    </a-space>
  </div>
</template>

<script setup>
import {Fullscreen, Minimise, NormalScreen, Quit, ReloadApp} from "../../../../wailsjs/go/setting/Service.js";
import {ref} from "vue";

const props = defineProps({
  info: {
    type: Object,
    default: {
      cateName: "",
      chatName: "",
      cateId: 0,
      chatId: 0,
      modelName: "",
      msgNum: 0,
      tokenNum: 0
    }
  }
})
const isFullScreen = ref(true);
const isWin = ref(window.isWindows);

const handleMinize = () => {
  Minimise()
}
const handleRefresh = () => {
  ReloadApp()
}
const handleClose = () => {
  Quit()
}
const handleFullscreen = () => {
  isFullScreen.value = false;
  Fullscreen()
}
const handleNormalScreen = () => {
  isFullScreen.value = true;
  NormalScreen()
}
</script>

<style scoped>
.page-header {
  background: var(--color-bg-2);
  top: 0;
  position: absolute;
  width: 100%;
  border-bottom: 1px solid var(--color-neutral-2);
  z-index: 9;
  height: 52px;
}

.page-header :deep(.arco-page-header) {
  padding: 10px 0 !important;
}
</style>