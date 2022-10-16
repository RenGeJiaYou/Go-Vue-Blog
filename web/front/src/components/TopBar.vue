<template>
  <div>
    <v-app-bar
      app
      color="primary darken-4"
    >
      <v-avatar
        color="grey lighten-2"
        size="43"
      ></v-avatar>
      <!--见文档 样式与动画-间距-->
      <v-container class="py-2 fill-height">
        <v-btn
          text
          color="white"
          @click="$router.push('/')"
        >首页</v-btn>
        <v-btn
          v-for="item in cateList"
          :key="item.id"
          text
          color="white"
        >{{item.name}}</v-btn>
      </v-container>

      <!-- 该标签会将前后的元素推向其父容器的两侧 -->
      <v-spacer></v-spacer>

      <v-responsive
        max-width="260"
        color="white"
      >
        <v-text-field
          dense
          flat
          hide-details
          solo-inverted
          dark
          rounde
          placeholder="搜索文章"
          clearable
        ></v-text-field>
      </v-responsive>
    </v-app-bar>
  </div>
</template>

<script>
export default {
  name: "TopBar",
  data() {
    return {
      //分类列表
      cateList: [],
    };
  },
  created() {
    this.getCateList();
  },
  methods: {
    //获取分类
    async getCateList() {
      const { data: res } = await this.$axios.get("categories");
      this.cateList = res.data.slice(0, 7); //标签先只要一部分，顶端栏放不下了。
    },
  },
};
</script>

<style>
</style>

