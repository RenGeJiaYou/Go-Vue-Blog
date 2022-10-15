<template>
  <!--关于v-app、v-main、app 属性等，参见 https://vuetifyjs.com/zh-Hans/features/layouts/-->
  <v-app app>
    <TopBar></TopBar>

    <v-main class="grey lighten-4">
      <v-container>
        <v-row>
          <v-col
            cols="3"
            md="3"
          >
            <Nav></Nav>
          </v-col>
          <v-col>
            <v-sheet
              rounded="lg"
              elevation="5"
              min-height="85vh"
            >
              <router-view></router-view>
            </v-sheet>
          </v-col>
        </v-row>
      </v-container>
    </v-main>

    <Footer></Footer>
  </v-app>
</template>

<script>
import TopBar from "@/components/TopBar.vue";
import Footer from "@/components/Footer.vue";
import Nav from "@/components/Nav.vue";
import Article from "@/components/Article.vue";
export default {
  name: "Home",

  components: { TopBar, Footer, Nav, Article },
};
</script>
