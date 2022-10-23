import Vue from "vue";
import axios from "axios";

let Url = "http://localhost:3000/api/v1/";

axios.defaults.baseURL = Url;
axios.interceptors.request.use(config => {
  //localstorage储存的TOKEN不会自己到请求报文首部,需要手动在请求前添加
  config.headers.Authorization = "Bearer " + localStorage.getItem("token");
  return config;
});
//挂载到原型上，让每一个 Vue 实例对象都拥有这个属性($只是和用户自定义属性/方法名区分开，无特殊含义)
Vue.prototype.$axios = axios;

export { Url };
