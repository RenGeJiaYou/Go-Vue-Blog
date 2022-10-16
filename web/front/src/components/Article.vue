<template>
  <v-col class="grey lighten-4">
    <v-card
      v-for=" i in articleList"
      :key="i.ID"
    >
      <v-row class="mb-6 mt-2">
        <v-col
          cols="3"
          class="d-flex justify-center align-center mx-3"
        >
          <img
            width="100%"
            :src="i.image"
          />
        </v-col>
        <v-col>
          <div>
            <v-card-title>
              <v-chip
                color="pink lighten-3"
                label
                class="white--text mr-5"
              >{{i.Category.name}}</v-chip>
              {{i.title}}
            </v-card-title>
            <v-icon
              dense
              class="ml-16 mr-2"
            >{{'mdi-calendar-month'}}</v-icon>
            <span class="grey--text text--darken-1 caption">{{i.CreatedAt | dateFilter}}</span>
          </div>
          <v-card-subtitle>{{i.desc}}</v-card-subtitle>
          <v-divider
            inset
            light
          ></v-divider>
          <v-card-text>{{i.content}}</v-card-text>
        </v-col>
      </v-row>
    </v-card>

    <div class="text-center">
      <v-pagination
        v-model="page"
        :length="currentPageNum"
        circle
      ></v-pagination>
    </div>
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
      queryParam: {
        // title: "",
        pagesize: 5,
        pagenum: 1,
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
          pagenum: this.queryParam.pagenum,
          pagesize: this.queryParam.pagesize,
        },
      });
      this.articleList = res.data;
      this.total = res.total;

      console.log(this.total);
      // console.table(res);
    },
  },
  computed: {
    //返回当前应该有多少页
    currentPageNum() {
      return this.total % this.queryParam.pagesize == 0
        ? this.total / this.queryParam.pagesize
        : this.total / this.queryParam.pagesize + 1; //有余数，分页器加 1 页
    },
  },
  filters: {
    dateFilter: function (value) {
      //先将字符串转换为 Date 的对象
      let dateValue = new Date(value);
      return (
        dateValue.getFullYear() +
        "年" +
        (dateValue.getMonth() + 1) +
        "月" +
        dateValue.getDate() +
        "日" +
        "  " +
        dateValue.getHours() +
        ":" +
        dateValue.getMinutes()
      );
    },
  },
};
</script>

<style>
</style>

