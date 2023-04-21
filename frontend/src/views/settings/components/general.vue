<template>
  <div class="general-container">
    <!-- general -->
    <a-card :title="$t('settings.general')" :bordered="false" :header-style="{borderColor: 'var(--color-fill-2)'}">
      <div class="general-div scrollbar">
        <a-space direction="vertical" size="medium">
          <a-typography-text>{{ $t('settings.general.keyTips') }}</a-typography-text>
          <a-input-password v-model="form.api_key" :style="{width:'460px'}" @blur="handleGeneralSave"
                            placeholder="sk-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" allow-clear>
            <template #prepend>
              {{ $t('settings.general.apiKey') }}
            </template>
          </a-input-password>
          <a-alert type="warning">
            <a-link @click="handleNotice">{{ $t('settings.general.alertTips') }}</a-link>
          </a-alert>
          <a-space>
            {{ $t('settings.general.apiModel') }}:
            <a-select v-model="form.chat_model" :style="{width:'383px'}" @change="handleModelChange"
                      :placeholder="$t('settings.general.apiModel.placeholder')">
              <a-option v-for="(item, index) in modelList" :key="index">{{ item }}</a-option>
            </a-select>
          </a-space>
        </a-space>
        <!-- language -->
        <a-card class="card-container" :title="$t('settings.language')" :bordered="false"
                :header-style="{borderColor: 'var(--color-fill-2)', paddingLeft: '0px'}">
          <a-select v-model="form.language" :field-names="fieldNames" :style="{width:'460px'}"
                    allow-search @change="handleChangeLocale">
            <a-option style="width: 100%;" v-for="(item, index) in locales" :value="item.value" :key="index">
              {{ item.label }}
              <template #extra>
                {{ item.desc }}
              </template>
            </a-option>
          </a-select>
        </a-card>
        <!-- theme -->
        <a-card class="card-container" :title="$t('settings.theme')" :bordered="false"
                :header-style="{borderColor: 'var(--color-fill-2)', paddingLeft: '0px'}">
          <a-radio-group v-model="form.theme" @change="handleCheckTheme">
            <template v-for="(item, index) in themeList" :key="index">
              <a-radio :value="item.id">
                <template #radio="{ checked }">
                  <a-space
                      align="start"
                      class="custom-radio-card"
                      :class="{ 'custom-radio-card-checked': form.theme === item.id }"
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
        <!-- advanced -->
        <a-card class="card-container" :title="$t('settings.advanced')" :bordered="false"
                :header-style="{borderColor: 'var(--color-fill-2)', paddingLeft: '0px'}">
          <template #extra>
            <a-button type="outline" @click="handleGeneralSave">
              <template #icon>
                <icon-save />
              </template>
              {{ $t('common.save') }}
            </a-button>
          </template>
          <a-list>
            <a-list-item v-for="(item, index) in advancedList" :key="index">
              <a-list-item-meta
                  :title="item.title"
                  :description="item.desc"
              >
              </a-list-item-meta>
              <template #actions>
                <div :style="{ width: '160px', marginLeft: '20px' }">
                  <a-slider :min="item.min" :max="item.max" v-if="item.type === 1" :marks="item.marks" :step="item.step"
                            :default-value="item.value" @change="handleSliderChange($event, item)"/>
                  <a-input v-else v-model="item.value"></a-input>
                </div>
              </template>
            </a-list-item>
          </a-list>
        </a-card>
      </div>
    </a-card>
    <!-- drawer -->
    <general-drawer :visible="visible" @cancel="handleCancelDrawer"></general-drawer>
  </div>
</template>

<script setup>
import {computed, onMounted, reactive, ref} from 'vue';
import GeneralDrawer from "@views/settings/components/general-drawer.vue";
import {LOCALE_OPTIONS} from "@/locale";
import useLocale from "@/hooks/locale";
import {useI18n} from "vue-i18n";
import {GetGeneralInfo, SetGeneralData} from "../../../../wailsjs/go/setting/Service.js";
import {Message} from "@arco-design/web-vue";

const {t, locale} = useI18n();
const fieldNames = {value: 'city', label: 'text'}
const themeList = computed(() => [
  {id: 2, name: t('settings.theme.toLight')},
  {id: 3, name: t('settings.theme.toDark')},
  {id: 1, name: t('settings.theme.system')},
])
const advancedList = reactive([
  {
    title: '随机性 (temperature)',
    tag: 'temperature',
    desc: '值越大恢复越随机，介于0-2之间，大于1的值可能会导致乱码',
    value: 0.7,
    min: 0,
    max: 2,
    step: 0.1,
    type: 1, // 1：滑块 2：输入框
    marks: {0: '0', 2: '2'},
  }, {
    title: '单次回复限制 (max_tokens)',
    tag: 'max_tokens',
    desc: '单次交互所用的最大Token数，设置为0表示按照当前模型允许的最大token数量自动计算',
    value: 0,
    min: 20,
    max: 9999,
    step: 1,
    type: 2,
    marks: {20: '20'}
  }, {
    title: '对话新鲜度 (presence_penalty)',
    tag: 'presence_penalty',
    desc: '正值会根据到目前为止是否出现在文本中来惩罚新标记，从而增加模型谈论新主题的可能性。介于-2.0 和 2.0之间',
    value: 0,
    min: -2,
    max: 2,
    step: 0.1,
    type: 1,
    marks: {'-2': '-2', 2: '2'},
  }, {
    title: '对话重复性 (frequency_penalty)',
    tag: 'frequency_penalty',
    desc: '正值会根据新标记在文本中的现有频率对其进行惩罚，从而降低模型逐字重复同一行的可能性。介于-2.0 和 2.0之间',
    value: 0,
    min: -2,
    max: 2,
    step: 0.1,
    type: 1,
    marks: {'-2': '-2', 2: '2'},
  }, {
    title: '返回数量 (N)',
    tag: 'n',
    desc: 'API会生成多少个可能的文本选项供用户选择，它会很快消耗你的令牌配额。请谨慎使用 默认1',
    value: 1,
    min: 1,
    max: 10,
    step: 1,
    type: 1,
    marks: {1: '1', 10: '10'},
  },
])
const modelList = reactive(['gpt-3.5-turbo', 'gpt-3.5-turbo-0301', 'gpt-4', 'gpt-4-0314', 'gpt-4-32k', 'gpt-4-32k-0314'])
const form = ref({
  api_key: '',
  chat_model: '',
  ask_model: '',
  language: locale.value,
  theme: 1,
  proxy_url: '',
  account: '',
  access_token: '',
  is_sync: 1,
  sync_time: 0,
  temperature: "0.7",
  max_tokens: 0,
  presence_penalty: "0",
  frequency_penalty: "0",
  n: 1,
})
const locales = [...LOCALE_OPTIONS]
const visible = ref(false);
const {changeLocale, currentLocale} = useLocale();

const handleNotice = () => {
  visible.value = true;
}
const handleCheckTheme = (e) => {
  form.value.theme = e;
  switch (e) {
    case 1:
      const darkThemeMq = window.matchMedia("(prefers-color-scheme: dark)");
      darkThemeMq.addListener(e => {
        console.log(e)
        if (e.matches) {
          document.body.setAttribute('arco-theme', 'dark');
        } else {
          document.body.removeAttribute('arco-theme');
        }
      });
    case 2:
      document.body.removeAttribute('arco-theme');
      break;
    case 3:
      document.body.setAttribute('arco-theme', 'dark');
      break;
  }
  handleGeneralSave();
}
const handleCancelDrawer = () => {
  visible.value = false;
}
const handleModelChange = (e) => {
  form.value.ask_model = e;
  handleGeneralSave()
}
const handleChangeLocale = (e) => {
  handleGeneralSave();
  changeLocale(e);
}
const handleGeneralSave = (e) => {
  SetGeneralData(form.value).then(res => {
    console.log(res)
  })
}
const handleSliderChange = (e, row) => {
  switch (row.tag) {
    case 'temperature':
      form.value.temperature = e.toString();
      break;
    case 'max_tokens':
      form.value.max_tokens = e;
      break;
    case 'presence_penalty':
      form.value.presence_penalty = e.toString();
      break;
    case 'frequency_penalty':
      form.value.frequency_penalty = e.toString();
      break;
    case 'n':
      form.value.n = e;
      break;
  }
  // handleGeneralSave(e)
}
const initGeneralInfo = (e) => {
  GetGeneralInfo().then(res => {
    if (res.code === 0) {
      form.value = res.data;
    }
  })
}
onMounted(() => {
  if (window.go === undefined) {
    Message.error(t('common.panic'));
  } else {
    initGeneralInfo();
  }
})
</script>

<style scoped>
.custom-radio-card {
  padding: 10px 16px;
  border: 1px solid var(--color-border-2);
  border-radius: 4px;
  width: 134px;
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

.general-div {
  height: 80vh;
}

.general-container :deep(.arco-card-body) {
  overflow-x: hidden;
  overflow-y: scroll;
}

.general-container :deep(.arco-card-body::-webkit-scrollbar) {
  display: none;
}

.general-container::-webkit-scrollbar {
  display: none;
}

.card-container {
  margin-top: 20px;
  /*border-color: var(--color-fill-2);*/
  /*padding-left: 0px;*/
}

.card-container :deep(.arco-card-body) {
  padding-left: 0px;
}

.card-container :deep(.arco-slider-with-marks) {
  margin-bottom: initial !important;
  padding: initial !important;
}
</style>