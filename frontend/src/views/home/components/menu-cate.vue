<template>
  <a-space class="menu-space-container scrollbar" direction="vertical" size="medium" style="">
    <!-- cate list -->
    <a-tooltip content="未分类" position="right">
      <div :style="{position: 'relative', width: '100%'}">
        <a-avatar class="avatar flash" shape="square" :size="36"
                  :style="{backgroundColor: '#999'}"
                  @click="handleCateClick({id: 0, name: '未分类'}, 0)">
          W
        </a-avatar>
        <!-- selected cate -->
        <div :class="{'avatar-selected': currCateId === 0}"></div>
      </div>
    </a-tooltip>
    <a-tooltip :content="item.name" position="right" v-for="(item, index) in cateList">
      <div :style="{position: 'relative', width: '100%'}">
        <a-avatar class="avatar flash" :key="index" shape="square" :size="36"
                  :style="{backgroundColor: item.color}"
                  @click="handleCateClick(item, index)">
          {{ item.letter }}
        </a-avatar>
        <!-- selected cate -->
        <div :class="{'avatar-selected': currCateId === item.id}"></div>
      </div>
    </a-tooltip>
  </a-space>
  <!-- menu list -->
  <div class="menu-cate-container">
    <a-space direction="vertical" :size="20">
      <a-tooltip :content="$t('home.menuCate.settings.addCate')" position="right">
        <a-avatar :size="36"
                  :style="{ backgroundColor: '#ff7c00', cursor: 'pointer' }"
                  @click="handleAddCate">
          <icon-plus />
        </a-avatar>
      </a-tooltip>
      <a-avatar :size="36"
                :style="{ backgroundColor: '#3370ff', cursor: 'pointer' }"
                @click="handleUserProfile">
        <icon-robot/>
      </a-avatar>
      <div style="position: relative; cursor: pointer;">
        <icon-menu @click="handleMenuIconHover" size="26"/>
        <div v-if="isShow" class="menu-div">
          <a-menu theme="dark" @menu-item-click="handleMenuClick">
            <a-menu-item key="0_1">{{ $t('home.menuCate.settings.prompt') }}</a-menu-item>
            <a-menu-item key="0_2">{{ $t('home.menuCate.settings.upgrade') }}</a-menu-item>
            <a-menu-item key="0_3">{{ $t('home.menuCate.settings.sysSettings') }}</a-menu-item>
            <a-menu-item key="0_4">{{ $t('home.menuCate.settings.chat') }}</a-menu-item>
          </a-menu>
        </div>
      </div>
    </a-space>

    <!-- add cate -->
    <cate-add :visible="visible" @cancel="handleCateCancel" @ok="handleCateOk"></cate-add>
  </div>
</template>

<script setup>
import {useRouter} from 'vue-router'
import {defineEmits, reactive, ref, watch} from "vue";
import {Message} from "@arco-design/web-vue";
import CateAdd from "@views/home/components/cate-add.vue";
import {SetChatCateData} from "../../../../wailsjs/go/chat/Service.js";

const props = defineProps({
  list: {
    type: Array,
    default: []
  }
})
const emits = defineEmits(["select"])
const router = useRouter()
const isShow = ref(false)
const visible = ref(false)
let cateList = reactive(props.list)
const currCateId = ref(0)

watch(() => props.list, () => {
  cateList.value = props.list
})
// handleMenuClick 点击菜单
const handleMenuClick = (e) => {
  switch (e) {
    case '0_1':
      router.push('/settings/index?tab=4');
      break
    case '0_2':
      Message.info("检查更新中");
      break
    case '0_3':
      router.push('/settings/index?tab=1');
      break
    case '0_4':
      router.push('/settings/index?tab=3');
      break
  }
}
const handleMenuIconHover = (e) => {
  isShow.value = !isShow.value
}
const handleCateClick = (row, index) => {
  currCateId.value = row.id
  emits('select', row);
}
const handleAddCate = () => {
  visible.value = true;
}
const handleUserProfile = () => {
  Message.info("handle user profile")
}
const handleCateCancel = () => {
  visible.value = false;
}
const handleCateOk = (form) => {
  SetChatCateData(form).then((res) => {
    if (res.code !== 0) {
      Message.error(res.msg);
      return;
    }
    cateList.unshift({
      id: res.data.id,
      name: form.name,
      letter: res.data.letter,
      color: form.color
    })
    visible.value = false;
  })
}
</script>

<style scoped>
.menu-space-container {
  width: 100%;
  margin-top: 50px;
  height: calc(100vh - 230px);
  overflow-y: scroll;
  overflow-x: hidden;
}

.menu-cate-container :deep(.arco-menu-vertical .arco-menu-inner) {
  padding: 8px !important;
}

.menu-cate-container {
  position: absolute;
  bottom: 0px;
  left: 0px;
  height: 155px;
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

.avatar {
  position: relative;
  cursor: pointer;
}

.avatar-selected {
  position: absolute;
  left: 0px;
  top: 0;
  height: 35px;
  border-top-right-radius: 4px;
  border-bottom-right-radius: 4px;
  border-left: 4px solid #fff;
}

@keyframes flash {
  0% {
    opacity: 0.1;
  }
  50% {
    opacity: 0.5;
  }
  100% {
    opacity: 1;
  }
}

.flash {
  animation: flash 0.2s linear;
  animation-iteration-count: 1;
}

</style>