<template>
  <a-card :title="$t('settings.updateCate')" :bordered="false" :header-style="{borderColor: 'var(--color-fill-2)'}">
    <template #extra>
      <a-button type="primary" @click="handleCheckUpdate">
        <template #icon>
          <icon-download />
        </template>
        {{ $t('settings.checkUpdate') }}
      </a-button>
    </template>
    <a-timeline :style="{marginTop: '20px'}">
      <a-timeline-item label="2017-03-10" dotColor="#00B42A">
        The first milestone
      </a-timeline-item>
      <a-timeline-item label="2018-05-22">The second milestone</a-timeline-item>
      <a-timeline-item label="2020-06-22" dotColor="#F53F3F">
        The third milestone
        <IconExclamationCircleFill
            :style="{ color: 'F53F3F', fontSize: '12px', marginLeft: '4px' }"
        />
      </a-timeline-item>
      <a-timeline-item label="2020-09-30" dotColor="#C9CDD4">
        The fourth milestone
      </a-timeline-item>
    </a-timeline>
  </a-card>
</template>

<script setup>
import {Message} from "@arco-design/web-vue";

const handleCheckUpdate = () => {
  Message.info("Checking update")
}
</script>

<style scoped>

</style>