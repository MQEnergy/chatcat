<template>
  <div>
    <a-card :title="$t('settings.general')" :bordered="false" :header-style="{borderColor: 'var(--color-fill-2)'}">
      <a-space direction="vertical" size="medium">
        <a-typography-text>{{ $t('settings.general.keyTips') }}</a-typography-text>
        <a-input-password :style="{width:'460px'}" placeholder="sk-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" allow-clear>
          <template #prepend>
            API key
          </template>
        </a-input-password>
        <a-alert type="warning">
          <a-link @click="handleNotice">{{ $t('settings.general.alertTips') }}</a-link>
        </a-alert>
      </a-space>
    </a-card>
    <a-card :title="$t('settings.language')" :bordered="false" :header-style="{borderColor: 'var(--color-fill-2)'}">
      <a-select v-model="currentLocale" :field-names="fieldNames" :style="{width:'460px'}"
                allow-search @change="changeLocale">
        <a-option style="width: 100%;" v-for="(item, index) in locales" :value="item.value" :key="index">
          {{ item.label }}
          <template #extra>
            {{ item.desc }}
          </template>
        </a-option>
      </a-select>
    </a-card>
    <a-card :title="$t('settings.theme')" :bordered="false" :header-style="{borderColor: 'var(--color-fill-2)'}">
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

    <general-drawer :visible="visible" @cancel="handleCancelDrawer"></general-drawer>
  </div>
</template>

<script setup>
import {computed, ref} from 'vue';
import GeneralDrawer from "@views/settings/components/general-drawer.vue";
import {LOCALE_OPTIONS} from "@/locale";
import useLocale from "@/hooks/locale";
import {useI18n} from "vue-i18n";

const {t} = useI18n();
const fieldNames = {value: 'city', label: 'text'}
const themeList = computed(() => [
  {id: 1, name: t('settings.theme.system'), checked: true},
  {id: 2, name: t('settings.theme.toLight'), checked: false},
  {id: 3, name: t('settings.theme.toDark'), checked: false},
])
const locales = [...LOCALE_OPTIONS]
const visible = ref(false);
const {changeLocale, currentLocale} = useLocale();

const handleNotice = () => {
  visible.value = true;
}
const handleCheckTheme = (e) => {
  console.log(e);
  const theme = themeList.value.filter((item) => {
    return item.id == e
  });
  theme[0].checked = true;
}
const handleCancelDrawer = () => {
  visible.value = false;
}
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