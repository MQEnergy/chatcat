<template>
  <a-drawer :width="600"
            :visible="props.visible"
            @ok="handleOk"
            @cancel="handleOk"
            :hide-cancel="true"
            :closable="true"
            unmountOnClose>
    <template #title>
      {{ $t('settings.prompt.drawerTitle') }}
    </template>
    <div>
      <a-row :gutter="[10, 10]">
        <a-col :span="12" v-for="(item, index) in props.prompts" :key="index">
          <a-card :title="item.name" class="prompt-card">
            <template #extra>
              <a-dropdown @select="handleSelect($event, item)">
                <a-link>{{ $t('common.operation') }}</a-link>
                <template #content>
                  <a-doption :value="1"><icon-plus />chat</a-doption>
                  <a-doption :value="2"><icon-edit />编辑</a-doption>
                  <a-doption :value="2"><icon-delete />删除</a-doption>
                </template>
              </a-dropdown>
            </template>
            {{ item.desc }}
          </a-card>
        </a-col>
      </a-row>
    </div>
  </a-drawer>
</template>

<script setup>
import {useRouter} from "vue-router";

const props = defineProps({
  visible: {
    type: Boolean,
    default: false
  },
  prompts: {
    type: Array,
    default: []
  }
});
const router = useRouter();

const emits = defineEmits(['cancel']);
const handleOk = () => {
  emits('cancel', false);
}
const handleSelect = (e, row) => {
  switch (e) {
    case 1:
      router.push('/index?prompt='+row.prompt)
      break
  }
}
</script>

<style scoped>
.prompt-card {
  border-radius: 4px;
}
.prompt-card :deep(.arco-card-header) {
  border: none;
}
</style>