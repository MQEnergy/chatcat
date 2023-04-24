<template>
  <a-space class="menu-space-container scrollbar" direction="vertical" size="medium" style="">
    <!-- cate list -->
    <a-tooltip :content="$t('common.nocate')" position="right">
      <div :style="{position: 'relative', width: '100%'}">
        <a-avatar class="avatar flash" shape="square" :size="36"
                  :style="{backgroundColor: '#999'}"
                  @click="handleCateClick({id: 0, name: $t('common.nocate')}, 0)">
          W
        </a-avatar>
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
        <div :class="{'avatar-selected': currCateId === item.id}"></div>
      </div>
    </a-tooltip>
  </a-space>
  <!-- menu list -->
  <div class="menu-cate-container">
    <a-space direction="vertical" :size="20">
      <a-tooltip :content="$t('home.menuCate.settings.addCate')" position="right">
        <a-avatar :size="36"
                  :style="{ backgroundColor: '#3370ff', cursor: 'pointer' }"
                  @click="handleAddCate">
          <icon-plus/>
        </a-avatar>
      </a-tooltip>
      <div style="position: relative; cursor: pointer;">
        <icon-menu :style="{color: 'var(--color-text-2)'}" @click="handleMenuIconHover" size="26"/>
        <div v-if="isShow" class="menu-div">
          <a-menu class="menu-container" @menu-item-click="handleMenuClick">
            <a-menu-item key="0_3">{{ $t('home.menuCate.settings.sysSettings') }}</a-menu-item>
            <a-menu-item key="0_4">{{ $t('home.menuCate.settings.chat') }}</a-menu-item>
            <a-menu-item key="0_1">{{ $t('home.menuCate.settings.prompt') }}</a-menu-item>
            <a-menu-item key="0_2">{{ $t('home.menuCate.settings.upgrade') }}</a-menu-item>
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
import {defineEmits, h, onMounted, reactive, ref} from "vue";
import {Button, Message, Notification, Progress, TypographyParagraph} from '@arco-design/web-vue';
import CateAdd from "@views/home/components/cate-add.vue";
import {GetChatCateList, SetChatCateData} from "../../../../wailsjs/go/chat/Service.js";
import {useI18n} from "vue-i18n";
import {UpgradeService} from "../../../plugins/upgrade/index.js";
import marked from "@plugins/markdown/marked.js";
import "highlight.js/styles/default.css";

const emits = defineEmits(["select"])
const router = useRouter()
const {t} = useI18n()

const isShow = ref(false)
const visible = ref(false)
const cateList = reactive([])
const currCateId = ref(0)
const currPage = ref(1);

const handleMenuClick = (e) => {
  switch (e) {
    case '0_1':
      router.push('/settings/index?tab=4');
      break
    case '0_2':
      Message.info("Checking update");
      handleUpgrade()
      break
    case '0_3':
      router.push('/settings/index?tab=1');
      break
    case '0_4':
      router.push('/settings/index?tab=3');
      break
  }
}
const handleUpgrade = () => {
  const percentRef = ref(0)
  const loadingRef = ref(false)
  const upgradeService = new UpgradeService();
  const handleUpgradeNotification = (newVersion, releaseNotes) => {
    const id = `${Date.now()}`;
    releaseNotes = marked.parse(releaseNotes);
    const UpdateFooter = {
      render() {
        return h('div', [
          h(Progress, {
            percent: percentRef.value,
            style: {
              display: !loadingRef.value ? 'none' : 'block',
              marginBottom: '16px',
            },
          }),
          h(Button, {
            type: 'primary',
            size: 'small',
            loading: loadingRef.value,
            onClick: async () => {
              let percent = percentRef.value = 0
              loadingRef.value = true;
              const ticker = setInterval(() => {
                percent += 10;
                if (percent < 95 && percent >= 90) {
                  clearInterval(ticker);
                }
                percentRef.value = percent / 100;
              }, 100);

              upgradeService.downloadApp().then(() => {
                percentRef.value = 1;
                setTimeout(() => {
                  loadingRef.value = false;
                }, 500);
                upgradeService.installApp()
              })
            }
          }, t('common.upgrade.btn'))
        ])
      }
    }
    Notification.info({
      id,
      showIcon: false,
      title: `${newVersion}` + t('common.upgrade.notice'),
      content: h(TypographyParagraph, {
        style: {marginTop: '10px', height: '200px', overflowY: 'scroll'},
        innerHTML: releaseNotes
      }),
      closable: true,
      duration: 0,
      position: "bottomRight",
      footer: () => h(UpdateFooter),
    })
  }
// 检查更新
  upgradeService.checkUpdate().then(() => {
    if (!upgradeService.isUpdate) {
      Message.clear();
      Message.info(t('common.upgrade.noTips'));
      return
    }
    const newVersion = upgradeService.lastVersionInfo.version
    const releaseNotes = upgradeService.lastVersionInfo.versionDes
    Notification.clear()
    handleUpgradeNotification(newVersion, releaseNotes)
    setInterval(() => {
      Notification.clear()
      handleUpgradeNotification(newVersion, releaseNotes)
    }, 1000 * 60 * 10)
  })
}
const initCateList = (page) => {
  GetChatCateList(page).then(res => {
    cateList.splice(0, cateList.length);
    cateList.push(...res.data.list);
  })
}
onMounted(() => {
  if (window.go === undefined) {
    Message.error(t('common.panic'))
  } else {
    initCateList(currPage.value);
  }
})
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
// const handleUserProfile = () => {
//   Message.info("handle user profile")
// }
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
  height: calc(100vh - 180px);
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
  height: 100px;
  width: 100%;
  color: #fff;
}

.menu-div {
  width: 110px;
  overflow: hidden;
  border-bottom-right-radius: 6px;
  border-top-right-radius: 6px;
  position: absolute;
  z-index: 9;
  bottom: 0px;
  left: 53px;
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
  border-left: 4px solid rgb(var(--gray-9));
}
</style>