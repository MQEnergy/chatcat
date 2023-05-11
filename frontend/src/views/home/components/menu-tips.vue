<template>
  <div class="menu-container">
    <a-trigger
        :trigger="['click']"
        clickToClose
        position="top"
        v-model:popupVisible="popupVisible"
    >
      <div :class="`button-trigger ${popupVisible ? 'button-trigger-active' : ''}`">
        <IconClose v-if="popupVisible"/>
        <IconMessage v-else/>
      </div>
      <template #content>
        <a-menu
            :style="{ marginBottom: '-4px' }"
            mode="popButton"
            :collapsed="true"
            :tooltipProps="{ position: 'left' }"
            showCollapseButton
            @menu-item-click="handleMenuClick"
        >
          <a-menu-item key="1" style="border: 1px solid #f5f5f5;">
            <template #icon>
              <IconBug></IconBug>
            </template>
            {{ $t('settings.contact') }}
          </a-menu-item>
          <a-trigger trigger="click" position="lb" auto-fit-position :unmount-on-close="true">
            <a-menu-item key="2" style="border: 1px solid #f5f5f5;">
              <template #icon>
                <IconBulb></IconBulb>
              </template>
              {{ $t('settings.prompt') }}
            </a-menu-item>
            <template #content>
              <div class="prompt-tips-container">
                <a-card class="prompt-tips-card scrollbar"
                        :body-style="{height: '340px', overflowY: 'scroll'}"
                        :title="$t('settings.prompt.common')">
                  <a-space direction="vertical">
                    <prompt-simple-item v-for="(item, index) in promptList"
                                        :item="item" :index="index" @add:prompt="handlePromptToChat">
                    </prompt-simple-item>
                  </a-space>
                </a-card>
              </div>
            </template>
          </a-trigger>
        </a-menu>
      </template>
    </a-trigger>
  </div>
</template>

<script setup>
import {IconBug, IconBulb, IconClose, IconMessage} from '@arco-design/web-vue/es/icon';
import {onMounted, reactive, ref} from "vue";
import {useRouter} from "vue-router";
import PromptSimpleItem from "@components/prompt/prompt-simple-item.vue";
import {GetNormalPromptList} from "../../../../wailsjs/go/prompt/Service.js";

const router = useRouter();
const position = ref('right');
const popupVisible = ref(false);
const emits = defineEmits(['add:prompt'])
const promptList = reactive([]);

const handleMenuClick = (e) => {
  switch (e) {
    case "1": // contact
      router.push('/settings/index?tab=5')
      break;
    case "2": // prompt
      break;
  }
}
const handlePromptToChat = (row) => {
  emits('add:prompt', row)
}
const initPromptList = async () => {
  const res = await GetNormalPromptList();
  promptList.splice(0, promptList.length, ...res.data);
}
onMounted(async () => {
  await initPromptList();
})
</script>

<style scoped>
.menu-container {
}

.button-trigger {
  position: absolute;
  bottom: 180px;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 40px;
  height: 40px;
  color: var(--color-white);
  font-size: 14px;
  border-radius: 50%;
  cursor: pointer;
  transition: all 0.1s;
  right: 10px;
  background-color: rgb(var(--purple-5));
}

.button-trigger .button-trigger-active {
  background-color: var(--color-primary-light-4);
}

.prompt-tips-container {
  width: 310px;
  height: 400px;
  position: absolute;
  right: -40px;
  bottom: 0px;
  overflow: hidden;
}

.prompt-tips-container :deep(.arco-card-header) {
  border: none;
}

.prompt-tips-container :deep(.arco-card-body) {
  padding: 0px 10px 10px;
}

.prompt-tips-card {
  border-radius: 8px;
  height: 398px;
}
</style>