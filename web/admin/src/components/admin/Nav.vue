<template>
  <a-layout-sider
    breakpoint="md"
    collapsible
    v-model="collapsed"
  >
    <div class="log">
      <a-icon
        v-if="collapsed"
        type="book"
      />
      <span v-if="!collapsed">Go-Vue Blog 系统</span>
    </div>
    <a-menu
      mode="inline"
      theme="dark"
      @click="handleClick"
    >
      <a-menu-item key="index">
        <a-icon type="dashboard" />
        <span>仪表盘</span>
      </a-menu-item>

      <a-sub-menu key="article">
        <span slot="title">
          <a-icon type="read" />
          <span>文章管理</span>
        </span>
        <a-menu-item key="addart">
          <a-icon type="form" />
          <span>写文章</span>
        </a-menu-item>
        <a-menu-item key="artlist">
          <a-icon type="profile" />
          <span>文章列表</span>
        </a-menu-item>
      </a-sub-menu>

      <a-menu-item key="catelist">
        <a-icon type="tags" />
        <span>分类列表</span>
      </a-menu-item>

      <a-menu-item key="userlist">
        <a-icon type="user" />
        <span>用户列表</span>
      </a-menu-item>
    </a-menu>
  </a-layout-sider>
</template>

<script>
export default {
  name: "AdminNav",

  data() {
    return {
      //标识侧边栏当前是否收起为一窄竖条
      collapsed: false,
    };
  },

  mounted() {},

  methods: {
    handleClick(item, key, keyPath) {
      //item      是当前所点击对象的引用
      //key       item 所绑定的 key 值.注意：在<template>中item 的 key 一定和 router 的 path 保持一致
      //keyPath   一个数组，记录父子项的所有key。index越大辈分越大。如点击《文章列表》0: "getArts" 1: "article"
      // console.log(item, key, keyPath);

      this.$router.push("/" + item.key).catch((err) => {
        console.log(err);
      });
    },
  },
};
</script>

<style  scoped>
.log {
  height: auto;
  margin: 16px;
  background-color: #fff;

  text-align: center;

  border-radius: 5px;
  margin-bottom: 20px;
}
</style>