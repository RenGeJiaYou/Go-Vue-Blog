<template>
  <div>
    <a-card>
      <h3>{{id?"编辑文章":"新增文章"}}</h3>

      <a-form-model
        :model="artInfo"
        :rules="rules"
        ref="artInfoRef"
        hideRequiredMark="true"
      >
        <a-form-model-item
          label="文章标题"
          prop="title"
        >
          <a-input
            style="width: 300px;"
            v-model="artInfo.title"
          ></a-input>
        </a-form-model-item>

        <a-form-model-item
          label="文章分类"
          prop="category"
        >
          <a-select
            placeholder="选择分类"
            v-model="artInfo.cid"
          >
            <a-select-option
              v-for="item in cateList"
              :key="item.id"
            >{{item.name}}</a-select-option>
          </a-select>
        </a-form-model-item>

        <a-form-model-item
          label="文章描述"
          prop="desc"
        >
          <a-textarea v-model="artInfo.desc"></a-textarea>
        </a-form-model-item>

        <a-form-model-item
          label="文章缩略图"
          prop="img"
        >
          <a-upload
            listType="picture"
            name="file"
            :action="uploadUrl"
            :headers="headers"
            :defaultFileList="fileList"
            @change="handleUpload"
          >
            <a-button>
              <a-icon type="upload" />点击上传
            </a-button>
            <br />

            <template v-if="this.id">
              <img
                :src="artInfo.image"
                style="height:100px"
              />
            </template>
          </a-upload>
        </a-form-model-item>

        <a-form-model-item
          label="文章内容"
          prop="content"
        >
          <a-input
            v-model="artInfo.content"
            placeholder="后面 input 标签会改成富文本编辑器"
          ></a-input>
        </a-form-model-item>

        <a-form-model-item>
          <!--[!]flex 居右-->
          <a-button
            type="primary"
            style="margin-right: 15px"
            @click="handleSubmit(artInfo.id)"
          >{{artInfo.id?"更新":"提交"}}</a-button>
          <a-button @click="handleCancel">取消</a-button>
        </a-form-model-item>
      </a-form-model>
    </a-card>
  </div>
</template>
  
<script>
import { Url } from "../../plugin/axios";

export default {
  name: "AdminIndex",

  props: ["id"],

  data() {
    return {
      artInfo: {
        id: 0,
        title: "",
        cid: undefined, //若不希望默认值或初始化时显示一个 0 / "",使用 undefined 消除影响
        desc: "",
        content: "",
        image: "",
      },

      cateList: [], //分类列表

      uploadUrl: Url + "upload", //上传 url

      headers: {}, //请求头

      fileList: [], //上传文件列表

      rules: {
        //表单验证
        title: [
          { required: true, message: "文章标题不得为空", trigger: "blur" },
          {
            min: 3,
            max: 100,
            message: "文章标题长度为3~100个字符",
            trigger: "change",
          },
        ],
        desc: [],
        img: [],
        content: [
          { required: true, message: "文章内容不得为空", trigger: "blur" },
        ],
      },
    };
  },

  mounted() {
    this.getCateList();
    this.headers = {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
    };
    if (this.id) {
      //当this.id 不为空时，说明是从「编辑按钮」跳转而来
      this.getArtInfo(this.id);
    }
  },

  methods: {
    //查询全部分类。编辑文章时，必须已加载全部标签。
    async getCateList() {
      //将所接收数据的data部分赋给res.
      const { data: res } = await this.$axios.get("categories");

      if (res.status != 200) return this.$message.error(res.message);

      //确认无误，保存返回结果
      this.cateList = res.data;
    },

    //将待编辑文章传入对应模型
    async getArtInfo(id) {
      const { data: res } = await this.$axios.get(`article/info/${id}`);
      if (res.status != 200) return this.$message.error(res.message);

      (this.artInfo = res.data),
        (this.artInfo.id = res.data.ID),
        console.log(this.artInfo);
    },

    //v-model 已经双向绑定artInfo.cid,不需要写 a-select 组件的 @change

    //上传图片操作
    handleUpload(info) {
      console.table(info);
      if (info.file.status !== "uploading") {
        console.log(info.file, info.fileList);
      }
      if (info.file.status === "done") {
        this.$message.success(`${info.file.name} 文件上传成功`);

        this.artInfo.image = info.file.response.url; //保存图片的外链地址
      } else if (info.file.status === "error") {
        this.$message.error(`${info.file.name} 文件上传失败`);
        this.$message.error(this.uploadUrl);
      }
    },

    //发布文章操作
    handleSubmit(artId) {
      this.$refs.artInfoRef.validate(async (valid) => {
        if (valid) {
          if (artId) {
            //编辑文章
            const { data: res } = await this.$axios.put(
              `article/${artId}`,
              this.artInfo
            );
            if (res.status != 200)
              return this.$message.error("编辑文章失败： " + res.message);

            this.$message.success("编辑文章成功");
            this.$router.push("artlist");
          } else {
            //添加文章
            const { data: res } = await this.$axios.post(
              "article/add",
              this.artInfo
            );
            if (res.status != 200)
              return this.$message.error("新增文章失败： " + res.message);

            this.$message.success("新增文章成功");
            this.$router.push("artlist");
          }
        }
      });
    },

    //取消发布按钮
    handleCancel() {
      //清空不再需要的数据
      this.$refs.artInfoRef.resetFileds();

      //跳转回文章列表
      this.$router.push("artlist");
    },
  },
};
</script>
  
<style scoped>
<style>
