<template>
  <div style="padding: 14px 10px;">
    <a-space>
      <a-input-search :style="{width:'100%'}" placeholder="请输入提示词搜索"/>
      <a-dropdown @select="handleSelect">
        <a-button>
          <template #icon>
            <icon-plus/>
          </template>
        </a-button>
        <template #content>
          <a-doption :value="1">分组</a-doption>
          <a-doption :value="2">对话</a-doption>
        </template>
      </a-dropdown>

    </a-space>
  </div>
  <a-list :bordered="false">
    <a-list-item v-for="(item, index) in list" :key="index">
      <a-list-item-meta style="position: relative;">
        <template #title>
          <a-input style="position: absolute; left: -8px; top: -5px;" v-if="index == editIdx" v-model="item.name"
                   allow-clear/>
          <div v-else>{{ item.name }}</div>
        </template>
      </a-list-item-meta>
      <template #actions>
        <icon-check class="check-icon" @click="handleCheck(item, index)" v-if="index == editIdx"/>
        <icon-edit v-else @click="handleEdit(item, index)"/>
        <icon-close @click="handleClose" v-if="index == editIdx"/>
        <a-popconfirm v-else :content="$t('common.confirmDel')" @ok="handleDelete(item, index)">
          <icon-delete/>
        </a-popconfirm>
      </template>
    </a-list-item>
  </a-list>
</template>

<script setup>
import {defineProps, ref, toRefs} from "vue";

let props = defineProps({
  list: {
    type: Array,
    default: []
  }
})
const {list} = toRefs(props)

const editIdx = ref(null)
const handleSelect = (e) => {
  console.log(e)
}
const handleEdit = (row, index) => {
  editIdx.value = index
}
const handleDelete = (row, index) => {
  list.splice(index, 1)
}
const handleClose = () => {
  editIdx.value = null
}
const handleCheck = (row, index) => {
  if (row.name.trim() === "") {
    return
  }
  editIdx.value = null
  list[index].name = row.name;
}
</script>

<style scoped>
.check-icon {
  color: rgb(22, 93, 255);
}
</style>