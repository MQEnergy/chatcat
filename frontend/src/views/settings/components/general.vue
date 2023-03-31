<template>
  <div>
    <!-- 配置key -->
    <a-card title="基础配置" :bordered="false" :header-style="{borderColor: 'var(--color-fill-2)'}">
      <a-space direction="vertical" size="medium">
        <a-typography-text>您的 API 密钥存储在您的浏览器本地，绝不会发送到其他任何地方。</a-typography-text>
        <a-input-password :style="{width:'440px'}" placeholder="sk-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" allow-clear>
          <template #prepend>
            API key
          </template>
        </a-input-password>
        <a-alert type="warning">
          <a-link @click="handleNotice">按照指南從 OpenAI 免費獲取 API 密鑰和高達 18 美元的積分</a-link>
        </a-alert>
      </a-space>
    </a-card>
    <!-- 国际化 -->
    <a-card title="国际化" :bordered="false" :header-style="{borderColor: 'var(--color-fill-2)'}">
      <a-select v-model="value" :field-names="fieldNames" :style="{width:'440px'}"
                placeholder="Please select ..." allow-search>
        <a-option style="width: 100%;" v-for="(item, index) in lanList" :value="item.letter" :key="index">
          {{ item.label }}
          <template #extra>
            {{ item.desc }}
          </template>
        </a-option>
      </a-select>
    </a-card>
    <!-- 主题 -->
    <a-card title="主题" :bordered="false" :header-style="{borderColor: 'var(--color-fill-2)'}">
      <a-radio-group @change="handleCheckTheme">
        <template v-for="(item, index) in themeList" :key="index">
          <a-radio :value="item.id">
            <template #radio="{ checked }">
              <a-space
                  align="start"
                  class="custom-radio-card"
                  :class="{ 'custom-radio-card-checked': item.checked }"
              >
                <div className="custom-radio-card-mask">
                  <div className="custom-radio-card-mask-dot"/>
                </div>
                <div>
                  <div className="custom-radio-card-title">
                    {{ item.name }}
                  </div>
                </div>
              </a-space>
            </template>
          </a-radio>
        </template>
      </a-radio-group>
    </a-card>

    <!-- 侧边栏 -->
    <general-drawer :visible="visible" @cancel="handleCancelDrawer"></general-drawer>
  </div>
</template>

<script>
import {reactive, ref} from 'vue';
import GeneralDrawer from "@views/settings/components/general-drawer.vue";

export default {
  components: {GeneralDrawer},
  setup() {
    const checked = ref(true)
    const value = ref('zh');
    const fieldNames = {value: 'city', label: 'text'}
    const themeList = reactive([
      {id: 1, name: '跟随系统', checked: true},
      {id: 2, name: '亮色模式', checked: false},
      {id: 3, name: '黑暗模式', checked: false},
    ])
    const lanList = reactive([
      {
        letter: 'en',
        label: 'English',
        desc: 'English'
      }, {
        letter: 'zh',
        label: '中文(简体)',
        desc: 'Chinese(Simplified)'
      },
    ]);
    const visible = ref(false);
    const handleNotice = () => {
      visible.value = true;
    }
    const handleCheckTheme = () => {
      checked.value = !checked.value;
    }
    const handleCancelDrawer = () => {
      visible.value = false;
    }
    return {
      checked,
      value,
      fieldNames,
      lanList,
      themeList,
      visible,
      handleNotice,
      handleCheckTheme,
      handleCancelDrawer
    }
  }
}
</script>

<style scoped>
.custom-radio-card {
  padding: 10px 16px;
  border: 1px solid var(--color-border-2);
  border-radius: 4px;
  width: 128px;
  box-sizing: border-box;
}

.custom-radio-card-mask {
  height: 14px;
  width: 14px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  border-radius: 100%;
  border: 1px solid var(--color-border-2);
  box-sizing: border-box;
}

.custom-radio-card-mask-dot {
  width: 8px;
  height: 8px;
  border-radius: 100%;
}

.custom-radio-card-title {
  color: var(--color-text-1);
  font-size: 14px;
  font-weight: bold;
}

.custom-radio-card:hover,
.custom-radio-card-checked,
.custom-radio-card:hover .custom-radio-card-mask,
.custom-radio-card-checked .custom-radio-card-mask {
  border-color: rgb(var(--primary-6));
}

.custom-radio-card-checked {
  background-color: var(--color-primary-light-1);
}

.custom-radio-card:hover .custom-radio-card-title,
.custom-radio-card-checked .custom-radio-card-title {
  color: rgb(var(--primary-6));
}

.custom-radio-card-checked .custom-radio-card-mask-dot {
  background-color: rgb(var(--primary-6));
}
</style>