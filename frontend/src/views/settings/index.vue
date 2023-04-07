<template>
  <div class="settings-container">
    <div class="close-container">
      <icon-close-circle-fill @click="handleClose" size="40"/>
    </div>
    <a-layout class="layout-container">
      <a-layout-sider collapsible hide-trigger :collapsed="false" :style="{ width: '170px' }">
        <a-menu
            :default-selected-keys="tabName"
            :style="{ width: '170px' }"
            @menu-item-click="onClickMenuItem"
        >
          <a-menu-item key="1">
            <icon-settings/>
            {{ $t('settings.general') }}
          </a-menu-item>
          <a-menu-item key="2">
            <icon-sync/>
            {{ $t('settings.datasync') }}
          </a-menu-item>
          <a-menu-item key="3">
            <icon-select-all/>
            {{ $t('settings.chat') }}
          </a-menu-item>
          <a-menu-item key="4">
            <icon-select-all/>
            {{ $t('settings.prompt') }}
          </a-menu-item>
          <a-menu-item key="5">
            <icon-tool/>
            {{ $t('settings.contact') }}
          </a-menu-item>
          <a-menu-item key="6">
            <icon-download/>
            {{ $t('settings.releaseNotes') }}
          </a-menu-item>
        </a-menu>
      </a-layout-sider>
      <a-layout style="padding: 0 16px; border-radius: 8px; overflow: hidden">
        <a-layout-content>
          <general v-if="tabValue === '1'"></general>
          <datasync v-if="tabValue === '2'"></datasync>
          <chat v-if="tabValue === '3'"></chat>
          <prompt v-if="tabValue === '4'"></prompt>
          <contact v-if="tabValue === '5'"></contact>
          <release-notes v-if="tabValue === '6'"></release-notes>
        </a-layout-content>
      </a-layout>
    </a-layout>
  </div>
</template>

<script>
import {defineComponent, reactive, ref} from 'vue';
import {useRouter} from 'vue-router';
import General from "@views/settings/components/general.vue";
import Datasync from "@views/settings/components/datasync.vue";
import Prompt from "@views/settings/components/prompt.vue";
import Contact from "@views/settings/components/contact.vue";
import ReleaseNotes from "@views/settings/components/release-notes.vue";
import Chat from "@views/settings/components/chat.vue";

export default defineComponent({
  components: {Chat, ReleaseNotes, Contact, Prompt, Datasync, General},
  setup() {
    const route = useRouter()
    let {tab} = route.currentRoute.value.query
    let tabName = reactive([tab])
    const tabValue = ref(tab)
    const onClickMenuItem = (key) => {
      tabValue.value = key
      route.push(`/settings/index?tab=${key}`);
    }
    return {
      tabName,
      tabValue,
      onClickMenuItem
    }
  },
  methods: {
    handleClose() {
      const router = this.$router;
      const location = {path: '/index'}
      router.push(location);
    }
  }
});
</script>

<style scoped>
.settings-container {
  width: 85%;
  height: 90%;
  margin: 0 auto;
  position: relative;
  margin-top: 60px;
}

.close-container {
  position: absolute;
  top: 0;
  left: -50px;
  height: 40px;
  width: 40px;
  z-index: 10;
  cursor: pointer;
}

.layout-container {
  height: 100vh;
}

.layout-container :deep(.arco-layout-sider) .logo {
  height: 32px;
  margin: 12px 0px;
  text-align: center;
  font-size: 20px;
  line-height: 32px;
}

.layout-container :deep(.arco-menu-vertical .arco-menu-inner) {
  padding: 8px;
}

.layout-container :deep(.arco-layout-sider-light) {
  box-shadow: none;
  border-radius: 8px;
}

.layout-container :deep(.arco-layout-sider-children) {
  border-radius: 8px;
}

.layout-container :deep(.arco-layout-header) {
  height: 64px;
  line-height: 64px;
  background: var(--color-bg-3);
}

.layout-container :deep(.arco-layout-footer) {
  height: 48px;
  color: var(--color-text-2);
  font-weight: 400;
  font-size: 14px;
  line-height: 48px;
}

.layout-container :deep(.arco-layout-content) {
  color: var(--color-text-2);
  font-weight: 400;
  font-size: 14px;
  background: var(--color-bg-3);
  border-radius: 8px;
  padding: 24px;
}
</style>