<template>
  <div class="chat-container" style="padding: 20px;">
    <a-space direction="vertical" :size="30">
      <a-space align="start" :size="10" v-for="n in 10">
        <a-avatar :style="{backgroundColor: '#165DFF'}" :size="32">A</a-avatar>
        <a-card hoverable size="medium" :style="{
          width: '500px', 'border-radius': '8px',
          background: n % 2 == 0 ? 'white': 'blue',
        }">
          <div :style="{color: n % 2 == 0 ? 'black': 'white'}">
            ByteDance's core product, Toutiao ("Headlines"), is a content platform in
            China and around the world. Toutiao started out as a news recommendation
            engine and gradually evolved into a platform delivering content in various
            formats.
          </div>
          <template #actions>
            <span :style="{ color: n % 2 == 0 ? 'black' : 'white'}"><IconShareInternal/></span>
            <a-dropdown @select="handleSelect" position="bl">
              <span :style="{ color: n % 2 == 0 ? 'black' : 'white'}"><IconMore/></span>
              <template #content>
                <a-doption>Option 1</a-doption>
                <a-doption>Option 2</a-doption>
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