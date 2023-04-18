<template>
  <div>
    <a-card :title="$t('settings.general')" :bordered="false" :header-style="{borderColor: 'var(--color-fill-2)'}">
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
    </a-card>
    <a-card :title="$t('settings.language')" :bordered="false" :header-style="{borderColor: 'var(--color-fill-2)'}">
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
    <a-card :title="$t('settings.theme')" :bordered="false" :header-style="{borderColor: 'var(--color-fill-2)'}">
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

const {t, locale} = useI18n();
const fieldNames = {value: 'city', label: 'text'}
const themeList = computed(() => [
  {id: 1, name: t('settings.theme.system')},
  {id: 2, name: t('settings.theme.toLight')},
  {id: 3, name: t('settings.theme.toDark')},
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
})
const locales = [...LOCALE_OPTIONS]
const visible = ref(false);
const {changeLocale, currentLocale} = useLocale();

const handleNotice = () => {
  visible.value = true;
}
const handleCheckTheme = (e) => {
  form.value.theme = e;
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
const initGeneralInfo = (e) => {
  GetGeneralInfo().then(res => {
    if (res.code === 0) {
      form.value = res.data;
    }
  })
}
onMounted(() => {
  initGeneralInfo();
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
</style>