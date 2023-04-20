<template>
  <div class="menu-search-container" style="">
    <a-space>
      <a-input-search :style="{width:'100%'}" :placeholder="$t('common.keyword.query')"/>
      <a-button @click="handleAddChat">
        <template #icon>
          <icon-plus/>
        </template>
      </a-button>
    </a-space>
  </div>
  <a-list class="menu-list-container scrollbar" size="small"
          :bordered="false" hoverable :loading="loading">
    <a-list-item class="menu-list-item" v-for="(item, index) in chatList" @click="handleSelectChat(item, index)"
                 :class="{'selected': currIdx === index}" :key="index">
      <a-list-item-meta style="position: relative;">
        <template #title>
          <div v-if="index === editIdx">
            <a-input size="medium" class="menu-input-item"
                     @press-enter="handleCheck(item, index)"
                     v-model="item.name"
                     allow-clear/>
          </div>
          <div v-else class="menu-text-item">{{ item.name }}</div>
        </template>
      </a-list-item-meta>
      <template #actions>
        <icon-check class="check-icon" @click="handleCheck(item, index)"
                    v-if="index === editIdx"/>
        <icon-edit v-else @click="handleEdit(item, index)"/>
        <icon-close @click="handleClose(item, index)" v-if="index === editIdx"/>
        <a-popconfirm v-else
                      :cancel-text="$t('common.cancel')"
                      :ok-text="$t('common.ok')"
                      :content="$t('common.confirmDel')"
                      @ok="handleDelete(item, index)">
          <icon-delete/>
        </a-popconfirm>
      </template>
    </a-list-item>
  </a-list>
</template>

<script setup>
import {defineProps, onMounted, reactive, ref, watch} from "vue";
import {DeleteChat, EditChat, GetChatList, SetChatData} from "../../../../wailsjs/go/chat/Service.js";
import {Message} from "@arco-design/web-vue";
import {useI18n} from "vue-i18n";

const {t} = useI18n();

let props = defineProps({
  cateid: {
    type: Number,
    default: 0
  }
})
let chatList = reactive([]);
const editIdx = ref(null);
const currIdx = ref(0);
const loading = ref(false);
const curPage = ref(1);
const emits = defineEmits(['new:chat', 'select:chat', 'header:info']);

const initChatList = (cateid, page) => {
  GetChatList(cateid, page).then(res => {
    chatList.splice(0, chatList.length);
    chatList.push(...res.data.list);
    let chatName = '';
    if (res.data.list.length > 0) {
      chatName = res.data.list[0].name;
    }
    emits('header:info', {
      chatName: chatName
    })
  })
}
watch(() => props.cateid, () => {
  initChatList(props.cateid, curPage.value);
})
onMounted(() => {
  if (window.go !== undefined) {
    initChatList(props.cateid, curPage.value);
  }
})
const handleAddChat = () => {
  let num = chatList.length + 1
  SetChatData({
    cate_id: props.cateid,
    name: t('common.newchat') + num,
    sort: 50,
  }).then((res) => {
    if (res.code !== 0) {
      Message.error(res.msg);
      return;
    }
    emits('new:chat', res.data);
    chatList.unshift({
      id: res.data.id,
      name: res.data.name,
      sort: res.data.sort
    })
  })
}
const handleEdit = (row, index) => {
  editIdx.value = index;
}
const handleDelete = (row, index) => {
  DeleteChat(row.id).then((res) => {
    if (res.code !== 0) {
      Message.error(res.msg)
      return;
    }
    chatList.splice(index, 1)
  })
}
const handleClose = (row, index) => {
  editIdx.value = null
  if (row.name == '') {
    handleDelete(row, index)
  }
}
const handleCheck = (row, index) => {
  if (row.name.trim() === "") {
    return
  }
  EditChat(row.id, row.name).then(res => {
    if (res.code !== 0) {
      Message.error(res.msg);
      return;
    }
    editIdx.value = null
    chatList[index].name = row.name;
  });
}
const handleSelectChat = (row, index) => {
  currIdx.value = index;
  emits('select:chat', row);
}
</script>

<style scoped>
.menu-search-container {
  padding: 10px;
}

.menu-list-container {
  position: absolute;
  max-height: calc(100%);
  overflow-y: scroll;
  cursor: pointer;
  width: 220px;
  padding: 0 10px;
}

.menu-list-container .menu-list-item {
  height: 49px;
  margin-bottom: 10px;
  border-radius: 4px;
}

.menu-list-container .menu-list-item .menu-text-item {
  width: 130px;
  overflow: hidden;
}

.menu-list-container .menu-list-item .menu-input-item {
  position: absolute;
  left: -8px;
  top: 0px;
}

.menu-list-container :deep(.arco-list-item-action > li:not(:last-child)) {
  margin-right: 10px;
}

.check-icon {
  color: rgb(22, 93, 255);
}

.selected {
  background: var(--color-bg-4);
  /*box-sizing: border-box;*/
  /*border: 1px solid rgb(22, 93, 255) !important;*/
}

</style>