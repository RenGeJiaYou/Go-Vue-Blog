import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home.vue";
import Article from "../components/Article.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "home",
    component: Home,
    children: [
      {
        path: "/",
        component: Article,
        meta: { title: "Welcome to Go Vue Blog !" },
      },
    ],
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

router.beforeEach((to, form, next) => {
  if (to.meta.title) {
    document.title = to.meta.title;
    next(); //这里的 next() 是 beforeEach（）内定义的闭包函数吗？
  }
});

export default router;
