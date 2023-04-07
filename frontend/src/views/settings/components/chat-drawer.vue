<template>
  <a-drawer :width="600"
            :visible="props.visible"
            @ok="handleOk"
            @cancel="handleOk"
            :hide-cancel="true"
            :closable="true"
            unmountOnClose>
    <template #title>
      {{ $t('settings.chat.drawerTitle') }}
    </template>
    <div>
      <a-row :gutter="[10, 10]">
        <a-col :span="12" v-for="(item, index) in props.chats" :key="index">
          <a-card :title="item.name" class="prompt-card">
            <template #extra>
              <a-dropdown @select="handleSelect($event, item)">
                <a-link style="color: #000">
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
                  <a-doption :value="4">
                    <icon-drive-file/>
                    {{ $t('common.view') }}
                  </a-doption>
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
import {Message} from "@arco-design/web-vue";

const props = defineProps({
  visible: {
    type: Boolean,
    default: false
  },
  chats: {
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
      router.push('/index?chats=' + row.name)
      break
    case 2:
      Message.info("Edit")
      break;
    case 3:
      Message.info("Delete")
      break;
    case 4:
      Message.info("View")
      break;
  }
}
const handleDelete = (row, index) => {

}
</script>

<style scoped>
.prompt-card {
  border-radius: 8px;
}

.prompt-card :deep(.arco-card-header) {
  border: none;
}
</style>