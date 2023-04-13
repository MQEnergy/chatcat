<template>
  <div class="chat-container" style="padding: 20px;">
    <a-space direction="vertical" :size="30">
      <a-space align="start" :size="10" v-for="(item, index) in chatList" :key="index">
        <a-avatar
            :style="{backgroundColor: item.from === 2 ? '#fff' : '#165DFF', overflow: 'hidden'}"
            :size="32">
          <img v-if="item.from === 2"
               alt="avatar"
               :src="ChatGPTLogo"
          />
          <template v-else>
            A
          </template>
        </a-avatar>
        <a-card hoverable :style="{
          width: '500px', borderTopRightRadius: '12px', borderColor: item.from === 2 ? '' : '#c7d7fa',
          borderBottomRightRadius: '12px', borderBottomLeftRadius: '12px',
          backgroundColor: item.from === 2 ? '#f8f8f8': '#e8f3ff'
        }">
          <div class="chat-div" :style="{color: item.from === 2 ? '#000': '#155dff'}">
            {{ item.content }}
          </div>
          <template #actions>
            <span :style="{ color: item.from === 2 ? '#000' : '#000'}"><IconShareInternal/></span>
            <a-dropdown @select="handleSelect" position="bl">
              <span :style="{ color: item.from === 2 ? '#000' : '#000'}"><IconMore/></span>
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
import {reactive, watch} from "vue";
// import 'highlight.js/lib/common';
// import hljsVuePlugin from "@highlightjs/vue-plugin";
const props = defineProps({
  list: {
    type: Array,
    default: []
  },

});
let chatList = reactive(props.list)
watch(() => props.list, () => {
  chatList = props.list
})
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