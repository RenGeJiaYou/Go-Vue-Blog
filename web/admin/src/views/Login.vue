<template>
  <div class="container">
    <div class="loginBox">
      <a-form-model
        ref="loginFormRef"
        class="loginForm"
        :model="formData"
        :rules="rules"
      >
        <a-form-model-item
          label="用户名"
          prop="username"
        >
          <!--表单项 item 的属性 prop :本项的唯一标识,通过item的prop 确认formData和rules 的对应字段-->
          <a-input
            placeholder="请输入用户名"
            v-model="formData.username"
          >
            <a-icon
              slot="prefix"
              type="user"
              style="color: rgba(0,0,0,.45)"
            />
          </a-input>
        </a-form-model-item>

        <a-form-model-item
          label="密码"
          prop="password"
        >
          <a-input
            placeholder="请输入密码"
            type="password"
            v-model="formData.password"
          >
            <a-icon
              slot="prefix"
              type="lock"
              style="color: rgba(0,0,0,.45)"
            />
          </a-input>
        </a-form-model-item>

        <a-form-model-item class="loginBtn">
          <a-button
            type="primary"
            style="margin-right:20px"
            @click="submitForm('loginFormRef')"
          >登录</a-button>
          <a-button
            type="info"
            @click="resetForm('loginFormRef')"
          >取消</a-button>
        </a-form-model-item>
      </a-form-model>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      //表单的绑定模型
      formData: {
        username: "",
        password: "",
      },

      //表单验证
      rules: {
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
        password: [
          {
            required: true,
            message: "请输入密码",
            trigger: "blur",
          },
          {
            min: 6,
            max: 20,
            message: "长度保持在6-20个字符",
            trigger: "blur",
          },
        ],
      },
    };
  },
  methods: {
    //传入的 formName 是表单的引用名
    //提交按钮
    submitForm(formName) {
      //rules 的校验是不够的，它只是在前端提示用户。必须有提交前的验证。否则将把错误的数据提交到后端
      this.$refs[formName].validate(async (valid) => {
        if (valid) {
          try {
            //单独提取返回的对象中的data字段，赋给 res
            const { data: res } = await this.$axios.post(
              "/login",
              this.formData
            );
            if (res.status != 200) {
              return this.$message.error(res.message);
            } else {
              //将token 放到本地存储.session 在关闭浏览器后将清空token
              // window.sessionStorage.setItem("token", res.token);
              window.localStorage.setItem("token", res.token);
              this.$router.push("/admin");
            }
          } catch (err) {
            console.log(err);
          }
        } else {
          this.$message.error("提交的账户信息有误！"); //报错message
          this.$refs[formName].resetFields(); //重置表单
        }
      });
    },
    //重置按钮
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
  },
};
</script>


<style scoped>
.container {
  height: 100%;
  background: #282c34;
  /*flex 布局常常产生不可描述的问题，改用绝对定位 
  display: flex;
  justify-content: center;
  align-items: center; */
}
.loginBox {
  width: 450px;
  height: 300px;
  background: #ffffff;
  /* 绝对定位,零点为文档左上角 */
  position: absolute;
  top: 50%;
  left: 70%;
  transform: translate(-50%, -50%);

  border-radius: 10px;
}
.loginForm {
  width: 100%;
  position: absolute;
  bottom: 10%;
  padding: 0 20px;
  /* 盒子的实际宽度 = border-left + padding-left + width + border-right + padding-right  */
  /* 盒子的实际高度 = border-top + padding-top + height + border-bottom + padding-bottom  */
  box-sizing: border-box; /*强制宽度 = width  高度 = height */
}
.loginBtn {
  display: flex;
  justify-content: flex-end;
}
</style>