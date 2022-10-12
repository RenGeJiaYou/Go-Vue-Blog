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
        <v-btn elevation="12"></v-btn>
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
          placeholder
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
      this.cateList = res;
    },
  },
};
</script>

<style>
</style>

