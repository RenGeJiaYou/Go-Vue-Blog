import Vue from "vue";
import VueRouter from "vue-router";
import Login from "../views/Login.vue";
import Admin from "../views/Admin.vue";

//页面路由组件
import Index from "../components/admin/index";
import AddArt from "../components/article/AddArt";
import Artlist from "../components/article/ArtList";
import CateList from "../components/category/CateList";
import UserList from "../components/user/UserList";

Vue.use(VueRouter);

const routes = [
  {
    path: "/login",
    name: "login",
    component: Login,
  },
  {
    path: "/",
    name: "admin",
    component: Admin,
    //子路由的path 如果前面写'/'将会定位到根路径
    children: [
      {
        path: "index",
        component: Index,
      },
      {
        path: "addart",
        component: AddArt,
      },
      {
        path: "addart/:id", //添加、编辑文章
        component: AddArt,
        props: true,
      },
      {
        path: "artlist", //文章列表
        component: Artlist,
      },
      {
        path: "catelist",
        component: CateList,
      },
      {
        path: "userlist",
        component: UserList,
      },
    ],
  },
];

const router = new VueRouter({
  routes,
});

//路由守卫
router.beforeEach((to, form, next) => {
  const isToken = localStorage.getItem("token") ? true : false;

  //放行无需审查的页面
  if (to.path === "/login") {
    next();
  }

  //放行需审查,且持有 token 的页面;禁止需审查,且无 token 的页面
  isToken ? next() : next("/login");
});

export default router;
