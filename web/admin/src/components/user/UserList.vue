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
        :pagination="pagination"
        bordered
        rowKey="ID"
        @change="handleTableChange"
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

      //每页的页码和当前页数,此处数据将传递给getUserList()的 axios 请求
      queryParam: {
        pagesize: 5,
        pagenum: 1,
      },

      visible: false,

      //分页器配置
      pagination: {
        pageSizeOptions: ["5", "10", "20"], //指定每页可以显示多少条
        pageSize: 5, ///每页条数
        current: 1, //当前页码
        total: 0, //数据总量
        showSizeChanger: true, //允许更改每页条数
        showTotal: (total, range) => {
          //拿到总数，控制分页器的描述字符串具体内容
          //range[0]：当前页码的第一条数据的序号
          //range[1]：当前页码的最后一条的序号
          `${range[0]}-${range[1]}  共 ${total} 条数据`;
        },
      },
    };
  },

  created() {
    //在页面创建前就获取相应数据
    this.getUserList();
  },

  methods: {
    //获取列表数据
    async getUserList() {
      //将旧res的data部分又赋给res.即只保留 res 的 data 部分
      const { data: res } = await this.$axios.get("users", {
        params: {
          pagenum: this.queryParam.pagenum,
          pagesize: this.queryParam.pagesize,
        },
      });

      if (res.status != 200) return this.$message.error(res.message);

      //确认无误，保存返回结果
      this.userList = res.data;
      this.pagination.total = res.total;
    },

    //Table 分页、排序、筛选变化时触发(onChange方法名已存在于 antdv API,遂更名)
    handleTableChange(pagination, filters, sorter) {
      /* pagination 的主要字段：
         current:当前页码
         pageSize:每页条目
         total: 数据总量
      */

      // pager 是一个临时变量。
      //用于保存用户操作后的新数据，并传递到 this.pagination 中，以完成页面的更新。
      let pager = { ...this.pagination };

      //1.【用户界面】拿到更新的值（由回调函数所给的参数 pagination 提供）
      pager.current = pagination.current;
      pager.pageSize = pagination.pageSize;

      //2.【后端数据请求函数】拿到更新的值（由回调函数所给的参数 pagination 提供）
      this.queryParam.pagenum = pagination.current;
      this.queryParam.pagesize = pagination.pageSize;

      //3.如果更新后的每页条目！==之前的每页条目，【用户界面】和【后端数据】都回到第一页
      if (pagination.pageSize !== this.pagination.pageSize) {
        pager.current = 1;
        this.queryParam.pagenum = 1;
      }

      //4.更新【用户界面】
      this.pagination = pager;

      //5.请求【后端数据】
      this.getUserList();
    },

    //查找用户事件
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
