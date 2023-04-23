<template>
  <a-drawer :width="600"
            :visible="props.visible"
            @ok="handleOk"
            @cancel="handleOk"
            :hide-cancel="true"
            :closable="true"
            unmountOnClose>
    <template #title>
      {{ $t('settings.chat.drawerTitle') }}
    </template>
    <div>
      <a-spin v-if="loading" tip="loading..." :style="{height: '85vh', width: '100%', textAlign: 'center'}"/>

      <a-row :gutter="[10, 10]">
        <a-col :span="12" v-for="(item, index) in chatList" :key="index">
          <a-card :title="item.name" class="prompt-card">
            <template #extra>
              <a-dropdown @select="handleSelect($event, item)">
                <a-link :style="{color: 'var(--color-text-2)'}">
                  <icon-more size="large"/>
                </a-link>
                <template #content>
                  <a-doption :value="1">
                    <icon-delete/>
                    {{ $t('common.del') }}
                  </a-doption>
                  <a-doption :value="2">
                    <icon-drive-file/>
                    {{ $t('common.view') }}
                  </a-doption>
                </template>
              </a-dropdown>
            </template>
            <a-card-meta>
              <template #avatar>
                {{ dayjs.unix(item.created_at).format('YYYY-MM-DD HH:mm:ss') }}
              </template>
            </a-card-meta>
          </a-card>
        </a-col>
      </a-row>
      <a-empty :style="{marginTop: '40px'}" v-if="chatList.length == 0"></a-empty>

    </div>
    <template #footer>
      <a-space align="start">
        <a-pagination v-if="total > 0" :total="total" :page-size="20" show-total @change="handlePageChange"/>
        <a-button type='primary' @click="handleOk">确定</a-button>
      </a-space>
    </template>
  </a-drawer>
  <!--  delete -->
  <a-modal :width="360" v-model:visible="delSeen" @ok="handleDelOk" :ok-loading="delLoading" @cancel="handleDelCancel">
    <template #title>
      {{ $t('common.notice') }}
    </template>
    <div>{{ $t('common.chat.confirmDel') }}</div>
  </a-modal>
</template>

<script setup>
import {useRouter} from "vue-router";
import {Message} from "@arco-design/web-vue";
import {DeleteChat, GetChatList} from "../../../../wailsjs/go/chat/Service.js";
import {reactive, ref} from "vue";
import dayjs from 'dayjs';

const props = defineProps({
  visible: {
    type: Boolean,
    default: false
  }
});
const router = useRouter();
const currPage = ref(1);
const total = ref(0);
let chatList = reactive([])
const cateId = ref(null);
const loading = ref(false);
const delLoading = ref(false);
const delSeen = ref(false);
const delData = ref(null);

const emits = defineEmits(['cancel']);
const handleOk = () => {
  emits('cancel', false);
}
const handleDelOk = () => {
  DeleteChat(delData.value.id).then(res => {
    if (res.code !== 0) {
      Message.error(res.msg)
      return;
    }
    initDrawChatList(cateId.value, currPage.value);
  }).finally(() => {
    delLoading.value = false;
  })
}
const handleDelCancel = () => {
  delSeen.value = false;
}
const handleSelect = (e, row) => {
  switch (e) {
    case 1:
      delData.value = row;
      delSeen.value = true;
      break;
    case 2:
      router.push('/index?chatid=' + row.id);
      break;
  }
}
const initDrawChatList = (cateid, page) => {
  loading.value = true;
  cateId.value = cateid;
  GetChatList(cateid, page).then(res => {
    if (res.code === 0) {
      currPage.value = page;
      chatList.splice(0, chatList.length);
      chatList.push(...res.data.list);
      total.value = res.data.total;
    }
  }).finally(() => {
    loading.value = false;
  })
}
const handlePageChange = (e) => {
  initDrawChatList(cateId.value, e);
}
defineExpose({
  initDrawChatList
})
</script>

<style scoped>
.prompt-card {
  border-radius: 8px;
}

.prompt-card :deep(.arco-card-header) {
  border: none;
}

.footer-pagination {
  position: absolute;
  bottom: 20px;
  left: 50%;
  transform: translateX(-50%);
}
</style>