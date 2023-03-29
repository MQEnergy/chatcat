<template>
  <div class="layout-demo">
    <a-layout style="height: 100vh;">
      <a-layout>
        <!-- 侧边分类栏 -->
        <a-layout-sider style="width: 76px; background: #121212; text-align: center; ">
          <menu-cate></menu-cate>
        </a-layout-sider>

        <!-- 侧边内容列表栏 -->
        <a-layout-sider style="width: 240px; position: relative; background: var(--color-neutral-3)">
          <menu-list></menu-list>
        </a-layout-sider>

        <!-- 内容区 -->
        <a-layout-content style="position: relative; width: 100%; height: 100vh;">
          <a-layout>
            <!-- 头部 -->
            <a-layout-header>
              <content-header></content-header>
            </a-layout-header>
            <!-- 内容区域 -->
            <a-layout-content class="absolute-div">
              <div style="padding: 20px;">
                <a-space direction="vertical" :size="30">
                  <a-space align="start" :size="10" v-for="n in 10">
                    <a-avatar :style="{backgroundColor: '#165DFF'}" :size="32">A</a-avatar>
                    <a-card hoverable :style="{width: '500px', 'border-radius': '8px'}">
                      ByteDance's core product, Toutiao ("Headlines"), is a content platform in
                      China and around the world. Toutiao started out as a news recommendation
                      engine and gradually evolved into a platform delivering content in various
                      formats.
                      <template #actions>
                        <span class="icon-hover"><IconThumbUp/></span>
                        <span class="icon-hover"><IconShareInternal/></span>
                        <span class="icon-hover"><IconMore/></span>
                      </template>
                    </a-card>
                  </a-space>
                </a-space>
              </div>
            </a-layout-content>
            <!-- 菜单提示 -->
            <menu-tips></menu-tips>
            <!-- 对话输入框 -->
            <a-layout-footer class="prompt-input-container"
                             style="">
              <a-space direction="horizontal" style="width: 100%;">
                <a-mention size="large" :data="['翻译', '润色', '总结', '分析', '解释', '解释代码', '检查代码']"
                           prefix="/" placeholder="输入 / 或直接输入对话关键词"/>
                <a-button size="large" type="primary">send</a-button>
              </a-space>
            </a-layout-footer>
          </a-layout>
        </a-layout-content>
      </a-layout>
    </a-layout>
  </div>
</template>

<script setup>
import MenuTips from "@components/home/menu-tips.vue";
import MenuCate from "@components/home/menu-cate.vue";
import MenuList from "@components/home/menu-list.vue";
import ContentHeader from "@components/home/content-header.vue";</script>

<style scoped>
.layout-demo :deep(.arco-space-item:nth-child(1)) {
  width: 100%;
}

.layout-demo :deep(.arco-layout-sider) {
  width: 300px;
}

.layout-demo :deep(.arco-layout-content) {
  background-color: var(--color-bg-2);
}

.arco-icon {
  stroke-width: 2;
}

.absolute-div {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  height: auto;
  min-height: 80px;
  max-height: calc(100% - 124px);
  width: 100%;
  background-color: blue;
  overflow-y: scroll;
}

.prompt-input-container {
  position: absolute;
  z-index: 9;
  padding: 0 20px;
  bottom: 20px;
  width: -webkit-fill-available;
}
</style>