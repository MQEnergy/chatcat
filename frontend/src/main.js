import {createApp} from 'vue'
import App from './App.vue'
import ArcoVue from "@arco-design/web-vue";
import ArcoVueIcon from "@arco-design/web-vue/es/icon";
import router from "./router/index.js";
import '@arco-design/web-vue/dist/arco.css';

const app = createApp(App)
app.use(ArcoVue,{
  // 用于改变使用组件时的前缀名称
  componentPrefix: 'a'
})
app.use(ArcoVueIcon)
app.use(router)
app.mount('#app')
