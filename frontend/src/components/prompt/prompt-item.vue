<template>
  <div>
    <a-card hoverable class="prompt-card">
      <template #title>
        <a-space>
          {{ item.name }}
        </a-space>
      </template>
      <template #extra>
        <a-dropdown @select="handleSelect($event, item)">
          <a-link>{{ $t('common.operation') }}</a-link>
          <template #content>
            <a-doption :value="1"><icon-send /> {{ $t('common.chat') }}</a-doption>
            <a-doption :value="2"><icon-edit /> {{ $t('common.edit') }}</a-doption>
            <a-doption :value="3"><icon-delete /> {{ $t('common.del') }}</a-doption>
            <a-doption :value="4"><icon-copy /> {{ $t('common.copy') }}</a-doption>
          </template>
        </a-dropdown>
      </template>
      <a-space direction="vertical">
        {{ item.desc }}
        <a-typography-text>
          {{ item.prompt }}
        </a-typography-text>
      </a-space>
      <a-card-meta :style="{ position: 'absolute', bottom: '10px', left: '10px' }">
        <template #description>
          <a-space>
            <a-tag size="small" bordered>常用</a-tag>
            <a-tag size="small" bordered>写作辅助</a-tag>
            <a-tag size="small" bordered>AI</a-tag>
          </a-space>
        </template>
      </a-card-meta>
      <a-space :style="{ position: 'absolute', bottom: '10px', right: '10px' }">
        <a-switch size="small" unchecked-color="#14C9C9">
          <template #checked>
            CN
          </template>
          <template #unchecked>
            EN
          </template>
        </a-switch>
      </a-space>
    </a-card>
  </div>
</template>

<script setup>
import {useRouter} from "vue-router";

const router = useRouter();
const props = defineProps({
  item: {
    type: Object,
    default: {}
  },
  index: {
    type: Number,
    default: 0
  }
})

const handleSelect = (e, row) => {
  switch (e) {
    case 1:
      router.push('/index?prompt='+row.prompt);
      break;
    case 2:
      break;
  }
}
</script>

<style scoped>

.prompt-card {
  border-radius: 8px;
  height: 260px;
  cursor: pointer;
}
.prompt-card :deep(.arco-card-header) {
  border: none;
}
</style>