import localeSettings from './en-US/settings';
import localeCommon from './en-US/common';
import localeHome from './en-US/home';

export default {
  'menu.faq': 'FAQ',
  'navbar.docs': 'Docs',
  'navbar.action.locale': 'Switch to English',
  ...localeCommon,
  ...localeHome,
  ...localeSettings,
};
