<template>
  <a-space direction="vertical" size="medium" style="margin-top: 50px;">
    <a-avatar shape="square" :style="{ backgroundColor: '#3370ff' }">A</a-avatar>
    <a-avatar shape="square" :style="{ backgroundColor: '#14C9C9' }">B</a-avatar>
    <a-avatar shape="square" :style="{ backgroundColor: '#168CFF' }">C</a-avatar>
    <a-avatar shape="square" :style="{ backgroundColor: '#FF7D00' }">D</a-avatar>
    <a-avatar shape="square" :style="{ backgroundColor: '#FFC72E' }">E</a-avatar>
    <a-avatar shape="square" :style="{ backgroundColor: '#3370ff' }">F</a-avatar>
  </a-space>
  <div class="menu-cate-container">
    <a-space direction="vertical" :size="20">
      <a-avatar :size="30" :style="{ backgroundColor: '#3370ff' }">
        <IconUser/>
      </a-avatar>
      <div style="position: relative; cursor: pointer;">
        <icon-menu @click="handleMenuIconHover" size="26"/>
        <div v-if="isShow" class="menu-div">
          <a-menu theme="dark" @menu-item-click="handleMenuClick">
            <a-menu-item key="0_1">提示词</a-menu-item>
            <a-menu-item key="0_2">检查更新</a-menu-item>
            <a-menu-item key="0_3">系统设置</a-menu-item>
          </a-menu>
        </div>
      </div>
    </a-space>
  </div>
</template>

<script setup>
import {useRouter} from 'vue-router'
import {ref} from "vue";

const router = useRouter()
const isShow = ref(false)
// handleMenuClick 点击菜单
const handleMenuClick = (e) => {
  switch (e) {
    case '0_1':
      router.push('/settings/prompt')
    case '0_2':
      router.push('/settings/notice')
    case '0_3':
      router.push('/settings/index')
  }
}

const handleMenuIconHover = (e) => {
  isShow.value = !isShow.value
}
</script>

<style scoped>
.menu-cate-container :deep(.arco-menu-vertical .arco-menu-inner) {
  padding: 8px !important;
}

.menu-cate-container {
  position: absolute;
  bottom: 0px;
  left: 0px;
  height: 95px;
  width: 100%;
  color: #fff;
}

.menu-div {
  width: 110px;
  background: #414241;
  overflow: hidden;
  border-bottom-right-radius: 6px;
  border-top-right-radius: 6px;
  position: absolute;
  z-index: 9;
  bottom: 0px;
  left: 53px;
}
</style>