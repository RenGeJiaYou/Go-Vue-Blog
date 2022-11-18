<template>
  <v-col>
    <v-card
      v-for=" i in articleList"
      :key="i.ID"
      link
      @click="$router.push(`details/${i.ID}`)"
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
          <v-card-text v-html="i.content.slice(0,100)"></v-card-text>
        </v-col>
      </v-row>
    </v-card>

    <div class="text-center">
      <v-pagination
        total-visible="7"
        v-model="queryParam.pagenum"
        :length="Math.ceil(total/queryParam.pagesize)"
        circle
        @input="getArtList"
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
    },
  },
  filters: {
    dateFilter: function (value) {
      //先将字符串转换为 Date 的对象
      let dateValue = new Date(value);
      let hours = dateValue.getHours();
      let minutes = dateValue.getMinutes();

      if (dateValue.getHours() < 10) {
        hours = "0" + dateValue.getHours();
      }
      if (dateValue.getMinutes() < 10) {
        minutes = "0" + dateValue.getMinutes();
      }
      return (
        dateValue.getFullYear() +
        "年" +
        (dateValue.getMonth() + 1) +
        "月" +
        dateValue.getDate() +
        "日" +
        "  " +
        hours +
        ":" +
        minutes
      );
    },
  },
};
</script>

<style>
</style>

