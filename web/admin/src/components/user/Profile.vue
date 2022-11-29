<template>
  <div>
    <a-card>
      <h3>个人简介</h3>
      <a-form-model
        :label-col="labelCol"
        :wrapper-col="wrapperCol"
      >
        <a-form-model-item label="作者名称">
          <a-input
            style="width:300px"
            v-model="profileInfo.name"
          ></a-input>
        </a-form-model-item>

        <a-form-model-item label="状态">
          <a-input
            style="width:300px"
            v-model="profileInfo.state"
          ></a-input>
        </a-form-model-item>

        <a-form-model-item label="描述">
          <a-input
            style="width:300px"
            v-model="profileInfo.desc"
          ></a-input>
        </a-form-model-item>

        <a-form-model-item label="QQ">
          <a-input
            style="width:300px"
            v-model="profileInfo.qqchat"
          ></a-input>
        </a-form-model-item>

        <a-form-model-item label="微信">
          <a-input
            style="width:300px"
            v-model="profileInfo.wechat"
          ></a-input>
        </a-form-model-item>

        <a-form-model-item label="Weibo">
          <a-input
            style="width:300px"
            v-model="profileInfo.weibo"
          ></a-input>
        </a-form-model-item>

        <a-form-model-item label="邮箱">
          <a-input
            style="width:300px"
            v-model="profileInfo.email"
          ></a-input>
        </a-form-model-item>

        <a-form-model-item label="头像">
          <a-upload
            listType="picture"
            name="file"
            :action="upUrl"
            :headers="headers"
            @change="avatarChange"
          >
            <a-button>
              <a-icon type="upload" />点击上传
            </a-button>
            <br />
            <template v-if="profileInfo.avatar">
              <img
                :src="profileInfo.avatar"
                style="width:120px;height:100px"
              />
            </template>
          </a-upload>
        </a-form-model-item>

        <a-form-model-item label="背景图">
          <a-upload
            listType="picture"
            name="file"
            :action="upUrl"
            :headers="headers"
            @change="imgChange"
          >
            <a-button>
              <a-icon type="upload" />点击上传背景图
            </a-button>
            <br />
            <!--显示原头像-->
            <template v-if="profileInfo.img">
              <img :src="profileInfo.img" />
            </template>
          </a-upload>
        </a-form-model-item>

        <a-form-model-item>
          <a-button
            type="danger"
            style="margin-left:50%"
            @click="updateProfile"
          >更新</a-button>
        </a-form-model-item>
      </a-form-model>
    </a-card>
  </div>
</template>

<script>
import Url from "@/plugin/axios";
export default {
  name: "",
  data() {
    return {
      //数据模型
      profileInfo: {
        id: 1,
        name: "",
        img: "",
        avatar: "",
        state: "",
        desc: "",
        qqchat: "",
        wechat: "",
        weibo: "",
        email: "",
      },
      wrapperCol: { span: 16 },
      labelCol: { span: 4 },
      upUrl: Url + "upload",
      headers: {},
    };
  },
  created() {
    this.getProfileInfo();
    this.headers = {
      Authorization: `Bearer ${window.localStorage.getItem("token")}`,
    };
  },
  methods: {
    // 获取个人简介
    async getProfileInfo() {
      const { data: res } = await this.$axios.get(
        `profile/${this.profileInfo.id}`
      );
      if (res.status !== 200) {
        if (res.status === 1004 || 1005 || 1006 || 1007) {
          window.sessionStorage.clear();
          this.$router.push("/login");
        }
        this.$message.error(res.message);
      }
      // console.log(res);
      this.profileInfo = res.data;
    },
    /*
「上传背景图」和「上传头像」函数将图片文件传到七牛云，七牛云返回一个地址（string）；
「提交」将表单数据（avatar和img的值是图片的云存储空间地址）
*/
    // 上传头像
    avatarChange(info) {
      if (info.file.status !== "uploading") {
      }
      if (info.file.status === "done") {
        this.$message.success(`图片上传成功`);
        const imgUrl = info.file.response.url;
        this.profileInfo.avatar = imgUrl;
      } else if (info.file.status === "error") {
        this.$message.error(`图片上传失败`);
      }
    },
    // 上传背景图
    imgChange(info) {
      if (info.file.status !== "uploading") {
      }
      if (info.file.status === "done") {
        this.$message.success(`图片上传成功`);
        const imgUrl = info.file.response.url;
        this.profileInfo.avatar = imgUrl;
      } else if (info.file.status === "error") {
        this.$message.error(`图片上传失败`);
      }
    },
    //更新
    async updateProfile() {
      const { data: res } = await this.$axios.put(
        `profile/${this.profileInfo.id}`,
        this.profileInfo
      );
      if (res.status !== 200) return this.$message.error(res.message);
      this.$message.success(`个人信息更新成功`);
      this.$router.push("/index");
    },
  },
};
</script>

<style>
</style>

