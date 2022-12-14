//在本文件内按需导入 antdv 组件
import Vue from "vue";
import {
  Button,
  FormModel,
  Input,
  Icon,
  Message,
  Layout,
  Menu,
  Card,
  Table,
  Row,
  Col,
  LocaleProvider,
  Modal,
  Select,
  Upload,
} from "ant-design-vue";

//不全局使用，只是挂载到对象上
Message.config({
  top: `40px`,
  duration: 2,
  maxCount: 3,
});
Vue.prototype.$message = Message;
Vue.prototype.$confirm = Modal.confirm; //加()就执行了

//插件通常用来为 Vue 添加全局功能，可以通过全局方法 Vue.use() 使用插件，
//需在调用 new Vue() 启动应用之前声明。
Vue.use(Button);
Vue.use(FormModel);
Vue.use(Input);
Vue.use(Icon);
Vue.use(Layout);
Vue.use(Menu);

Vue.use(Card);
Vue.use(Table);
Vue.use(Row);
Vue.use(Col);

Vue.use(LocaleProvider);
Vue.use(Modal);
Vue.use(Select);
Vue.use(Upload);
