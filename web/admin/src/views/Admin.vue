<template>
  <a-layout class="container">
    <a-layout-sider>
      <div class="log">Go-Vue Blog 系统</div>
      <Nav></Nav>
    </a-layout-sider>
    <a-layout>
      <a-layout-header class="headerBtn">
        <a-button
          type="danger"
          @click="logout"
        >退出</a-button>
      </a-layout-header>
      <a-layout-content>Content</a-layout-content>
      <a-layout-footer class="footerContainer">
        <Footer></Footer>
      </a-layout-footer>
    </a-layout>
  </a-layout>
</template>

<script>
import Nav from "../components/admin/Nav";
import Footer from "../components/admin/Footer";
export default {
  components: {
    Nav,
    Footer,
  },
  data() {
    return {};
  },
  methods: {
    logout() {
      //1.清空token
      window.localStorage.removeItem("token");
      //2.回到路由 '/login'
      this.$router.push("/login");
    },
  },
};
</script>

<style scoped>
.container {
  height: 100%;
}
.footerContainer {
  background-color: #aaa;
  padding: 0;
}
.headerBtn {
  display: flex;
  justify-content: end;
  align-items: center;
}
.log {
  height: auto;
  margin: 16px;
  background-color: #fff;

  text-align: center;

  border-radius: 5px;
}
</style>