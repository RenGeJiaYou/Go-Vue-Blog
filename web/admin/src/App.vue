<template>
  <a-locale-provider :locale="locale">
    <div id="app">
      <!--如果 App.vue 没有该标签， Vue Router 引导的所有 view 均不能展现-->
      <router-view />
    </div>
  </a-locale-provider>
</template>

<script>
import zhCN from "ant-design-vue/lib/locale-provider/zh_CN";
export default {
  data() {
    return {
      locale: zhCN,
    };
  },
};
</script>



