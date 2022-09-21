<template>
  <div>
    <a-card>
      <a-row :gutter="16">
        <a-col :span="6">
          <a-input-search
            placeholder="输入用户名查找"
            enter-button
            v-model="queryParam.username"
            @search="getUserList"
            allowClear
          />
        </a-col>
        <a-col :span="4">
          <!--vue文件中，<template>调用数据不用像<script>那样加this.-->
          <a-button
            type="primary"
            @click="addUserVisible = true"
          >新增用户</a-button>
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
        <!--1.【组件传值：父→子】“管理员”/“订阅者” 这两个数据是属于 父组件UserList.vue 的，想传给子组件显示。使用具名插槽slot 和子组件<a-table>的columns[] 定义的 scopedSlots 所声明的具名插槽值一致 -->
        <!--2.【组件传值：子→父】当前行的role值是 子组件<a-table> 用自己的 dataSource 得到的，想传给父组件（用于验证）。使用slot-scoped 在父组件调用时产生一个接收数据的接口 -->

        <span
          slot="role"
          slot-scope="role"
        >{{role==1?'管理员':'订阅者'}}</span>

        <!--1.【组件传值：父→子】UserList.vue调用 <a-table> ,在两处约定具名插槽 "action" ，将<template>内的两个按钮元素传到“操作”一列-->
        <!--2.【组件传值：子→父】经过console 打印测试，可得知data是子组件 <a-table> 的返回数据-->

        <template
          slot="action"
          slot-scope="data"
        >
          <div class="actionSlot">
            <a-button
              type="primary"
              @click="editUser(data)"
              style="margin-right: 15px"
            >编辑</a-button>

            <a-button
              type="danger"
              @click="deleteConfirm(data.ID)"
              style="margin-right: 15px"
            >删除</a-button>

            <a-button @click="resetPassword(data.ID)">重置密码</a-button>
          </div>
        </template>
        <!--↑↑↑范围内是父组件 => 子组件的数据，父组件应该用slot表明数据要传入子组件的哪个部分；而子组件<a-table>早已用columns[]所定义的key指定具名插槽↑↑↑------------------------------------------------->
      </a-table>
    </a-card>

    <!--新增用户 弹框-->
    <a-modal
      title="新增用户"
      :visible="addUserVisible"
      @ok="handleAddUserOK"
      @cancel="handleAddUserCancel"
    >
      <a-form-model
        :model="userInfo"
        :rules="rules"
        ref="addUserRef"
        :label-col="labelCol"
        :wrapper-col="wrapperCol"
      >
        <a-form-model-item
          label="用户名"
          prop="username"
        >
          <a-input v-model="userInfo.username"></a-input>
        </a-form-model-item>
        <a-form-model-item
          label="密码"
          prop="password"
          has-feedback
        >
          <a-input-password v-model="userInfo.password"></a-input-password>
        </a-form-model-item>
        <a-form-model-item
          label="确认密码"
          prop="checkPass"
          has-feedback
        >
          <a-input-password v-model="userInfo.checkPass"></a-input-password>
        </a-form-model-item>
        <a-form-model-item label="是否为管理员">
          <!--此处无需v-model双向绑定，数据的更新可用change回调函数完成-->
          <a-select
            default-value="订阅者"
            style="width: 120px"
            @change="addUserRoleChange"
          >
            <a-select-option
              key="1"
              value="1"
            >管理员</a-select-option>
            <a-select-option
              key="2"
              value="2"
            >订阅者</a-select-option>
          </a-select>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
    <!--/新增用户 弹框-->

    <!--编辑用户 弹框-->
    <a-modal
      title="编辑用户"
      :visible="editUserVisible"
      @ok="handleEditUserOK"
      @cancel="handleEditUserCancel"
    >
      <a-form-model
        :model="userEdit"
        :rules="rules"
        ref="editUserRef"
        :label-col="labelCol"
        :wrapper-col="wrapperCol"
      >
        <a-form-model-item
          label="用户名"
          prop="username"
        >
          <a-input v-model="userEdit.username"></a-input>
        </a-form-model-item>
        <!--
          编辑用户功能 - 修改用户密码功能日后再写
          <a-form-model-item
          label="密码"
          prop="password"
          has-feedback
        >
          <a-input-password v-model="userEdit.password"></a-input-password>
        </a-form-model-item>-->
        <!-- <a-form-model-item
          label="确认密码"
          prop="checkPass"
          has-feedback
        >
          <a-input-password v-model="userEdit.checkPass"></a-input-password>
        </a-form-model-item>-->
        <a-form-model-item label="是否为管理员">
          <!--此处无需v-model双向绑定，数据的更新可用change回调函数完成-->
          <a-select
            :default-value="userEdit.role"
            style="width: 120px"
            @change="editUserRoleChange"
          >
            <a-select-option
              key="1"
              :value="1"
            >管理员</a-select-option>
            <a-select-option
              key="2"
              :value="2"
            >订阅者</a-select-option>
          </a-select>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
    <!--/编辑用户 弹框-->

    <!--编辑用户 弹框-->
    <a-modal
      title="编辑用户"
      :visible="editUserVisible"
      @ok="handleEditUserOK"
      @cancel="handleEditUserCancel"
    >
      <a-form-model
        :model="userEdit"
        :rules="rules"
        ref="editUserRef"
        :label-col="labelCol"
        :wrapper-col="wrapperCol"
      >
        <a-form-model-item
          label="用户名"
          prop="username"
        >
          <a-input v-model="userEdit.username"></a-input>
        </a-form-model-item>
        <!--
          编辑用户功能 - 修改用户密码功能日后再写
          <a-form-model-item
          label="密码"
          prop="password"
          has-feedback
        >
          <a-input-password v-model="userEdit.password"></a-input-password>
        </a-form-model-item>-->
        <!-- <a-form-model-item
          label="确认密码"
          prop="checkPass"
          has-feedback
        >
          <a-input-password v-model="userEdit.checkPass"></a-input-password>
        </a-form-model-item>-->
        <a-form-model-item label="是否为管理员">
          <!--此处无需v-model双向绑定，数据的更新可用change回调函数完成-->
          <a-select
            :default-value="userEdit.role"
            style="width: 120px"
            @change="editUserRoleChange"
          >
            <a-select-option
              key="1"
              :value="1"
            >管理员</a-select-option>
            <a-select-option
              key="2"
              :value="2"
            >订阅者</a-select-option>
          </a-select>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
    <!--/编辑用户 弹框-->
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

    scopedSlots: { customRender: "id" }, //是Vue子组件内具名插槽<slot name="xxx">写法的封装
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
    scopedSlots: { customRender: "role" }, //子组件的具名插槽，值要和slot="role"一致
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
    let validatePass = (rule, value, callback) => {
      if (value == "") {
        callback(new Error("密码不得为空"));
      } else if (value.length < 6 || value.length > 20) {
        callback(new Error("密码位数应该在 6~20 位之间"));
      } else {
        //默认正确路径。  如果没有这个路径，若之前已进入其它错误路径，之前的错误信息将不会消失。
        callback();
      }
    };

    //密码确认框 校验器
    let validatePass2 = (rule, value, callback) => {
      if (value === "") {
        callback(new Error("请再次输入密码"));
      } else if (value !== this.userInfo.password) {
        callback(new Error("两次输入不匹配!"));
      } else {
        callback();
      }
    };

    return {
      //用户列表的表单数据
      userList: [],

      //规定列格式
      columns,

      //每页的页码和当前页数,此处数据将传递给getUserList()的 axios 请求
      queryParam: {
        username: "",
        pagesize: 5,
        pagenum: 1,
      },

      //visible: false, //视频说这是删除确认的visible属性。但是官方示例中没有。实测删除后也无影响
      addUserVisible: false, //隐藏/浮现添加用户弹框
      editUserVisible: false, //隐藏/浮现编辑用户弹框

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
          return `共 ${total} 条数据  `;
        },
      },

      //添加用户弹框绑定的数据对象
      userInfo: {
        id: 0,
        username: "",
        password: "",
        checkPass: "",
        role: 2,
      },
      //编辑用户弹框绑定的数据对象
      userEdit: {
        id: 0,
        username: "",
        password: "",
        checkPass: "",
        role: 0,
      },

      //添加&编辑用户弹框文本宽度和控件宽度
      labelCol: { span: 6 },
      wrapperCol: { span: 14, offset: 1 },
      //添加&编辑用户弹框表单验证
      rules: {
        //两种方式：使用验证器+自定义回调函数 和 使用预设校验规则
        //效果是一样的。此处分别写出以供学习
        username: [
          {
            required: true,
            message: "请输入用户名",
            trigger: "blur",
          },
          {
            min: 4,
            max: 12,
            message: "长度保持在4-12个字符",
            trigger: "blur",
          },
        ],
        password: [{ validator: validatePass, trigger: "change" }],
        checkPass: [{ validator: validatePass2, trigger: "change" }],
      },
    };
  },

  created() {
    //在页面创建前就获取相应数据
    this.getUserList();
  },

  methods: {
    //获取列表数据(囊括带搜索条件&&无搜索条件)
    // @@重构tip,@search 提供了一个回调函数(value,event)，包含了搜索值，试试能不能用
    async getUserList() {
      //将所接收数据的data部分赋给res.
      const { data: res } = await this.$axios.get("users", {
        params: {
          username: this.queryParam.username, //为""时说明是无搜索条件，不为空时说明带搜索条件
          pagenum: this.queryParam.pagenum,
          pagesize: this.queryParam.pagesize,
        },
      });

      if (res.status != 200) return this.$message.error(res.message);

      //确认无误，保存返回结果
      this.userList = res.data;
      this.pagination.total = res.total;
      // console.log("实际返回的数据条数", res.total);
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
    //==========================================
    //删除当前用户
    async deleteUser(id) {
      // 子组件<a-table>所提供给父组件的值 data 是当前行数据，即getUserList()的结果数组中的当前一条，字段包括
      /*  CreatedAt
          DeletedAt
          ID
          UpdatedAt
          password
          role
          username
      */
      // console.log("<a-table>向父组件的传输值:", data);
      const res = await this.$axios.delete(`user/${id}`);
      if (res.status !== 200) {
        return this.$message.error(res.message);
      }
      this.$message.success("删除成功");

      //5.删除后刷新列表
      this.getUserList();
    },
    //删除确认,调用deleteUser()
    deleteConfirm(id) {
      this.$confirm({
        title: "确定删除该项吗?",
        content: "删除后将不可恢复",
        okText: "确认",
        okType: "primary",
        cancelText: "取消",
        onOk: () => {
          this.deleteUser(id); //使用箭头函数，this 指向当前Vue实例，而非调用该 this 的函数
        },
        onCancel: () => {
          this.$message.info("已取消");
        },
      });
    },
    //==========================================
    //添加用户弹窗 确认按钮
    handleAddUserOK() {
      //在 validate 中完成全部操作
      this.$refs.addUserRef.validate(async (valid) => {
        if (!valid) return this.$message.error("未通过验证，请重新填写"); //回调函数的实现总要写 return
        //axios 调用API
        const { data: res } = await this.$axios.post("user/add", {
          username: this.userInfo.username,
          password: this.userInfo.password,
          role: this.userInfo.role,
        });
        if (res.status !== 200) return this.$message.error(res.message);

        this.addUserVisible = false; //关闭弹窗
        this.$refs.addUserRef.resetFields(); //清空数据。组件提供的方法可通过ref的方式调用
        this.$message.success("添加用户成功"); //成功提示
        this.getUserList(); //数据更新后要及时刷新
      });
    },
    //添加用户弹窗 取消按钮
    handleAddUserCancel() {
      //清空 model{} 中废弃的数据

      this.addUserVisible = false;
      this.$refs.addUserRef.resetFields();
      this.$message.info("已取消添加");
    },
    //添加用户弹窗 "是否为管理员"多选框变更事件
    addUserRoleChange(val) {
      //将更新后的值传到 userInfo 中，以待传到后端
      this.userInfo.role = val;
      console.log(this.userInfo.role);
    },
    //==========================================
    //编辑用户按钮
    editUser(data) {
      // console.log(data); //data包含了当前行所有数据

      //1.将 data 导入v-model数据模型
      this.userEdit.id = data.ID;
      this.userEdit.username = data.username;
      this.userEdit.password = data.password;
      this.userEdit.role = data.role;

      //2.打开弹框
      this.editUserVisible = true;

      //3.此时用户编辑，修改数据。响应式地更新到模型里
    },
    //编辑用户弹窗 确认按钮
    handleEditUserOK() {
      //在 validate 中完成全部操作
      this.$refs.editUserRef.validate(async (valid) => {
        if (!valid) return this.$message.error("未通过验证，请重新填写"); //回调函数的实现总要写 return

        //axios 调用API，只能修改用户名和身份，修改密码是独立的功能点
        const { data: res } = await this.$axios.put(
          `user/${this.userEdit.id}`,
          {
            username: this.userEdit.username,
            role: this.userEdit.role,
          }
        );

        if (res.status !== 200) return this.$message.error(res.message);
        this.editUserVisible = false; //关闭弹窗
        // this.$refs.editUserRef.resetFields(); //无需清空数据。反正每次打开编辑弹窗都会重新加载本行数据
        this.$message.success("编辑用户成功"); //成功提示
        this.getUserList(); //数据更新后要及时刷新
      });
    },
    //编辑用户弹窗 取消按钮
    handleEditUserCancel() {
      this.editUserVisible = false;
      // this.$refs.editUserRef.resetFields();//无需清空数据，反正每次点击编辑都会重新加载数据
      this.$message.info("已取消编辑");
    },
    //编辑用户弹窗 "是否为管理员"多选框变更事件
    editUserRoleChange(val) {
      //将更新后的值传到 userInfo 中，以待传到后端
      this.userEdit.role = val;
    },
    //==========================================
    // 重置密码按钮
    resetPassword(id) {
      /*
      密码安全校验（配合 rules 使用）
      获取数据向后端发送 PUT 请求
      $message.error()/success()
      */
    },
  },
};
</script>
    
<style scoped>
.actionSlot {
  display: flex;
  justify-content: center;
}
</style>
