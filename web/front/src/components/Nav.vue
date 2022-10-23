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
            <img src="@/assets/img/avatar.jpg" />
          </v-avatar>
          <div class="ma-4 white--text">Joey</div>
        </v-col>
      </v-card-title>

      <v-divider></v-divider>

      <!-- text 展示个人介绍 -->
      <v-col class="au">
        <div class="ma-3">关于作者:</div>
        <div class="ma-3">电子信息研究生在读</div>
        <div class="ma-3">一名 Java、Go、JS 开发者，对任何新技术永远抱有热情与求知欲</div>
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
          <v-list-item-content class="grey-text">123456789</v-list-item-content>
        </v-list-item>
        <v-list-item>
          <v-list-item-icon>
            <v-icon>{{'mdi-sina-weibo'}}</v-icon>
          </v-list-item-icon>
          <v-list-item-content class="grey-text">123456789</v-list-item-content>
        </v-list-item>
        <v-list-item>
          <v-list-item-icon>
            <v-icon>{{'mdi-email'}}</v-icon>
          </v-list-item-icon>
          <v-list-item-content class="grey-text">123456789@outlook.com</v-list-item-content>
        </v-list-item>

        <v-list-item>
          <v-list-item-icon>
            <v-icon>{{'mdi-github'}}</v-icon>
          </v-list-item-icon>
          <v-list-item-content class="grey-text">123456789</v-list-item-content>
        </v-list-item>
      </v-list>
    </v-img>
  </v-card>
</template>

<script>
export default {
  name: "NavBar",
  data() {
    return {};
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

