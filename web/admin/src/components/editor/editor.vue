<template>
  <div class>
    <Editor :init="init"></Editor>
  </div>
</template>

<script>
import Editor from "@tinymce/tinymce-vue"; //tinymce 的官方 npm 包
import tinymce from "./tinymce.min.js"; //tinymce 的本地包
import "./icons/default/icons.min.js";
import "./themes/silver/theme.min.js";
import "./langs/zh-Hans.js";
export default {
  name: "",
  components: { Editor },
  data() {
    return {
      init: {
        language: "zh-CN",
      },
    };
  },
};
</script>

<style>
@import url("./skins/ui/oxide/skin.min.css");
</style>

