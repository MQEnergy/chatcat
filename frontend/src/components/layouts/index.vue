<template>
  <div class="layout-demo">
    <a-layout style="height: 100vh;">
      <a-layout-header style="--wails-draggable: drag">聊天猫chatGPT工具</a-layout-header>
      <a-layout>
        <a-layout-sider style="width: 300px; position: relative">
          <a-list>
            <a-list-item v-for="idx in 4" :key="idx">
              <a-list-item-meta
                  title="Beijing Bytedance Technology Co., Ltd."
                  description="Beijing ByteDance Technology "
              >
              </a-list-item-meta>
              <template #actions>
                <icon-edit/>
                <icon-delete/>
              </template>
            </a-list-item>
          </a-list>
          <!--底部-->
          <div style="position: absolute; bottom: 0px; left: 0px; height: 40px; width: 100%; border-top: 1px solid #e3e3e3;">
            setting
          </div>
        </a-layout-sider>
        <a-layout-content>
          <router-view/>
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