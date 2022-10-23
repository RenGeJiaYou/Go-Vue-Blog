import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import "./plugin/axios";
import "./plugin/antui"; //此处导入按需加载的 antui.js
import "./assets/css/style.css"; //导入初始化的 css 文件,必须有后缀名，否则报错

Vue.config.productionTip = false;

new Vue({
  router,
  render: h => h(App),
}).$mount("#app");
