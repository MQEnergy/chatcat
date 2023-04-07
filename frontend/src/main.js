import {createApp} from 'vue'
import ArcoVue from "@arco-design/web-vue";
import ArcoVueIcon from "@arco-design/web-vue/es/icon";
import router from "./router/index.js";
import '@arco-design/web-vue/dist/arco.css';
import i18n from './locale';
import App from './App.vue'
// code pre
import 'highlight.js/styles/stackoverflow-light.css'
import 'highlight.js/lib/common';
import hljsVuePlugin from "@highlightjs/vue-plugin";

const app = createApp(App)
app.use(ArcoVue, {
  // 用于改变使用组件时的前缀名称
  // componentPrefix: 'a'
})
app.use(ArcoVueIcon)
app.use(router)
app.use(i18n)
app.use(hljsVuePlugin)
app.mount('#app')
