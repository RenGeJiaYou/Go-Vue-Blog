<template>
  <a-card>
    <a-form-model></a-form-model>
  </a-card>
</template>

<script>
export default {
  name: "",
  data() {
    return {
      //数据模型
      profileInfo: {
        id: 1,
        name: "",
        img: "",
        avatar: "",
        state: "",
        desc: "",
        qqchat: "",
        wechat: "",
        weibo: "",
        email: "",
      },
    };
  },
  created() {
    this.getProfileInfo();
    this.headers = {
      Authorization: `Bearer ${window.localStorage.getItem("token")}`,
    };
  },
  methods: {
    // 获取个人简介
    async getProfileInfo() {
      const { data: res } = await this.$axios.get(
        `profile/${this.profileInfo.id}`
      );
      if (res.status !== 200) {
        if (res.status === 1004 || 1005 || 1006 || 1007) {
          window.sessionStorage.clear();
          this.$router.push("/login");
        }
        this.$message.error(res.message);
      }
      // console.log(res);
      this.profileInfo = res.data;
    },

    // 上传头像

    //更新
  },
};
</script>

<style>
</style>

