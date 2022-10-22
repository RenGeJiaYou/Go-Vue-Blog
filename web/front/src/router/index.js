import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home.vue";
import Article from "../components/Article.vue";
import Details from "../components/Details.vue";

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
      {
        path: "/details/:id", //:id 作为key,实际传入的文章 ID 作为value 组成一个k-v.可在子组件中通过props:['id']获得，并像一个data那样用
        component: Details,
        meta: { title: "文章详情" },
        props: true,
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
