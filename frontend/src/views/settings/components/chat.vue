<template>
  <div class="chat-container">
    <a-card :title="$t('settings.chat')" :bordered="false" :header-style="{borderColor: 'var(--color-fill-2)'}">
      <!-- chat list -->
      <div class="chat-list scrollbar">
        <a-row :gutter="[20, 20]">
          <a-col :span="8">
            <a-card hoverable :style="{ borderRadius: '8px' }">
              <div :style="{
                  display: 'flex',
                  alignItems: 'center',
                  justifyContent: 'space-between',
                }">
                  <span :style="{ display: 'flex', alignItems: 'center', color: '#1D2129' }">
                    <a-avatar :style="{ marginRight: '8px', backgroundColor: '#999' }" :size="28">
                      W
                    </a-avatar>
                    <a-typography-text>未分类</a-typography-text>
                  </span>
                <a-dropdown @select="handleSelect($event, {id:0})">
                  <a-link style="color: #000">
                    <icon-more size="large"/>
                  </a-link>
                  <template #content>
                    <a-doption :value="1">
                      <icon-copy/>
                      {{ $t('common.view') }}
                    </a-doption>
                  </template>
                </a-dropdown>
              </div>
            </a-card>
          </a-col>
          <a-col :span="8" v-for="(item, index) in cateList" :key="index">
            <a-card hoverable :style="{ borderRadius: '8px' }">
              <div :style="{
                  display: 'flex',
                  alignItems: 'center',
                  justifyContent: 'space-between',
                }">
                  <span :style="{ display: 'flex', alignItems: 'center', color: '#1D2129' }">
                    <a-avatar :style="{ marginRight: '8px', backgroundColor: item.color }" :size="28">
                      {{ item.letter }}
                    </a-avatar>
                    <a-typography-text>{{ item.name }}</a-typography-text>
                  </span>
                <a-dropdown @select="handleSelect($event, item)">
                  <a-link style="color: #000">
                    <icon-more size="large"/>
                  </a-link>
                  <template #content>
                    <a-doption :value="1">
                      <icon-copy/>
                      {{ $t('common.view') }}
                    </a-doption>
                    <a-doption :value="2">
                      <icon-edit/>
                      {{ $t('common.edit') }}
                    </a-doption>
                    <a-doption :value="3">
                      <icon-delete/>
                      {{ $t('common.del') }}
                    </a-doption>
                  </template>
                </a-dropdown>
              </div>
            </a-card>
          </a-col>
        </a-row>
      </div>
    </a-card>
    <div style="width: 400px; margin: 0 auto;">
      <a-pagination :total="50"/>
    </div>
  </div>
  <!-- chat drawer -->
  <chat-drawer ref="chatRecordListRef" :visible="drawerSeen" @cancel="handleCancel"></chat-drawer>
</template>

<script setup>
import {onMounted, reactive, ref} from "vue";
import ChatDrawer from "@views/settings/components/chat-drawer.vue";
import {GetChatCateList} from "../../../../wailsjs/go/chat/Service.js";

const visible = ref(false);
const drawerSeen = ref(false);
const cateList = reactive([]);
const currPage = ref(1);
const chatRecordListRef = ref(null);

const initChatCateList = (page) => {
  GetChatCateList(page).then(res => {
    cateList.push(...res.data.list);
  })
}
onMounted(() => {
  initChatCateList(currPage.value);
})

const handleCancel = (e) => {
  visible.value = e;
  drawerSeen.value = e;
}
const handlePromptClick = (row) => {
  drawerSeen.value = true;
  chatRecordListRef.value.initDrawChatList(row.id, 1);
}
const handleSelect = (e, row) => {
  switch (e) {
    case 1:
      handlePromptClick(row)
      break;
    case 2:
      handlePromptClick(row)
      break;
  }
}
const handleDelete = (row, index) => {

}
</script>

<style scoped>

.chat-container {
  position: relative;
}

.chat-container :deep(.arco-card-body) {
  /*padding: 10px;*/
  /*padding-top: 0px;*/
}

.chat-list {
  width: 100%;
  height: 700px;
  padding-top: 10px;
  overflow-y: scroll;
}
</style>