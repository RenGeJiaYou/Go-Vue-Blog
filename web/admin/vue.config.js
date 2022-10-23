// defineConfig 是 vue 提供的帮手函数，具有提示功能
const { defineConfig } = require("@vue/cli-service");

module.exports = defineConfig({
  transpileDependencies: true,
  //关闭eslint校验
  lintOnSave: false,

  //后台页面统一部署到 '/admin/' url下,这意味着后台页面访问时 url 为'http://www.xxx.com/admin/yyyyyyy'
  publicPath: "/admin/",

  //存放打包结果的目录
  outputDir: "dist",

  //存放图片、音频等静态资源的目录
  assetsDir: "static",
});
