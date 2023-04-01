import { createI18n } from 'vue-i18n';
import en from './en-US';
import cn from './zh-CN';

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
