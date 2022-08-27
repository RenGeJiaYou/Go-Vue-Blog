//在本文件内按需导入 antdv 组件
import Vue from "vue";
import { Button, FormModel, Input, Icon } from "ant-design-vue";

//插件通常用来为 Vue 添加全局功能，可以通过全局方法 Vue.use() 使用插件，
//需在调用 new Vue() 启动应用之前声明。
Vue.use(Button);
Vue.use(FormModel);
Vue.use(Input);
Vue.use(Icon);
