import { createI18n } from 'vue-i18n';
import en from './en-US';
import cn from './zh-CN';

import zhCN from '@arco-design/web-vue/es/locale/lang/zh-cn';
import enUS from '@arco-design/web-vue/es/locale/lang/en-us';
import esES from '@arco-design/web-vue/es/locale/lang/es-es';
import jaJP from '@arco-design/web-vue/es/locale/lang/ja-jp';
import idID from '@arco-design/web-vue/es/locale/lang/id-id';
import frFR from '@arco-design/web-vue/es/locale/lang/fr-fr';
import ptPT from '@arco-design/web-vue/es/locale/lang/pt-pt';
import deDE from '@arco-design/web-vue/es/locale/lang/de-de';
import koKR from '@arco-design/web-vue/es/locale/lang/ko-kr';
import itIT from '@arco-design/web-vue/es/locale/lang/it-it';
import thTH from '@arco-design/web-vue/es/locale/lang/th-th';
import viVN from '@arco-design/web-vue/es/locale/lang/vi-vn';

const locales = {
  'zh-CN': zhCN,
  'en-US': enUS,
  'es-ES': esES,
  'ja-JP': jaJP,
  'id-ID': idID,
  'fr-FR': frFR,
  'pt-PT': ptPT,
  'de-DE': deDE,
  'ko-KR': koKR,
  'it-IT': itIT,
  'th-TH': thTH,
  'vi-VN': viVN,
};

export const LOCALE_OPTIONS = [
  { label: '中文(简体)', value: 'zh-CN', desc: 'Chinese(Simplified)' },
  { label: 'English', value: 'en-US', desc: 'English' },
];
// Todo 这一部分可能需要在sqlite里面获取
const defaultLocale = localStorage.getItem('cc-locale') || 'zh-CN';

const i18n = createI18n({
  locale: defaultLocale,
  fallbackLocale: 'en-US',
  legacy: false,
  allowComposition: true,
  messages: {
    'en-US': en,
    'zh-CN': cn,
  },
});

export default i18n;
