<template>
  <div class="chat-container" style="padding: 20px;">
    <a-space direction="vertical" :size="30">
      <a-space align="start" :size="10" v-for="n in 10">
        <a-avatar :style="{backgroundColor: n % 2 == 0 ? '#fff' : '#165DFF', overflow: 'hidden'}" :size="32">
            <img v-if="n % 2 == 0"
                alt="avatar"
                :src="ChatGPTLogo"
            />
          <template v-else>
            A
          </template>
        </a-avatar>
        <a-card hoverable :style="{
          width: '500px', borderTopRightRadius: '12px', borderColor: n % 2 == 0 ? '' : '#c7d7fa',
          borderBottomRightRadius: '12px', borderBottomLeftRadius: '12px',
          backgroundColor: n % 2 == 0 ? '#f8f8f8': '#e8f3ff'
        }">
          <div class="chat-div" :style="{color: n % 2 == 0 ? '#000': '#155dff'}">
            <template v-if="n % 2 == 0">
              <code>
                {{ codePre }}
              </code>
            </template>
            vue3 子组件A通过emits给父组件B赋值给一个reactive数组，这个数组在渲染到子组件C，子组件C的props接受如何动态变化，请给出示例。
          </div>
          <template #actions>
            <span :style="{ color: n % 2 == 0 ? '#000' : '#000'}"><IconShareInternal/></span>
            <a-dropdown @select="handleSelect" position="bl">
              <span :style="{ color: n % 2 == 0 ? '#000' : '#000'}"><IconMore/></span>
              <template #content>
                <a-doption>保存</a-doption>
                <a-doption>复制</a-doption>
              </template>
            </a-dropdown>
          </template>
        </a-card>
      </a-space>
    </a-space>
  </div>
</template>

<script setup>
import ChatGPTLogo from '@assets/images/chatgpt_black_logo.svg';
import {ref} from "vue";
const codePre = ref(`
<!-- ChildComponent.vue -->
<template>
  <div>
    <p>子组件接收到的名字：{{ name }}</p>
    <button @click="changeName">点击修改名字</button>
  </div>
</template>

<script>
import { defineComponent } from 'vue';

export default defineComponent({
  props: {
    name: {
      type: String,
      required: true
    }
  },
  methods: {
    changeName() {
      this.$emit('update:name', '新名字');
    }
  }
});
<\/script>`)
const handleSelect = (e) => {

}
</script>

<style scoped>
.chat-container :deep(.arco-card-size-medium) {
  /*font-size: 16px;*/
}
.chat-container :deep(.arco-card-size-medium .arco-card-body) {
  padding: 15px;
}
</style>