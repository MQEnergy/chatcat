<template>
  <div style="width: 80%; margin: 0 auto; position: relative; margin-top: 60px;">
    <div class="close-container">
      <icon-close-circle-fill @click="handleClose" size="40"/>
    </div>
    <a-layout class="layout-demo">
      <a-layout-sider collapsible hide-trigger :collapsed="false">
        <div class="logo">
          基础设置
        </div>
        <a-menu
            :default-selected-keys="tabName"
            :style="{ width: '100%' }"
            @menu-item-click="onClickMenuItem"
        >
          <a-menu-item key="1">
            <icon-settings />
            通用配置
          </a-menu-item>
          <a-menu-item key="2">
            <icon-sync />
            数据同步
          </a-menu-item>
          <a-menu-item key="3">
            <icon-select-all />
            提示词管理
          </a-menu-item>
          <a-menu-item key="4">
            <icon-message />
            更新信息
          </a-menu-item>
          <a-menu-item key="5">
            <icon-tool />
            其他工具
          </a-menu-item>
        </a-menu>
      </a-layout-sider>
      <a-layout>
        <a-layout style="padding: 0 16px; border-radius: 8px;">
          <a-layout-content>
            <general></general>
          </a-layout-content>
        </a-layout>
      </a-layout>
    </a-layout>
  </div>
</template>

<script>
import {defineComponent, reactive} from 'vue';
import { useRouter } from 'vue-router';
import {Message} from '@arco-design/web-vue';
import General from "@views/settings/components/general.vue";

export default defineComponent({
  components: {General},
  setup() {
    const route = useRouter()
    const {tab} = route.currentRoute.value.query
    const tabName = reactive([tab])

    return {
      tabName
    }
  },
  methods: {
    onClickMenuItem(key) {
      Message.info({content: `You select ${key}`, showIcon: true});
    },
    handleClose() {
      const router = this.$router;
      const location = {path: '/index'}
      router.push(location);
    }
  }
});
</script>

<style scoped>
.close-container {
  position: absolute;
  top: 0;
  left: -50px;
  height: 40px;
  width: 40px;
  z-index: 10;
  cursor: pointer;
}

.layout-demo :deep(.arco-layout-sider) .logo {
  height: 32px;
  margin: 12px 0px;
  text-align: center;
  font-size: 20px;
  line-height: 32px;
}

.layout-demo :deep(.arco-layout-sider-light) {
  box-shadow: none;
  border-radius: 8px;
}

.layout-demo :deep(.arco-layout-sider-children) {
  border-radius: 8px;
}

.layout-demo :deep(.arco-layout-header) {
  height: 64px;
  line-height: 64px;
  background: var(--color-bg-3);
}

.layout-demo :deep(.arco-layout-footer) {
  height: 48px;
  color: var(--color-text-2);
  font-weight: 400;
  font-size: 14px;
  line-height: 48px;
}

.layout-demo :deep(.arco-layout-content) {
  color: var(--color-text-2);
  font-weight: 400;
  font-size: 14px;
  background: var(--color-bg-3);
  border-radius: 8px;
  padding: 24px;
}
</style>