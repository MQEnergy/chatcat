<template>
  <div class="prompt-container">
    <a-card :title="$t('settings.prompt')" :bordered="false" :header-style="{borderColor: 'var(--color-fill-2)'}">
      <template #extra>
        <a-button type="primary" @click="handleAddCate">
          <template #icon>
            <icon-plus/>
          </template>
          {{ $t('settings.prompt.cateAddBtn') }}
        </a-button>
      </template>
      <!-- prompt list -->
      <div class="prompt-list scrollbar">
        <a-row :gutter="[20, 20]">
          <a-col :span="12" v-for="(item, index) in promptList" :key="index">
            <prompt-item :item="item" :index="index"></prompt-item>
          </a-col>
        </a-row>
      </div>
    </a-card>
  </div>
  <!-- 添加分类 -->
  <prompt-add :visible="visible" @cancel="handleCancel"></prompt-add>
</template>

<script setup>
import {reactive, ref} from "vue";
import PromptAdd from "@views/settings/components/prompt-add.vue";
import {useRouter} from "vue-router";
import PromptItem from "@components/prompt/prompt-item.vue";
import PromptEnums from "@config/prompt.js";

const router = useRouter();
const visible = ref(false);
const drawerSeen = ref(false);
const promptList = reactive(PromptEnums)
const handleCancel = (e) => {
  visible.value = e;
  drawerSeen.value = e;
}
const handleAddCate = () => {
  visible.value = true;
}
</script>

<style scoped>
.prompt-container {
  position: relative;
}
.prompt-container :deep(.arco-card-body) {
  padding: 10px;
  padding-top: 0px;
}
.prompt-list {
  height: 700px;
  padding-top: 20px;
  overflow-y: scroll;
}
</style>