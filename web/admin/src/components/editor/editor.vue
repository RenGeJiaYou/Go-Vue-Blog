<template>
  <div class>
    <Editor
      :init="init"
      v-model="content"
    ></Editor>
  </div>
</template>

<script>
import Editor from "@tinymce/tinymce-vue"; //tinymce 的官方 npm 包
import tinymce from "./tinymce.min.js"; //tinymce 的本地包
import "./icons/default/icons.min.js";
import "./themes/silver/theme.min.js";
import "./langs/zh-Hans.js";

//注册插件
import "./plugins/preview/plugin.min.js"; // 预览
import "./plugins/paste/plugin.min.js";
import "./plugins/wordcount/plugin.min.js"; // 字数统计
import "./plugins/code/plugin.min.js"; // 代码块

import "./plugins/image/plugin.min.js";
import "./plugins/imagetools/plugin.min.js";
import "./plugins/media/plugin.min.js";
import "./plugins/codesample/plugin.min.js";
import "./plugins/lists/plugin.min.js";
import "./plugins/table/plugin.min.js";

export default {
  name: "",
  components: { Editor },
  props: {
    value: {
      type: String,
      default: "",
    },
  },
  data() {
    return {
      //富文本编辑器的初始化配置
      init: {
        language: "zh_CN",
        height: "500px",
        //加载组件
        plugins:
          "preview paste wordcount code imagetools image media codesample lists table",
        margin: "0",
        padding: "0",
        branding: false, //隐藏右下角'powered by tinymce' 字样
        //在toolbar 显示这些组件
        toolbar: [
          "undo redo | styleselect |fontsizeselect| bold italic underline strikethrough |alignleft aligncenter alignright alignjustify |blockquote removeformat |numlist bullist table",
          "preview paste code codesample |image media",
        ],
        // 自定义图片上传函数
        images_upload_handler: async (blobInfo, success, failure) => {
          let formData = new FormData();
          formData.append("file", blobInfo.blob(), blobInfo.filename()); //不管是官方 doc 的 filename(),还是B站视频的name().两种函数结果一致

          const { data: res } = await this.$axios.post("upload", formData);

          success(res.url);
          failure(this.$message.error("tinymce 中上传图片失败"));
        },
        imagetools_cors_hosts: ["*"],
        imagetools_proxy: "*",
      },
      // 接收来自 AddArt.vue 的文本内容，并让<Editor>组件双向绑定该内容
      content: this.value,
    };
  },
  watch: {
    value(newVal) {
      this.content = newVal;
    },
    content(newVal) {
      this.$emit("input", newVal);
    },
  },
};
</script>

<style>
@import url("./skins/ui/oxide/skin.min.css");
@import url("./skins/content/default/content.min.css");
</style>

