<template>
  <div>
    <a-card>
      <!--表头部分-->
      <a-row :gutter="16">
        <a-col :span="6">
          <a-input-search
            placeholder="输入文章标题查找"
            enter-button
            v-model="queryParam.title"
            @search="getArtList"
            allowClear
          />
        </a-col>
        <a-col :span="4">
          <!--vue文件中，<template>调用数据不用像<script>那样加this.-->
          <a-button
            type="primary"
            @click="handleAddArt"
          >写文章</a-button>
        </a-col>

        <a-col :span="8">
          <a-select
            show-search
            allowClear
            placeholder="按分类查找"
            style="width: 200px"
            @change="handleCateChange"
          >
            <!--@change 回调函数提供的value 来自下面的 :value 绑定值-->
            <!-- v-for 使用两种循环方式 in/of 没有区别 -->
            <a-select-option
              v-for="item in cateList"
              :key="item.id"
            >{{item.name}}</a-select-option>
          </a-select>
        </a-col>
      </a-row>
      <!--/表头部分-->

      <!--表单部分-->
      <a-table
        :dataSource="articleList"
        :columns="columns"
        :pagination="pagination"
        bordered
        rowKey="ID"
        @change="handleTableChange"
      >
        <span
          class="artImg"
          slot="img"
          slot-scope="img"
        >
          <img
            :src="img"
            class="artImgItem"
          />
        </span>

        <template
          slot="action"
          slot-scope="data"
        >
          <div class="actionSlot">
            <!--push()中的字符串，如果首位是'/'，将定位到绝对路径-->
            <a-button
              type="primary"
              @click="$router.push(`/addart/${data.ID}`)"
              style="margin-right: 15px"
            >编辑</a-button>

            <a-button
              type="danger"
              @click="deleteConfirm(data.ID)"
              style="margin-right: 15px"
            >删除</a-button>
          </div>
        </template>
      </a-table>
      <!--/表单部分-->
    </a-card>
  </div>
</template>
    
<script>
const columns = [
  {
    title: "ID",
    dataIndex: "ID", //接收来自json的数据，和json字段一致
    sorter: true,
    width: "5%",
    align: "center",
    key: "id",

    scopedSlots: { customRender: "id" }, //是Vue子组件内具名插槽<slot name="xxx">写法的封装
  },
  {
    title: "所属分类",
    /*关于json数据的对象内的嵌套对象，不同版本的 AntDV： 
      v1 文档是这么写的：
        列数据在数据项中对应的 key，支持 a.b.c 的嵌套写法。
      v3 文档中进行了重大升级：
      “...此外，比较重大的改动为 dataIndex 从支持路径嵌套如 user.age 
          改成了数组路径如 ['user', 'age']。
          以解决过去属性名带 . 需要额外的数据转化问题。”
    */
    dataIndex: "Category.name",
    width: "10%",
    align: "center",
    key: "name",
  },
  {
    title: "文章标题",
    dataIndex: "title",
    align: "center",
    width: "20%",
    key: "title",
  },
  {
    title: "描述",
    dataIndex: "desc",
    width: "20%",
    align: "center",
    key: "desc",
  },
  {
    title: "缩略图",
    dataIndex: "image",
    width: "20%",
    align: "center",
    key: "image",
    scopedSlots: { customRender: "img" },
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
      //文章列表的表单数据
      articleList: [],
      //分类列表的表单数据
      cateList: [],

      //规定列格式
      columns,

      //每页的页码和当前页数,此处数据将传递给getArtList()的 axios 请求
      queryParam: {
        title: "",
        pagesize: 5,
        pagenum: 1,
      },

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
      //添加&编辑文章弹框文本宽度和控件宽度
      labelCol: { span: 6 },
      wrapperCol: { span: 14, offset: 1 },
      //添加&编辑文章弹框表单验证
      rules: {
        //[!]看看一会的弹框需要什么字段再修改
        title: [
          {
            required: true,
            message: "请输入文章名",
            trigger: "blur",
          },
          {
            min: 4,
            max: 12,
            message: "长度保持在4-12个字符",
            trigger: "blur",
          },
        ],
      },
    };
  },

  created() {
    //在页面创建前就获取相应数据
    this.getArtList();
    this.getCateList();
  },

  methods: {
    //查询全部文章
    async getArtList() {
      //将所接收数据的data部分赋给res.
      const { data: res } = await this.$axios.get("articles", {
        params: {
          title: this.queryParam.title, //为""时说明是无搜索条件，不为空时说明带搜索条件
          pagenum: this.queryParam.pagenum,
          pagesize: this.queryParam.pagesize,
        },
      });

      if (res.status != 200) return this.$message.error(res.message);

      //确认无误，保存返回结果
      this.articleList = res.data;
      this.pagination.total = res.total;
    },

    //查询全部分类
    async getCateList() {
      //将所接收数据的data部分赋给res.
      const { data: res } = await this.$axios.get("categories");

      if (res.status != 200) return this.$message.error(res.message);

      //确认无误，保存返回结果
      this.cateList = res.data;
      this.pagination.total = res.total;
    },
    //==========================================

    //分类下拉菜单选中后的回调，将调用getCateArt（）
    handleCateChange(value) {
      /*回调函数提供的参数 value 实际上来自 <a-select-option>的 :value 绑定值
        如果没绑定 :value ,传入的将是 :key 所绑定的值。本项目中即 item.id
      */
      //选择分类时，可以按分类查找；退出选择时，可以查找全部
      value != undefined ? this.getCateArt(value) : this.getArtList();
    },
    //按分类查找 具体实现
    async getCateArt(id) {
      const { data: res } = await this.$axios.get(`article/list/${id}`);

      if (res.status != 200) return this.$message.error(res.message);

      this.articleList = res.data;
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
      //用于保存文章操作后的新数据，并传递到 this.pagination 中，以完成页面的更新。
      let pager = { ...this.pagination };

      //1.【文章界面】拿到更新的值（由回调函数所给的参数 pagination 提供）
      pager.current = pagination.current;
      pager.pageSize = pagination.pageSize;

      //2.【后端数据请求函数】拿到更新的值（由回调函数所给的参数 pagination 提供）
      this.queryParam.pagenum = pagination.current;
      this.queryParam.pagesize = pagination.pageSize;

      //3.如果更新后的每页条目！==之前的每页条目，【文章界面】和【后端数据】都回到第一页
      if (pagination.pageSize !== this.pagination.pageSize) {
        pager.current = 1;
        this.queryParam.pagenum = 1;
      }

      //4.更新【文章界面】
      this.pagination = pager;

      //5.请求【后端数据】
      this.getArtList();
    },
    //==========================================
    //删除当前文章
    async deleteArt(id) {
      // 子组件<a-table>所提供给父组件的值 data 是当前行数据，即getArtList()的结果数组中的当前一条，字段包括
      /*  CreatedAt
          DeletedAt
          ID
          UpdatedAt
          password
          role
          title
      */
      // console.log("<a-table>向父组件的传输值:", data);
      const res = await this.$axios.delete(`article/${id}`);
      if (res.status !== 200) {
        return this.$message.error(res.message);
      }
      this.$message.success("删除成功");

      //5.删除后刷新列表
      this.getArtList();
    },
    //删除确认,调用deleteArt()
    deleteConfirm(id) {
      this.$confirm({
        title: "确定删除该文章吗?",
        content: "删除文章后将不可恢复",
        okText: "确认",
        okType: "primary",
        cancelText: "取消",
        onOk: () => {
          this.deleteArt(id); //使用箭头函数，this 指向当前Vue实例，而非调用该 this 的函数
        },
        onCancel: () => {
          this.$message.info("已取消");
        },
      });
    },
    //==========================================
    //新建文章,仅仅是跳转页面
    handleAddArt() {
      this.$router.push("/addart").catch((err) => {
        console.log(err);
      });
    },
    //==========================================
    //编辑文章按钮
    editArt(data) {
      // console.log(data); //data包含了当前行所有数据

      //1.将 data 导入v-model数据模型
      this.articleEdit.id = data.ID;
      this.articleEdit.title = data.title;
      this.articleEdit.password = data.password;
      this.articleEdit.role = data.role;

      //2.打开弹框
      this.editArtVisible = true;

      //3.此时文章编辑，修改数据。响应式地更新到模型里
    },
    //编辑文章弹窗 确认按钮
    handleEditArtOK() {
      //在 validate 中完成全部操作
      this.$refs.editArtRef.validate(async (valid) => {
        if (!valid) return this.$message.error("未通过验证，请重新填写"); //回调函数的实现总要写 return

        //axios 调用API，只能修改文章名和身份，修改密码是独立的功能点
        const { data: res } = await this.$axios.put(
          `article/${this.articleEdit.id}`,
          {
            title: this.articleEdit.title,
            role: this.articleEdit.role,
          }
        );

        if (res.status !== 200) return this.$message.error(res.message);
        this.editArtVisible = false; //关闭弹窗
        // this.$refs.editArtRef.resetFields(); //无需清空数据。反正每次打开编辑弹窗都会重新加载本行数据
        this.$message.success("编辑文章成功"); //成功提示
        this.getArtList(); //数据更新后要及时刷新
      });
    },
    //[!]编辑文章弹窗 取消按钮
    handleEditArtCancel() {
      this.editArtVisible = false;
      // this.$refs.editArtRef.resetFields();//无需清空数据，反正每次点击编辑都会重新加载数据
      this.$message.info("已取消编辑");
    },
    //==========================================
  },
};
</script>
    
<style scoped>
.actionSlot {
  display: flex;
  justify-content: center;
}

.artImg {
  display: flex;
  justify-content: center;
}
.artImg img {
  height: 55px;
  width: 55px;
}
</style>
