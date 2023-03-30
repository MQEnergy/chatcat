<template>
  <div style="--wails-draggable: drag;">
    <a-page-header
        :style="{
          background: 'var(--color-bg-2)',
          '--wails-draggable': 'drag', 'top': 0, 'position': 'absolute',
          'width': '100%', 'height': '30px', 'border-bottom': '1px solid var(--color-neutral-3)',
          'z-index': '9'
      }"
        title="这是prompt分类"
        subtitle="这是初始化的prompt词"
        :show-back="false"
    >
      <template #extra>
        <a-radio-group type="button" default-value="large">
          <a-radio value="mini">GPT3.5</a-radio>
          <a-radio value="small">6 messages</a-radio>
          <a-radio value="large">193 tokens</a-radio>
        </a-radio-group>
      </template>
    </a-page-header>
  </div>
</template>

<script>
export default {
  name: "content-header"
}
</script>

<style scoped>

</style>