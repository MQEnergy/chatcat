<template>
  <a-space direction="vertical" size="medium" style="margin-top: 50px;">
    <a-avatar v-for="(item, index) in cateList" :key="index" shape="square" :size="36"
              :style="{ backgroundColor: item.color }">{{ item.letter }}
    </a-avatar>
  </a-space>
  <div class="menu-cate-container">
    <a-space direction="vertical" :size="20">
      <a-avatar :size="36" :style="{ backgroundColor: '#3370ff' }">
        <icon-robot/>
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
import {reactive, ref, watch} from "vue";
import {Message} from "@arco-design/web-vue";

const props = defineProps({
  list: {
    type: Array,
    default: []
  }
})

const router = useRouter()
const isShow = ref(false)

let cateList = reactive(props.list)
watch(() => props.list, () => {
  cateList = props.list
})
// handleMenuClick 点击菜单
const handleMenuClick = (e) => {
  switch (e) {
    case '0_1':
      router.push('/settings/index?tab=3');
      break
    case '0_2':
      Message.info("检查更新中");
      break
    case '0_3':
      router.push('/settings/index?tab=1');
      break
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
  left: 56px;
}
</style>