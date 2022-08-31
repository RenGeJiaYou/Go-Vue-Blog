<template>
  <div>
    <a-button
      type="danger"
      @click="logout"
    >退出</a-button>
  </div>
</template>

<script>
export default {
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
.headerBtn {
  display: flex;
  justify-content: end;
  align-items: center;
}
</style>