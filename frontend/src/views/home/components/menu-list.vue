<template>
  <div class="menu-search-container" style="">
    <a-space>
      <a-input-search v-model="keyword" :style="{width:'100%'}" :placeholder="$t('common.keyword.query')"
                      @press-enter="handleSearchChat"/>
      <a-button @click="handleAddChat">
        <template #icon>
          <icon-plus/>
        </template>
      </a-button>
    </a-space>
  </div>
  <a-list ref="chatlistRef" class="menu-list-container scrollbar" size="small"
          :bordered="false" :max-height="980" @reach-bottom="fetchData" :scrollbar="false" hoverable>
    <template v-if="!bottom" #scroll-loading>
      <a-spin/>
    </template>
    <a-list-item v-if="chatList.length > 0" class="menu-list-item flash" v-for="(item, index) in chatList"
                 @click="handleSelectChat(item, index)"
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
import {defineProps, reactive, ref, watch} from "vue";
import {DeleteChat, EditChat, GetChatList, SearchChatList, SetChatData} from "../../../../wailsjs/go/chat/Service.js";
import {Message} from "@arco-design/web-vue";
import {useI18n} from "vue-i18n";

const {t} = useI18n();

const props = defineProps({
  cateid: {
    type: Number,
    default: 0
  },
  chatid: {
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
const bottom = ref(false);
const chatlistRef = ref(null);
const keyword = ref('');
const isSearch = ref(false);
const currChatId = ref(0);
const currCateId = ref(0);

const initChatList = async (cateid, page, type) => {
  if (window.go === undefined) {
    return;
  }
  loading.value = true;
  isSearch.value = false;
  let res = await GetChatList(cateid, page);
  if (res.code === 0) {
    if (type === 1) {
      chatList.splice(0, chatList.length);
    }
    chatList.push(...res.data.list);
    if (chatList.length > 0) {
      curPage.value = res.data.current_page + 1;
      if (currChatId.value === 0) {
        currIdx.value = 0;
      } else {
        chatList.forEach((item, index) => {
          if (item.id === currChatId.value) {
            currIdx.value = index;
          }
        })
      }
      let chatName = chatList[0].name;
      currChatId.value = props.chatid || chatList[0].id;
      emits('header:info', {
        chatName: chatName,
        chatId: currChatId.value,
        cateId: currCateId.value,
      });
    }
  }
  loading.value = false;
  bottom.value = true;
}
const fetchData = async () => {
  if (isSearch.value) {
    searchChatList(keyword.value, curPage.value, 2);
  } else {
    await initChatList(props.cateid, curPage.value, 2);
  }
}

watch(() => props.cateid, () => {
  currCateId.value = props.cateid;
  initChatList(props.cateid, 1, 1);
})
watch(() => props.chatid, () => {
  currChatId.value = props.chatid;
  bottom.value = true;
})

const handleAddChat = () => {
  SetChatData({
    cate_id: props.cateid,
    name: t('common.newchat'),
    sort: 50,
  }).then((res) => {
    if (res.code !== 0) {
      Message.error(res.msg);
      return;
    }
    chatList.unshift({
      id: res.data.id,
      name: res.data.name,
      sort: res.data.sort
    })
    currIdx.value = 0;
    emits('new:chat', res.data);
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
    initChatList(props.cateid, 1, 1);
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
const searchChatList = (keyword, page, type) => {
  isSearch.value = true;
  SearchChatList(keyword, page).then(res => {
    if (res.code !== 0) {
      return;
    }
    if (res.data.list.length === 0) {
      bottom.value = true;
    }
    if (type === 1) {
      chatList.splice(0, chatList.length);
    }
    chatList.push(...res.data.list);
    curPage.value = res.data.current_page + 1;

    let chatName = '';
    if (chatList.length > 0 && currIdx.value === 0) {
      chatName = chatList[0].name;
      currChatId.value = chatList[0].id;
      currCateId.value = props.cateid;
      emits('header:info', {
        chatName: chatName,
        chatId: currChatId,
        cateId: currCateId,
      });
    }
  }).finally(() => {
    loading.value = false;
    bottom.value = true;
  })
}
const handleSearchChat = () => {
  if (keyword.value.trim() === "") {
    initChatList(currCateId.value, 1, 1);
    return
  }
  searchChatList(keyword.value.trim(), 1, 1);
}
</script>

<style scoped>
.menu-search-container {
  padding: 10px;
  border-bottom: 1px solid var(--color-neutral-3);
  background: var(--color-bg-2);
}

.menu-list-container {
  position: absolute;
  max-height: calc(100% - 70px);
  overflow-y: scroll;
  cursor: pointer;
  width: 224px;
  padding: 8px;
}

.menu-list-container :deep(.arco-list::-webkit-scrollbar) {
  display: none;
}

.menu-list-container .menu-list-item {
  height: 49px;
  margin-bottom: 10px;
  border-radius: 4px;
  border: 1px solid var(--color-neutral-3);
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
  background: var(--color-bg-2);
}

</style>