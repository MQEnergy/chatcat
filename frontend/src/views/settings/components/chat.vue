<template>
  <div class="chat-container">
    <a-card :title="$t('settings.chat')" :bordered="false" :header-style="{borderColor: 'var(--color-fill-2)'}">
      <!-- chat list -->
      <a-space>
        <div class="chat-list scrollbar">
          <a-row :gutter="[20, 20]">
            <a-col :span="8" v-for="(item, index) in cateList" :key="index">
              <a-card hoverable :style="{ borderRadius: '8px' }">
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
      </a-space>
    </a-card>
  </div>
  <!-- chat drawer -->
  <chat-drawer :visible="drawerSeen" :chats="chatList" @cancel="handleCancel"></chat-drawer>
</template>

<script setup>
import {reactive, ref} from "vue";
import ChatDrawer from "@views/settings/components/chat-drawer.vue";

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
let chatList = reactive([])
const handleCancel = (e) => {
  visible.value = e;
  drawerSeen.value = e;
}
const handleAddCate = () => {
  visible.value = true;
}
const handlePromptClick = () => {
  chatList = [{
    id: 1,
    name: '这是chat对话语句1',
    sort: 50
  }, {
    id: 2,
    name: '这是chat对话语句2',
    sort: 50
  }, {
    id: 3,
    name: '这是chat对话语句3',
    sort: 50
  }, {
    id: 1,
    name: '这是chat对话语句1',
    sort: 50
  }, {
    id: 2,
    name: '这是chat对话语句2',
    sort: 50
  }];
  drawerSeen.value = true;
}
const handleSelect = (e, row) => {
  switch (e) {
    case 1:
      handlePromptClick()
      break;
    case 2:
      handlePromptClick()
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
  height: 700px;
  padding-top: 10px;
  overflow-y: scroll;
}
</style>