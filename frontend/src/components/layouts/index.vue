<template>
  <div class="layout-demo">
    <a-layout style="height: 100vh;">
      <a-layout-header style="--wails-draggable: drag">聊天猫chatGPT工具</a-layout-header>
      <a-layout>
        <a-layout-sider style="width: 300px; position: relative; border-right: 1px solid var(--color-neutral-3)">
          <a-list :bordered="false">
            <a-list-item v-for="idx in 4" :key="idx">
              <a-list-item-meta
                  title="这是prompt提示词"
              >
              </a-list-item-meta>
              <template #actions>
                <icon-edit/>
                <icon-delete/>
              </template>
            </a-list-item>
          </a-list>
          <!--底部-->
          <div
              style="position: absolute; bottom: 0px; left: 0px; height: 40px; width: 100%; border-top: 1px solid #e3e3e3;">
            <a-avatar @click="toast" :style="{ backgroundColor: '#14C9C9' }">
              <IconUser/>
              <template #trigger-icon>
                <IconEdit/>
              </template>
            </a-avatar>
            setting
          </div>
        </a-layout-sider>
        <a-layout-content style="position: relative; width: 100%">
          <router-view/>
          <div style="position: absolute; bottom: 10px; padding: 0 10px; box-sizing: border-box; width: 100%">
            <a-space direction="horizontal" style="width: 100%;">
              <a-mention size="large" :data="['Bytedance', 'Bytedesign', 'Bytenumner']"
                         prefix="/" placeholder="input /"/>
              <!--                  <a-input :style="{width:'100%'}" placeholder="Please enter something" size="large" allow-clear />-->
              <a-button size="large" type="primary">send</a-button>
            </a-space>
          </div>
        </a-layout-content>
      </a-layout>
    </a-layout>
  </div>
</template>

<script setup>

</script>

<style scoped>
.layout-demo :deep(.arco-layout-header),
.layout-demo :deep(.arco-layout-footer),
.layout-demo :deep(.arco-layout-content) {
  display: flex;
  flex-direction: column;
  justify-content: center;
  color: var(--color-white);
  font-size: 16px;
  font-stretch: condensed;
  text-align: center;
}

.layout-demo :deep(.arco-space-item:nth-child(1)) {
  width: 100%;
}

.layout-demo :deep(.arco-layout-header),
.layout-demo :deep(.arco-layout-footer) {
  height: 50px;
  background-color: var(--color-primary-light-3);
}

.layout-demo :deep(.arco-layout-sider) {
  width: 300px;
  background-color: var(--color-bg-1);
}

.layout-demo :deep(.arco-layout-content) {
  background-color: var(--color-bg-2);
}
</style>