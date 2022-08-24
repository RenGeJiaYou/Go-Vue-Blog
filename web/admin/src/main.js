import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import axios from "axios";
import "./plugin/antui"; //此处导入按需加载的 antui.js

axios.defaults.baseURL = "http://localhost:3000/api/v1";
//挂载到原型上，让每一个 Vue 实例对象都拥有这个属性($只是和用户自定义属性/方法名区分开，无特殊含义)
Vue.prototype.$axios = axios;

Vue.config.productionTip = false;

new Vue({
  router,
  render: h => h(App),
}).$mount("#app");
