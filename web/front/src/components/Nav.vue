<template>
  <!-- 展示一些个人介绍 -->
  <v-card
    max-width="300"
    elevation="5"
  >
    <!-- title 展示头像及名字 -->
    <v-img src="@/assets/img/NavBG.jpg">
      <v-card-title>
        <v-col align="center">
          <v-avatar
            size="100"
            color="grey"
          >
            <img :src="profile.avatar" />
          </v-avatar>
          <div class="ma-4 white--text">{{profile.name}}</div>
        </v-col>
      </v-card-title>

      <v-divider></v-divider>

      <!-- text 展示个人介绍 -->
      <v-col class="au">
        <div class="ma-3">关于作者:</div>
        <div class="ma-3">{{profile.state}}</div>
        <div class="ma-3">{{profile.desc}}</div>
      </v-col>

      <v-divider></v-divider>

      <!-- v-list 的默认白色背景会覆盖掉图片背景 -->
      <v-list
        dense
        nav
      >
        <v-list-item>
          <v-list-item-icon>
            <v-icon>{{'mdi-qqchat'}}</v-icon>
          </v-list-item-icon>
          <v-list-item-content class="grey-text">{{profile.qqchat}}</v-list-item-content>
        </v-list-item>

        <v-list-item>
          <v-list-item-icon>
            <v-icon>{{'mdi-wechat'}}</v-icon>
          </v-list-item-icon>
          <v-list-item-content class="grey-text">123456789</v-list-item-content>
        </v-list-item>

        <v-list-item>
          <v-list-item-icon>
            <v-icon>{{'mdi-sina-weibo'}}</v-icon>
          </v-list-item-icon>
          <v-list-item-content class="grey-text">{{profile.weibo}}</v-list-item-content>
        </v-list-item>

        <v-list-item>
          <v-list-item-icon>
            <v-icon>{{'mdi-email'}}</v-icon>
          </v-list-item-icon>
          <v-list-item-content class="grey-text">{{profile.email}}</v-list-item-content>
        </v-list-item>
      </v-list>
    </v-img>
  </v-card>
</template>

<script>
export default {
  name: "NavBar",
  data() {
    return {
      profile: {},
    };
  },
  created() {
    this.getProfile();
  },
  methods: {
    async getProfile(param) {
      // if (id == null || id == undefined || id <= 0) id = 1; //使用 ?? 简化空值判断
      let id = param ?? 1;

      const { data: res } = await this.$axios.get(`/profile/${id}`);
      this.profile = res.data;
      console.log(this.profile);
    },
  },
};
</script>

<style>
.au {
  font-size: medium;
  background-color: rgba(0, 0, 0, 0.5);
  color: #fff;
}
</style>

