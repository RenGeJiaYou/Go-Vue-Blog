<template>
  <div>
    <v-footer
      color="primary lighten-1"
      padless
    >
      <v-row
        justify="center"
        no-gutters
      >
        <v-col
          class="primary lighten-2 py-4 text-center white--text"
          cols="12"
        >
          {{ new Date().getFullYear() }} —
          <strong>Go & Vue Blog Project</strong>

          <br />
Powered by Shi Jiajun
        </v-col>
      </v-row>
    </v-footer>
  </div>
</template>

<script>
export default {
  name: "FooterBar",
  data() {
    return {};
  },
};
</script>

<style>
</style>

