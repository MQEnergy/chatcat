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
          <a-link :style="{color: 'var(--color-text-2)'}">
            <icon-more size="large"/>
          </a-link>
          <template #content>
            <a-doption :value="1">
              <icon-send/>
              {{ $t('common.chat') }}
            </a-doption>
            <a-doption :value="2">
              <icon-edit/>
              {{ $t('common.edit') }}
            </a-doption>
            <a-doption :value="3">
              <icon-delete/>
              {{ $t('common.del') }}
            </a-doption>
          </template>
        </a-dropdown>
      </template>
      <a-space direction="vertical">
        {{ item.desc }}
        <div style="height: 160px; overflow: hidden; overflow-y: scroll">
          <a-typography-text>
            {{ promptId == item.id ? item.enprompt : item.prompt }}
          </a-typography-text>
        </div>
      </a-space>
      <a-card-meta v-if="item.tag_name" :style="{ position: 'absolute', bottom: '10px', left: '10px' }">
        <template #description>
          <a-space>
            <a-tag v-for="(val, idx) in item.tag_name.split(',')" :key="idx" size="small" bordered>{{ val }}</a-tag>
          </a-space>
        </template>
      </a-card-meta>
      <a-space :style="{ position: 'absolute', bottom: '10px', right: '10px' }">
        <a-switch size="small" unchecked-color="#14C9C9" @change="handleSwitch($event, item)">
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
import {ref} from "vue";

const promptId = ref(null)
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
const emits = defineEmits(['del', 'edit']);

const handleSelect = (e, row) => {
  switch (e) {
    case 1:
      router.push('/index?prompt=' + row.prompt);
      break;
    case 2:
      emits('edit', row);
      break;
    case 3:
      emits('del', row);
      break;
  }
}
const handleSwitch = (e, row) => {
  if (e) {
    promptId.value = row.id;
  } else {
    promptId.value = null;
  }
}
</script>

<style scoped>

.prompt-card {
  border-radius: 8px;
  height: 280px;
  cursor: pointer;
}

.prompt-card :deep(.arco-card-header) {
  border: none;
  padding: 10px;
}

</style>