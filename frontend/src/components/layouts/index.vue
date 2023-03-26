<template>
  <div class="layout-demo">
    <a-layout style="height: 100vh;">
      <a-layout-header style="--wails-draggable: drag">聊天猫chatGPT工具</a-layout-header>
      <a-layout>
        <a-layout-sider style="width: 70px;">Sider</a-layout-sider>
        <a-layout-sider style="width: 245px; margin-left: 1px;">
          <a-list>
            <a-list-item v-for="idx in 4" :key="idx">
              <a-list-item-meta
                  title="Beijing Bytedance Technology Co., Ltd."
                  description="Beijing ByteDance Technology Co., Ltd. is an enterprise located in China."
              >
                <template #avatar>
                  <a-avatar shape="square">
                    <img
                        alt="avatar"
                        src="https://p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/3ee5f13fb09879ecb5185e440cef6eb9.png~tplv-uwbnlip3yd-webp.webp"
                    />
                  </a-avatar>
                </template>
              </a-list-item-meta>
              <template #actions>
                <icon-edit />
                <icon-delete />
              </template>
            </a-list-item>
          </a-list>
        </a-layout-sider>
        <a-layout-content>
          <router-view />
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
.layout-demo :deep(.arco-layout-sider-children),
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
  height: 64px;
  background-color: var(--color-primary-light-4);
}

.layout-demo :deep(.arco-layout-sider) {
  width: 206px;
  background-color: var(--color-primary-light-3);
}

.layout-demo :deep(.arco-layout-content) {
  background-color: rgb(var(--arcoblue-6));
}
</style>