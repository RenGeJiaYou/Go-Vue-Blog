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
    path: "/",
    redirect: "/login", //重定向到"/index"这条路由项
  },
  {
    path: "/login",
    name: "login",
    component: Login,
  },
  {
    path: "/admin",
    name: "admin",
    component: Admin,
    children: [
      {
        path: "/index",
        component: Index,
      },
      {
        path: "/addart",
        component: AddArt,
      },
      {
        path: "/artlist",
        component: Artlist,
      },
      {
        path: "/catelist",
        component: CateList,
      },
      {
        path: "/userlist",
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
