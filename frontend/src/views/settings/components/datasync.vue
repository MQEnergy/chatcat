<template>
  <div>
    <!-- 关联账号 -->
    <a-card title="关联账号" :bordered="false" :header-style="{borderColor: 'var(--color-fill-2)'}">
      <a-space size="medium">
        <a-button type="primary">登录</a-button>
        <a-button>暂无账号，请去注册 <icon-arrow-right /></a-button>
      </a-space>
    </a-card>
    <!-- 开启同步 -->
    <a-card title="开启同步" :bordered="false" :header-style="{borderColor: 'var(--color-fill-2)'}">
      <a-space size="medium">
        <a-switch></a-switch>
        5s自动同步一次
      </a-space>
    </a-card>
  </div>
</template>

<script setup>
</script>

<style scoped>

</style>