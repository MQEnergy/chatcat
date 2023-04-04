<template>
  <div>
    <a-card :title="$t('settings.datasync.linkAccount')" :bordered="false" :header-style="{borderColor: 'var(--color-fill-2)'}">
      <a-space direction="horizontal" size="medium" style="margin-bottom: 16px;">
        <a-button type="primary">{{ $t('settings.datasync.loginBtn') }}</a-button>
        <a-button>{{ $t('settings.datasync.registerBtn') }}  <icon-arrow-right /></a-button>
      </a-space>
      <a-alert type="warning">
        <a-typography-text>{{ $t('settings.datasync.accountTips') }}</a-typography-text>
      </a-alert>
    </a-card>
    <a-card :title="$t('settings.datasync.beginSync')" :bordered="false" :header-style="{borderColor: 'var(--color-fill-2)'}">
      <a-space size="medium">
        <a-switch></a-switch>
        {{ $t('settings.datasync.syncTips') }}
      </a-space>
    </a-card>
  </div>
</template>

<script setup>
</script>

<style scoped>

</style>