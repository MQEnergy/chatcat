<template>
  <div class="menu-container">
    <a-trigger
        :trigger="['click', 'hover']"
        clickToClose
        position="top"
        v-model:popupVisible="popupVisible2"
    >
      <div :class="`button-trigger ${popupVisible2 ? 'button-trigger-active' : ''}`">
        <IconClose v-if="popupVisible2"/>
        <IconMessage v-else/>
      </div>
      <template #content>
        <a-menu
            :style="{ marginBottom: '-4px' }"
            mode="popButton"
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
          <a-menu-item key="2" style="border: 1px solid #f5f5f5;">
            <template #icon>
              <IconBulb></IconBulb>
            </template>
            {{ $t('settings.prompt') }}
          </a-menu-item>
        </a-menu>
      </template>
    </a-trigger>
    <div v-if="tipSeen"
        style="width: 400px; height: 400px; background: #fff; border-radius: 8px; position: absolute;
             right: 80px; bottom: 80px; border: 1px solid #eee; box-shadow: 0px 1px 5px rgba(0,0,0,0.2)">
      这是prompt热门提示词
    </div>
    <!-- drawer -->
    <a-drawer
        style="width: 300px;"
        :height="340"
        :visible="visible"
        :mask="false"
        :placement="position"
        @ok="handleOk"
        @cancel="handleCancel"
        unmountOnClose
    >
      <template #title>
        Title
      </template>
      <div>这是热门prompt词
      </div>
    </a-drawer>
  </div>
</template>

<script setup>
import {IconBug, IconBulb, IconClose, IconMessage,} from '@arco-design/web-vue/es/icon';
import {defineComponent, ref} from "vue";
import {useRouter} from "vue-router";

const router = useRouter();
const visible = ref(false);
const position = ref('bottom');
const tipSeen = ref(false);
defineComponent({
  IconBug, IconBulb, IconClose, IconMessage
})
const popupVisible2 = ref(false);
const handleMenuClick = (e) => {
  console.log("e", e)
  switch (e) {
    case "1": // contact
      router.push('/settings/index?tab=4')
      break;
    case "2": // prompt
      // handleClick()
      tipSeen.value = !tipSeen.value
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
</style>