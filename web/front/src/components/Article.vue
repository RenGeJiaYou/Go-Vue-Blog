<template>
  <v-col class="grey lighten-4">
    <v-card>
      <v-row
        class="ma-1"
        v-for=" i in articleList"
        :key="i.ID"
        no-gutters
      >
        <v-col cols="3">
          <img
            width="100%"
            :src="i.image"
          />
        </v-col>
        <v-col>
          <v-card-title>
            <v-chip
              color="pink lighten-3"
              label
              class="white--text mr-5"
            >{{i.Category.name}}</v-chip>
            {{i.title}}
          </v-card-title>
          <v-card-subtitle>{{i.desc}}</v-card-subtitle>
          <v-card-text>{{i.content}}</v-card-text>
        </v-col>
      </v-row>
    </v-card>
  </v-col>
</template>

<script>
export default {
  name: "ArticleBar",
  data() {
    return {
      //文章列表的集合
      articleList: [],

      //请求参数
      queryParams: {
        pageSize: 5,
        PageNum: 1,
      },
      //保存文章总数
      total: 0,
    };
  },
  created() {
    this.getArtList();
  },
  methods: {
    //获取文章列表
    async getArtList() {
      const { data: res } = await this.$axios.get("articles", {
        params: {
          pagesize: this.queryParams.pageSize,
          pagenum: this.queryParams.PageNum,
        },
      });
      this.articleList = res.data;
      this.total = res.total;

      // console.table(res);
    },
  },
};
</script>

<style>
</style>

