<template>
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
