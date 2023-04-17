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
  <a-list class="menu-list-container scrollbar" :bordered="false">
    <a-list-item v-for="(item, index) in list" :key="index">
      <a-list-item-meta style="position: relative;">
        <template #title>
          <a-input style="position: absolute; left: -8px; top: -5px;"
                   @press-enter="handleCheck(item, index)"
                   v-if="index == editIdx" v-model="item.name"
                   placeholder="请输入对话关键词"
                   allow-clear/>
          <div v-else style="width: 150px; overflow: hidden;">{{ item.name }}</div>
        </template>
      </a-list-item-meta>
      <template #actions>
        <icon-check class="check-icon" @click="handleCheck(item, index)"
                    v-if="index == editIdx"/>
        <icon-edit v-else @click="handleEdit(item, index)"/>
        <icon-close @click="handleClose(item, index)" v-if="index == editIdx"/>
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
import {defineProps, ref, toRefs} from "vue";
import {Message} from "@arco-design/web-vue";
import {DeleteChat, EditChat, SetChatData} from "../../../../wailsjs/go/chat/Service.js";

let props = defineProps({
  list: {
    type: Array,
    default: []
  },
  cateid: {
    type: Number,
    default: 0
  }
})
let {list, cateid} = toRefs(props)
const editIdx = ref(null)
const emits = defineEmits(['add:chat']);
const handleAddChat = () => {
  SetChatData({
    cate_id: props.cateid,
    name: '新的聊天',
    sort: 50,
  }).then((res) => {
    if (res.code !== 0) {
      Message.error(res.msg);
      return;
    }
    emits('add:chat', res.data);
    list.value.unshift({
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
    list.value.splice(index, 1)
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
    list.value[index].name = row.name;
  });
}
</script>

<style scoped>
.menu-search-container {
  padding: 14px 10px;
}

.menu-list-container {
  position: absolute;
  top: 63px;
  width: 100%;
  max-height: calc(100%);
  overflow-y: scroll;
}

.menu-list-container :deep(.arco-list-item-action > li:not(:last-child)) {
  margin-right: 10px;
}

.check-icon {
  color: rgb(22, 93, 255);
}

</style>