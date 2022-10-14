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
        <div>关于作者:</div>
      </v-col>

      <v-divider></v-divider>

      <v-list
        dense
        nav
      >
        <v-list-item>
          <v-list-item-icon>
            <v-icon>{{'mdi-qqchat'}}</v-icon>
          </v-list-item-icon>
          <v-list-item-content>1178972961</v-list-item-content>
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

