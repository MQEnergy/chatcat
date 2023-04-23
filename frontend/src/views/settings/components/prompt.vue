<template>
  <div class="prompt-container">
    <a-card :title="$t('settings.prompt')" :bordered="false" :header-style="{borderColor: 'var(--color-fill-2)'}">
      <template #extra>
        <a-button type="primary" @click="handleAddCate">
          <template #icon>
            <icon-plus/>
          </template>
          {{ $t('settings.prompt.addBtn') }}
        </a-button>
      </template>
      <!-- prompt list -->
      <div class="prompt-list scrollbar">
        <a-spin v-if="loading" tip="loading..." :style="{height: '80vh', width: '100%', textAlign: 'center'}"/>
        <a-row :gutter="[20, 20]">
          <a-col :span="12" v-for="(item, index) in promptList" :key="index">
            <prompt-item :item="item" :index="index" @del="handleDelPrompt" @edit="handleEditPrompt"></prompt-item>
          </a-col>
        </a-row>
      </div>
    </a-card>
    <div class="footer-pagination" v-if="total > 0">
      <a-pagination :total="total" :page-size="21" show-total @change="handlePageChange"/>
    </div>
  </div>
  <!-- add prompt -->
  <prompt-add :visible="visible" :form-data="formData" @cancel="handleCancel" @ok="handleAddOk"></prompt-add>
  <!--  delete -->
  <a-modal :width="360" v-model:visible="delSeen" @ok="handleDelOk" :ok-loading="delLoading" @cancel="handleDelCancel">
    <template #title>
      {{ $t('common.notice') }}
    </template>
    <div>{{ $t('common.prompt.confirmDel') }}</div>
  </a-modal>
</template>

<script setup>
import {nextTick, onMounted, reactive, ref} from "vue";
import PromptAdd from "@views/settings/components/prompt-add.vue";
import {useRouter} from "vue-router";
import PromptItem from "@components/prompt/prompt-item.vue";
import {DeletePrompt, GetPromptList, SetPromptData, UpdatePromptData} from "../../../../wailsjs/go/prompt/Service.js";
import {Message} from "@arco-design/web-vue";

const router = useRouter();
const visible = ref(false);
const drawerSeen = ref(false);
const promptList = reactive([])
const currPage = ref(1);
const total = ref(0);
const loading = ref(false);
const delSeen = ref(false);
const delData = ref(null);
const delLoading = ref(false);
const formData = ref(null);

const initPromptList = (page) => {
  loading.value = true;
  GetPromptList(page).then(res => {
    if (res.code === 0) {
      currPage.value = page;
      promptList.splice(0, promptList.length);
      promptList.push(...res.data.list);
      total.value = res.data.total;
    }
  }).finally(() => {
    loading.value = false;
  })
}
onMounted(() => {
  initPromptList(currPage.value);
});
const handlePageChange = (page) => {
  currPage.value = page;
  initPromptList(page);
}
const handleCancel = (e) => {
  visible.value = e;
  drawerSeen.value = e;
}
const handleAddOk = (form) => {
  if (form.id) {
    UpdatePromptData(form).then(res => {
      if (res.code !== 0) {
        Message.error(res.msg)
        return;
      }
      visible.value = false;
    }).finally(() => {
      nextTick(() => {
        initPromptList(currPage.value);
      })
    });
  } else {
    SetPromptData(form).then(res => {
      if (res.code !== 0) {
        Message.error(res.msg)
        return;
      }
      visible.value = false;
    }).finally(() => {
      nextTick(() => {
        initPromptList(currPage.value);
      })
    })
  }
}
const handleAddCate = () => {
  formData.value = null;
  visible.value = true;
}
const handleDelPrompt = (e) => {
  delSeen.value = true;
  delData.value = e;
}
const handleEditPrompt = (e) => {
  visible.value = true;
  formData.value = e;
}
const handleDelCancel = () => {
  delSeen.value = false;
}
const handleDelOk = () => {
  loading.value = true;
  DeletePrompt(delData.value.id).then(res => {
    if (res.code !== 0) {
      Message.error(res.msg)
      return;
    }
    initPromptList(currPage.value);
  }).finally(() => {
    delLoading.value = false;
  })
}
</script>

<style scoped>
.prompt-container {
  position: relative;
}

.prompt-container :deep(.arco-card-body) {
  padding: 0px 10px 10px;
}

.prompt-list {
  height: 84vh;
  padding-top: 20px;
  overflow-x: hidden;
  overflow-y: scroll;
}

.prompt-list :deep(.arco-row) {
  margin-bottom: 40px !important;
}

.footer-pagination {
  position: absolute;
  bottom: 20px;
  left: 50%;
  transform: translateX(-50%);
  background: var(--color-bg-2);
  padding: 10px;
  border-radius: 6px;
}
</style>