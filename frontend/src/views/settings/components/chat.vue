<template>
  <div class="chat-container">
    <a-card :title="$t('settings.chat')" :bordered="false" :header-style="{borderColor: 'var(--color-fill-2)'}">
      <!-- chat list -->
      <div class="chat-list scrollbar">
        <a-spin v-if="loading" tip="loading..." :style="{height: '80vh', width: '100%', textAlign: 'center'}"/>
        <a-row :gutter="[20, 20]">
          <a-col :span="8">
            <a-card hoverable :style="{ borderRadius: '8px', cursor: 'pointer' }">
              <div :style="{
                  display: 'flex',
                  alignItems: 'center',
                  justifyContent: 'space-between',
                }">
                  <span :style="{ display: 'flex', alignItems: 'center', color: '#1D2129', width: '80%' }"
                        @click="handlePromptClick({id:0})">
                    <a-avatar :style="{ marginRight: '8px', backgroundColor: '#999' }" :size="28">
                      W
                    </a-avatar>
                    <a-typography-text>{{ $t('common.nocate') }}</a-typography-text>
                  </span>
                <a-dropdown @select="handleSelect($event, {id:0})">
                  <a-link :style="{color: 'var(--color-text-2)'}">
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
            <a-card hoverable :style="{ borderRadius: '8px', cursor: 'pointer' }">
              <div :style="{ display: 'flex', alignItems: 'center', justifyContent: 'space-between'}">
                  <span :style="{ display: 'flex', alignItems: 'center', color: '#1D2129', width: '80%' }"
                        @click="handlePromptClick(item)">
                    <a-avatar :style="{ marginRight: '8px', backgroundColor: item.color }" :size="28">
                      {{ item.letter }}
                    </a-avatar>
                    <a-typography-text>{{ item.name }}</a-typography-text>
                  </span>
                <a-dropdown @select="handleSelect($event, item)">
                  <a-link :style="{color: 'var(--color-text-2)'}">
                    <icon-more size="large"/>
                  </a-link>
                  <template #content>
                    <a-doption :value="1">
                      <icon-edit/>
                      {{ $t('common.edit') }}
                    </a-doption>
                    <a-doption :value="2">
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
    <div class="footer-pagination" v-if="total > 0">
      <a-pagination :total="total" :page-size="21" show-total @change="handlePageChange"/>
    </div>
  </div>
  <!-- chat drawer -->
  <chat-drawer ref="chatRecordListRef" :visible="drawerSeen" @cancel="handleCancel"></chat-drawer>
  <!-- cate edit drawer -->
  <cate-add ref="cateAddRef" :visible="cateSeen" :form-data="formData" @cancel="handleCancel"
            @ok="handleCateOk"></cate-add>
  <!--  delete -->
  <a-modal :width="360" v-model:visible="delSeen" @ok="handleDelOk" :ok-loading="delLoading" @cancel="handleDelCancel">
    <template #title>
      {{ $t('common.notice') }}
    </template>
    <div>{{ $t('common.chat.confirmDel') }}</div>
  </a-modal>
</template>

<script setup>
import {onMounted, reactive, ref} from "vue";
import ChatDrawer from "@views/settings/components/chat-drawer.vue";
import {DeleteChatCate, GetChatCateList, UpdateChatCateData} from "../../../../wailsjs/go/chat/Service.js";
import CateAdd from "../../home/components/cate-add.vue";
import {Message} from "@arco-design/web-vue";

const drawerSeen = ref(false);
const cateSeen = ref(false);
const formData = ref(null);
const cateList = reactive([]);
const currPage = ref(1);
const total = ref(0);
const delSeen = ref(false);
const delData = ref(null);
const delLoading = ref(false);

const chatRecordListRef = ref(null);
const cateAddRef = ref(null);
const loading = ref(false);
const initChatCateList = (page) => {
  loading.value = true;
  GetChatCateList(page).then(res => {
    if (res.code === 0) {
      currPage.value = page;
      cateList.splice(0, cateList.length);
      cateList.push(...res.data.list);
      total.value = res.data.total;
    }
  }).finally(() => {
    loading.value = false;
  })
}
onMounted(() => {
  if (window.go !== undefined) {
    initChatCateList(currPage.value);
  }
})

const handleCancel = (e) => {
  drawerSeen.value = e;
  cateSeen.value = e;
}
const handleCateOk = (form) => {
  UpdateChatCateData(form).then((res) => {
    if (res.code !== 0) {
      Message.error(res.msg);
      return;
    }
    initChatCateList(currPage.value);
    cateSeen.value = false;
  })
}
const handlePromptClick = (row) => {
  drawerSeen.value = true;
  chatRecordListRef.value.initDrawChatList(row.id, 1);
}
const handleCateAdd = (row) => {
  cateSeen.value = true;
  formData.value = row;
}
const handleCateDel = (row) => {
  delData.value = row;
  delSeen.value = true;
}
const handleSelect = (e, row) => {
  switch (e) {
    case 1: // edit
      handleCateAdd(row);
      break;
    case 2: // del
      handleCateDel(row);
      break;
  }
}
const handlePageChange = (e) => {
  initChatCateList(e);
}
const handleDelOk = () => {
  delLoading.value = true;
  DeleteChatCate(delData.value.id).then(res => {
    if (res.code !== 0) {
      Message.error(res.msg)
      return;
    }
    initChatCateList(currPage.value);
  }).finally(() => {
    delLoading.value = false;
  })
}
const handleDelCancel = () => {
  delSeen.value = false;
}
</script>

<style scoped>
.chat-container {
  position: relative;
  width: 100%;
  height: 90vh;
}

.chat-list {
  width: 100%;
  height: 80vh;
  overflow-y: scroll;
  overflow-x: hidden;
}

.footer-pagination {
  position: absolute;
  bottom: 20px;
  left: 50%;
  transform: translateX(-50%);
}
</style>