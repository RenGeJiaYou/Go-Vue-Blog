<template>
  <!--关于v-app、v-main、app 属性等，参见 https://vuetifyjs.com/zh-Hans/features/layouts/-->
  <v-app app>
    <TopBar></TopBar>
    <v-main>
      <v-container></v-container>
    </v-main>
    <FooterBar></FooterBar>
  </v-app>
</template>

<script>
import TopBar from "@/components/TopBar.vue";
import FooterBar from "@/components/Footer.vue";
export default {
  name: "Home",

  components: { TopBar, FooterBar },
};
</script>
