<template>
  <div>
    <a-card hoverable class="prompt-card">
      <template #title>
        <a-space>
          {{ item.name }}
        </a-space>
      </template>
      <template #extra>
        <a-link @click="handleAddChat(item)">
          <icon-send/>
          {{ $t('common.chat') }}
        </a-link>
      </template>
      <a-space direction="vertical">
        {{ item.desc }}
      </a-space>
    </a-card>
  </div>
</template>

<script setup>
import {useRouter} from "vue-router";

const router = useRouter();
const emits = defineEmits(['add:prompt'])
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
const handleAddChat = (row) => {
  emits('add:prompt', row)
}
</script>

<style scoped>

.prompt-card {
  border-radius: 8px;
  height: 100px;
  cursor: pointer;
}

.prompt-card :deep(.arco-card-header) {
  border: none;
  padding: 10px;
}

</style>