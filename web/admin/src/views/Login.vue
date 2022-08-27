<template>
  <div class="container">
    <div class="loginBox">
      <a-form-model
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
            <a-tooltip
              slot="suffix"
              title="用户名包括英文，数字和下划线"
            >
              <a-icon
                type="info-circle"
                style="color: rgba(0,0,0,.45)"
              />
            </a-tooltip>
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
          >登录</a-button>
          <a-button type="info">取消</a-button>
        </a-form-model-item>
      </a-form-model>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      //表单的绑定模型
      formData: {
        username: "",
        password: "",
        role: "",
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