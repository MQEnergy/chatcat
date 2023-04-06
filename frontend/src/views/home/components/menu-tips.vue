<template>
  <div class="menu-container">
    <a-trigger
        :trigger="['click', 'hover']"
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
                <a-card :style="{ width: '100%', borderRadius: '8px', height: '400px' }" title="这是prompt热门提示词">
                  <div :style="{ height: '400px', overflow: 'hidden', overflowY: 'scroll' }">
                    <a-list>
                      <a-list-item v-for="idx in 4" :key="idx">
                        <a-list-item-meta
                            title="Beijing Bytedance Technology Co., Ltd."
                            description="Beijing ByteDance Technology Co., Ltd. is an enterprise located in China."
                        >
                          <template #avatar>
                            <a-avatar shape="square">
                              <img
                                  alt="avatar"
                                  src="https://p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/3ee5f13fb09879ecb5185e440cef6eb9.png~tplv-uwbnlip3yd-webp.webp"
                              />
                            </a-avatar>
                          </template>
                        </a-list-item-meta>
                        <template #actions>
                          <icon-edit/>
                          <icon-delete/>
                        </template>
                      </a-list-item>
                    </a-list>
                  </div>
                </a-card>
              </div>
            </template>
          </a-trigger>

        </a-menu>
      </template>
    </a-trigger>

    <!-- drawer -->
    <a-drawer
        :width="500"
        :visible="visible"
        :mask="true"
        :placement="position"
        @ok="handleOk"
        @cancel="handleCancel"
        unmountOnClose
    >
      <template #title>
        常用热门提示词
      </template>
      <div>这是热门prompt词
      </div>
    </a-drawer>
  </div>
</template>

<script setup>
import {IconBug, IconBulb, IconClose, IconMessage} from '@arco-design/web-vue/es/icon';
import {ref} from "vue";
import {useRouter} from "vue-router";

const router = useRouter();
const visible = ref(false);
const position = ref('right');
const tipSeen = ref(false);
const popupVisible = ref(false);
const handleMenuClick = (e) => {
  switch (e) {
    case "1": // contact
      router.push('/settings/index?tab=5')
      break;
    case "2": // prompt
      // handleClick()
      // tipSeen.value = !tipSeen.value
      // router.push('/settings/index?tab=3')
      break;
  }
}
const handleClick = () => {
  visible.value = true;
};
const handleOk = () => {
  visible.value = false;
};
const handleCancel = () => {
  visible.value = false;
}

</script>

<style scoped>
.menu-demo {
}

.button-trigger {
  position: absolute;
  bottom: 80px;
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
}

/* button right */
.button-trigger {
  right: 20px;
  background-color: rgb(var(--arcoblue-6));
}

.button-trigger .button-trigger-active {
  background-color: var(--color-primary-light-4);
}

.prompt-tips-container {
  width: 400px;
  height: 400px;
  position: absolute;
  right: -40px;
  bottom: 0px;
}
</style>