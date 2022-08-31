<template>
  <div>
    <a-card>
      <a-row :gutter="16">
        <a-col :span="6">
          <a-input-search
            placeholder="输入用户名查找"
            enter-button
            @search="onSearch"
          />
        </a-col>
        <a-col :span="4">
          <a-button type="primary">新增用户</a-button>
        </a-col>
      </a-row>

      <a-table></a-table>
    </a-card>
  </div>
</template>
    
    <script>
export default {
  name: "AdminIndex",

  data() {
    return {};
  },

  mounted() {},

  methods: {},
};
</script>
    
    <style scoped>
</style>
