<template>
  <div class="prompt-container">
    <a-card title="分类管理" :bordered="false" :header-style="{borderColor: 'var(--color-fill-2)'}">
      <template #extra>
        <a-button type="primary" @click="handleAddCate">
          <template #icon>
            <icon-plus/>
          </template>
          分类
        </a-button>
      </template>
      <!-- 分类列表 -->
      <a-space>
        <a-row :gutter="[10, 10]">
          <a-col :span="8" v-for="(item, index) in cateList" :key="index">
            <a-card hoverable :style="{ borderRadius: '4px' }">
              <div :style="{
                  display: 'flex',
                  alignItems: 'center',
                  justifyContent: 'space-between',
                }">
                <span :style="{ display: 'flex', alignItems: 'center', color: '#1D2129' }">
                  <a-avatar :style="{ marginRight: '8px', backgroundColor: '#165DFF' }" :size="28">
                    A
                  </a-avatar>
                  <a-typography-text>{{ item.name }}</a-typography-text>
                </span>
                <a-link @click="handlePromptClick">查看</a-link>
              </div>
            </a-card>
          </a-col>
        </a-row>
      </a-space>
    </a-card>
  </div>
  <!-- 添加分类 -->
  <prompt-add :visible="visible" @cancel="handleCancel"></prompt-add>
  <!-- 查看prompt关键词列表 -->
  <prompt-drawer :visible="drawerSeen" @cancel="handleCancel"></prompt-drawer>
</template>

<script setup>
import {reactive, ref} from "vue";
import PromptAdd from "@views/settings/components/prompt-add.vue";
import PromptDrawer from "@views/settings/components/prompt-drawer.vue";

const visible = ref(false);
const drawerSeen = ref(false);
const cateList = reactive([
  {
    id: 1,
    name: '分类1'
  },
  {
    id: 2,
    name: '分类2'
  },
  {
    id: 3,
    name: '分类3'
  },
  {
    id: 4,
    name: '分类4'
  },
  {
    id: 5,
    name: '分类5'
  },
  {
    id: 6,
    name: '分类6'
  },
  {
    id: 7,
    name: '分类7'
  },
  {
    id: 8,
    name: '分类8'
  },
  {
    id: 9,
    name: '分类9'
  },
  {
    id: 10,
    name: '分类10'
  },
  {
    id: 11,
    name: '分类11'
  },
]);
const handleCancel = (e) => {
  visible.value = e;
  drawerSeen.value = e;
}
const handleAddCate = () => {
  visible.value = true;
}
const handlePromptClick = () => {
  drawerSeen.value = true;
}
</script>

<style scoped>
.prompt-container :deep(.arco-card-size-medium .arco-card-body) {
  padding: 10px;
}
</style>