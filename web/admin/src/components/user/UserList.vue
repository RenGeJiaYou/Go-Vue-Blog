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

      <a-table
        :dataSource="userList"
        :columns="columns"
        :pagination="paginationOption"
        bordered
        rowKey="ID"
      >
        <!--↓↓↓范围内是父组件 => 子组件的数据，父组件应该用slot表明数据要传入子组件的哪个部分；而子组件<a-table>早已用columns[]所定义的key指定具名插槽↓↓↓------------------------------------------------->
        <!--希望将 role的1、2换成管理员、订阅者 -->
        <span
          slot="role"
          slot-scope="data"
        >{{data==1?'管理员':'订阅者'}}</span>

        <template
          slot="action"
          slot-scope="data"
        >
          <div class="actionSlot">
            <a-button
              type="primary"
              style="margin-right: 15px"
            >编辑</a-button>
            <a-button type="danger">删除</a-button>
          </div>
        </template>
        <!--↑↑↑范围内是父组件 => 子组件的数据，父组件应该用slot表明数据要传入子组件的哪个部分；而子组件<a-table>早已用columns[]所定义的key指定具名插槽↑↑↑------------------------------------------------->
      </a-table>
    </a-card>
  </div>
</template>
    
<script>
const columns = [
  {
    title: "ID",
    dataIndex: "ID", //接收来自json的数据，和json字段一致
    sorter: true,
    width: "10%",
    align: "center",
    key: "id",

    scopedSlots: { customRender: "id" },
  },
  {
    title: "用户名",
    dataIndex: "username",
    align: "center",
    width: "20%",
    key: "username",
  },
  {
    title: "角色",
    dataIndex: "role",
    width: "20%",
    align: "center",
    key: "role",
    scopedSlots: { customRender: "role" }, //指明要渲染的是slot="role"的那一行
  },
  {
    title: "操作",
    width: "30%",
    align: "center",
    key: "action",
    scopedSlots: { customRender: "action" },
  },
];

export default {
  name: "AdminIndex",

  data() {
    return {
      //用户列表的表单数据
      userList: [],

      //规定列格式
      columns,

      //分页器配置
      paginationOption: {
        pageSizeOptions: ["5", "10", "15", "20"], //指定每页可以显示多少条

        defaultCurrent: 1, //默认的当前页数
        defaultPageSize: 5, ///默认的每页条数

        total: 0, //数据总量
        showSizeChanger: true, //允许更改每页条数
        showTotal: (total, range) => {
          //拿到总数，控制分页器的描述字符串具体内容
          //range[0]：当前页码的第一条数据的序号
          //range[1]：当前页码的最后一条的序号
          `${range[0]}-${range[1]}  共 ${total} 条数据`;
        },

        //分页器的时间直接写到这里时，函数名前要加'on'
        onChange: (page, pageSize) => {
          //页码改变的回调
          this.paginationOption.defaultCurrent = page;
          this.paginationOption.defaultPageSize = pageSize;
        },
        // onShowSizeChange: (current, size) => {
        //   //pageSize变化的回调
        //   //current: 当前页数
        //   //size:当前尺寸
        //   //todo 确定是这两个default 而不是current/pageSize 吗？
        //   this.paginationOption.defaultCurrent = current;
        //   this.paginationOption.defaultPageSize = size;
        // },
      },
    };
  },

  created() {
    //在页面创建前就获取相应数据
    this.getUserList();
  },

  methods: {
    async getUserList() {
      //将旧res的data部分又赋给res.即只保留 res 的 data 部分
      const { data: res } = await this.$axios.get("users", {
        params: {
          pagenum: this.paginationOption.defaultCurrent,
          pagesize: this.paginationOption.defaultPageSize,
        },
      });

      if (res.status != 200) return this.$message.error(res.message);

      //确认无误，赋值
      this.userList = res.data;
      this.paginationOption.total = res.total;
    },
    onSearch() {},
  },
};
</script>
    
<style scoped>
.actionSlot {
  display: flex;
  justify-content: center;
}
</style>
