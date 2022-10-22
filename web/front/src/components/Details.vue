<template>
  <div>
    <div class="d-flex justify-center pa-1 text-h4 font-weight-bold">{{artInfo.title}}</div>

    <v-divider class="ma-10"></v-divider>

    <v-alert
      class="ma-3"
      color="grey lighten-5"
      border="left"
      elevation="1"
    >{{artInfo.desc}}</v-alert>

    <!-- 绑定的数据(string 类型)会被看做 html文档而被解析 -->
    <div
      class="ma-5 pa-3 text-justify"
      v-html="artInfo.content"
    ></div>
  </div>
</template>

<script>
// 文章内容详情
export default {
  props: ["id"],
  data() {
    return {
      //文章模型
      artInfo: {},
    };
  },
  created() {
    this.getArtById();
  },
  methods: {
    async getArtById() {
      const { data: res } = await this.$axios.get(`article/info/${this.id}`);

      this.artInfo = res.data;
      console.log(this.aritcleInfo);
    },
  },
};
</script>

<style>
</style>

