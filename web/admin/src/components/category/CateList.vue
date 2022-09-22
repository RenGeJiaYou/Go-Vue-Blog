<template>
  <div>
    <a-card>
      <!--表头部分-->
      <a-row :gutter="16">
        <a-col :span="6">
          <a-input-search
            placeholder="输入分类名查找"
            enter-button
            v-model="queryParam.name"
            @search="getCateList"
            allowClear
          />
        </a-col>
        <a-col :span="4">
          <a-button
            type="primary"
            @click="addCateVisible = true"
          >新增分类</a-button>
        </a-col>
      </a-row>
      <!--/表头部分-->

      <!--表单部分-->
      <a-table
        :dataSource="cateList"
        :columns="columns"
        :pagination="pagination"
        bordered
        rowKey="ID"
        @change="handleTableChange"
      >
        <template
          slot="action"
          slot-scope="data"
        >
          <div class="actionSlot">
            <a-button
              type="primary"
              @click="editCate(data)"
              style="margin-right: 15px"
            >编辑</a-button>

            <a-button
              type="danger"
              @click="deleteConfirm(data.id)"
              style="margin-right: 15px"
            >删除</a-button>
          </div>
        </template>
      </a-table>
      <!--/表单部分-->
    </a-card>

    <!--新增分类 弹框-->
    <a-modal
      title="新增分类"
      :visible="addCateVisible"
      @ok="handleAddCateOK"
      @cancel="handleAddCateCancel"
    >
      <a-form-model
        :model="cateInfo"
        :rules="rules"
        ref="addCateRef"
        :label-col="labelCol"
        :wrapper-col="wrapperCol"
      >
        <a-form-model-item
          label="分类名"
          prop="name"
        >
          <a-input v-model="cateInfo.name"></a-input>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
    <!--/新增分类 弹框-->

    <!--编辑分类 弹框-->
    <a-modal
      title="编辑分类"
      :visible="editCateVisible"
      @ok="handleEditCateOK"
      @cancel="handleEditCateCancel"
    >
      <a-form-model
        :model="cateEdit"
        :rules="rules"
        ref="editCateRef"
        :label-col="labelCol"
        :wrapper-col="wrapperCol"
      >
        <a-form-model-item
          label="分类名"
          prop="name"
        >
          <a-input v-model="cateEdit.name"></a-input>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
    <!--/编辑分类 弹框-->
  </div>
</template>
    
<script>
const columns = [
  {
    title: "ID",
    dataIndex: "id", //接收来自json的数据，和json字段一致
    sorter: true,
    width: "10%",
    align: "center",
    key: "id",

    scopedSlots: { customRender: "id" }, //是Vue子组件内具名插槽<slot name="xxx">写法的封装
  },
  {
    title: "分类名",
    dataIndex: "name",
    align: "center",
    width: "30%",
    key: "name",
  },
  {
    title: "操作",
    width: "20%",
    align: "center",
    key: "action",
    scopedSlots: { customRender: "action" },
  },
];

export default {
  name: "AdminIndex",

  data() {
    return {
      //分类列表的表单数据
      cateList: [],

      //规定列格式
      columns,

      //每页的页码和当前页数,此处数据将传递给getCateList()的 axios 请求
      queryParam: {
        name: "",
        pagesize: 5,
        pagenum: 1,
      },

      //visible: false, //视频说这是删除确认的visible属性。但是官方示例中没有。实测删除后也无影响
      addCateVisible: false, //隐藏/浮现添加分类弹框
      editCateVisible: false, //隐藏/浮现编辑分类弹框

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

      //添加分类弹框绑定的数据对象
      cateInfo: {
        id: 0,
        name: "",
      },
      //编辑分类弹框绑定的数据对象
      cateEdit: {
        id: 0,
        name: "",
      },
      //添加&编辑分类弹框文本宽度和控件宽度
      labelCol: { span: 6 },
      wrapperCol: { span: 14, offset: 1 },
      //添加&编辑分类弹框表单验证
      rules: {
        name: [
          {
            required: true,
            message: "请输入分类名",
            trigger: "blur",
          },
          {
            min: 1,
            max: 12,
            message: "长度保持在1-12个字符",
            trigger: "blur",
          },
        ],
      },
    };
  },

  created() {
    //在页面创建前就获取相应数据
    this.getCateList();
  },

  methods: {
    //获取列表数据(囊括带搜索条件&&无搜索条件)
    // @@重构tip,@search 提供了一个回调函数(value,event)，包含了搜索值，试试能不能用
    async getCateList() {
      //将所接收数据的data部分赋给res.
      const { data: res } = await this.$axios.get("categories", {
        params: {
          name: this.queryParam.name, //为""时说明是无搜索条件，不为空时说明带搜索条件
          pagenum: this.queryParam.pagenum,
          pagesize: this.queryParam.pagesize,
        },
      });

      if (res.status != 200) return this.$message.error(res.message);

      //确认无误，保存返回结果
      this.cateList = res.data;
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
      //用于保存分类操作后的新数据，并传递到 this.pagination 中，以完成页面的更新。
      let pager = { ...this.pagination };

      //1.【分类界面】拿到更新的值（由回调函数所给的参数 pagination 提供）
      pager.current = pagination.current;
      pager.pageSize = pagination.pageSize;

      //2.【后端数据请求函数】拿到更新的值（由回调函数所给的参数 pagination 提供）
      this.queryParam.pagenum = pagination.current;
      this.queryParam.pagesize = pagination.pageSize;

      //3.如果更新后的每页条目！==之前的每页条目，【分类界面】和【后端数据】都回到第一页
      if (pagination.pageSize !== this.pagination.pageSize) {
        pager.current = 1;
        this.queryParam.pagenum = 1;
      }

      //4.更新【分类界面】
      this.pagination = pager;

      //5.请求【后端数据】
      this.getCateList();
    },
    //==========================================
    //删除当前分类
    async deleteCate(id) {
      console.log("删除函数的id的传输值:", id);
      const res = await this.$axios.delete(`category/${id}`);
      if (res.status !== 200) {
        return this.$message.error(res.message);
      }
      this.$message.success("删除成功");

      //5.删除后刷新列表
      this.getCateList();
    },
    //删除确认,调用deleteCate()
    deleteConfirm(id) {
      this.$confirm({
        title: "确定删除该项吗?",
        content: "删除后将不可恢复",
        okText: "确认",
        okType: "primary",
        cancelText: "取消",
        onOk: () => {
          this.deleteCate(id); //使用箭头函数，this 指向当前Vue实例，而非调用该 this 的函数
        },
        onCancel: () => {
          this.$message.info("已取消");
        },
      });
    },
    //==========================================
    //添加分类弹窗 确认按钮
    handleAddCateOK() {
      //在 validate 中完成全部操作
      this.$refs.addCateRef.validate(async (valid) => {
        if (!valid) return this.$message.error("未通过验证，请重新填写"); //回调函数的实现总要写 return
        //axios 调用API
        const { data: res } = await this.$axios.post("category/add", {
          name: this.cateInfo.name,
        });
        if (res.status !== 200) return this.$message.error(res.message);

        this.addCateVisible = false; //关闭弹窗
        this.$refs.addCateRef.resetFields(); //清空数据。组件提供的方法可通过ref的方式调用
        this.$message.success("添加分类成功"); //成功提示
        this.getCateList(); //数据更新后要及时刷新
      });
    },
    //添加分类弹窗 取消按钮
    handleAddCateCancel() {
      //清空 model{} 中废弃的数据

      this.addCateVisible = false;
      this.$refs.addCateRef.resetFields();
      this.$message.info("已取消添加");
    },
    //==========================================
    //编辑分类按钮
    editCate(data) {
      // console.log(data); //data包含了当前行所有数据

      //1.将 data 导入v-model数据模型
      this.cateEdit.id = data.id;
      this.cateEdit.name = data.name;

      //2.打开弹框
      this.editCateVisible = true;

      //3.此时分类编辑，修改数据。响应式地更新到模型里
    },
    //编辑分类弹窗 确认按钮
    handleEditCateOK() {
      //在 validate 中完成全部操作
      this.$refs.editCateRef.validate(async (valid) => {
        if (!valid) return this.$message.error("未通过验证，请重新填写"); //回调函数的实现总要写 return

        //axios 调用API，只能修改分类名和身份，修改密码是独立的功能点
        const { data: res } = await this.$axios.put(
          `category/${this.cateEdit.id}`,
          {
            name: this.cateEdit.name,
          }
        );

        if (res.status !== 200) return this.$message.error(res.message);
        this.editCateVisible = false; //关闭弹窗
        this.$message.success("编辑分类成功"); //成功提示
        this.getCateList(); //数据更新后要及时刷新
      });
    },
    //编辑分类弹窗 取消按钮
    handleEditCateCancel() {
      this.editCateVisible = false;
      this.$message.info("已取消编辑");
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
