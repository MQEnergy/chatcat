<template>
  <div class="menu-demo">
    <a-trigger
        :trigger="['click', 'hover']"
        clickToClose
        position="top"
        v-model:popupVisible="popupVisible2"
    >
      <div :class="`button-trigger ${popupVisible2 ? 'button-trigger-active' : ''}`">
        <IconClose v-if="popupVisible2" />
        <IconMessage v-else />
      </div>
      <template #content>
        <a-menu
            :style="{ marginBottom: '-4px' }"
            mode="popButton"
            :tooltipProps="{ position: 'left' }"
            showCollapseButton
        >
          <a-menu-item key="1">
            <template #icon><IconBug></IconBug></template>
            Bugs
          </a-menu-item>
          <a-menu-item key="2">
            <template #icon><IconBulb></IconBulb></template>
            Ideas
          </a-menu-item>
        </a-menu>
      </template>
    </a-trigger>
  </div>
</template>

<script>
import {
  IconBug,
  IconBulb,
  IconClose,
  IconMessage,
} from '@arco-design/web-vue/es/icon';

export default {
  components: {
    IconBug,
    IconBulb,
    IconClose,
    IconMessage,
  },
  data() {
    return {
      popupVisible1: false,
      popupVisible2: false,
    };
  }
};
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