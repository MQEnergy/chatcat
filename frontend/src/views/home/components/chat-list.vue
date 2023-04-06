<template>
  <div class="chat-container" style="padding: 20px;">
    <a-space direction="vertical" :size="30">
      <a-space align="start" :size="10" v-for="n in 10">
        <a-avatar :style="{backgroundColor: n % 2 == 0 ? '' : '#165DFF', overflow: 'hidden'}" :size="32">
            <img v-if="n % 2 == 0"
                alt="avatar"
                src="../../../assets/images/chatgpt.jpg"
            />
          <template v-else>
            A
          </template>
        </a-avatar>
        <a-card hoverable size="medium" :style="{
          width: '500px', borderRadius: '8px',
          backgroundColor: n % 2 == 0 ? '#f8f8f8': '#165DFF'
        }">
          <div :style="{color: n % 2 == 0 ? '#000': '#fff'}">
            vue3 子组件A通过emits给父组件B赋值给一个reactive数组，这个数组在渲染到子组件C，子组件C的props接受如何动态变化，请给出示例。
          </div>
          <template #actions>
            <span :style="{ color: n % 2 == 0 ? '#000' : '#fff'}"><IconShareInternal/></span>
            <a-dropdown @select="handleSelect" position="bl">
              <span :style="{ color: n % 2 == 0 ? '#000' : '#fff'}"><IconMore/></span>
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
const handleSelect = (e) => {

}
</script>

<style scoped>
.chat-container :deep(.arco-card-size-medium) {
  font-size: 16px;
}
</style>