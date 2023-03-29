<template>
  <div style="padding: 14px 10px;">
    <a-space>
      <a-input-search :style="{width:'100%'}" placeholder="请输入prompt关键词"/>
      <a-button>
        <template #icon>
          <icon-plus/>
        </template>
      </a-button>
    </a-space>
  </div>
  <a-list :bordered="false">
    <a-list-item v-for="idx in 4" :key="idx">
      <a-list-item-meta
          title="这是prompt提示词"
      >
      </a-list-item-meta>
      <template #actions>
        <icon-edit/>
        <icon-delete/>
      </template>
    </a-list-item>
  </a-list>
</template>

<script>
export default {
  name: "menu-list"
}
</script>

<style scoped>

</style>