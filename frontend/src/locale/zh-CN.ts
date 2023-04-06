import localeSettings from './zh-CN/settings';
import localeCommon from './zh-CN/common';
import localeHome from './zh-CN/home';

export default {
  'menu.faq': '常见问题',
  'navbar.docs': '文档中心',
  'navbar.action.locale': '切换为中文',
  ...localeCommon,
  ...localeHome,
  ...localeSettings,
};
