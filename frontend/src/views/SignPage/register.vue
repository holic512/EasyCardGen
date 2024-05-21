<template>
  <div class="Con" v-show="ifshow">
    <div style="margin-right: 100px;">
      <img src="./public/注册.png" alt="" class="Img1">
    </div>
    <div>
      <form class="form_container" @submit.prevent="submitForm">
        <div class="title_container">
          <p class="title">商户注册</p>
        </div>
        <div class="input_container">
          <label class="input_label">用户名</label>
          <el-icon :size="24" class="icon">
            <User/>
          </el-icon>
          <input placeholder="请输入用户名" type="text" class="input_field" v-model="formData.username">
        </div>
        <div class="input_container">
          <label class="input_label">密码</label>
          <el-icon :size="24" class="icon">
            <Key/>
          </el-icon>
          <input placeholder="请输入密码" type="password" class="input_field" v-model="formData.password">
        </div>
        <div class="input_container">
          <label class="input_label">确认密码</label>
          <el-icon :size="24" class="icon">
            <Key/>
          </el-icon>
          <input placeholder="请输入确认密码" type="password" class="input_field" v-model="formData.confirmPassword">
        </div>
        <div class="input_container">
          <label class="input_label">手机号</label>
          <el-icon :size="24" class="icon">
            <Phone/>
          </el-icon>
          <input placeholder="请输入手机号" type="text" class="input_field" v-model="formData.phoneNumber">
        </div>

        <div class="input_container">
          <label class="input_label">邮箱</label>
          <el-icon :size="24" class="icon">
            <Iphone/>
          </el-icon>
          <input placeholder="请输入邮箱" type="text" class="input_field" v-model="formData.email">
        </div>

        <div class="input_container">
          <label class="input_label">邮箱验证码</label>
          <div class="input_with_button"> <!-- 新添加的 div 包含输入框和按钮 -->
            <el-icon :size="24" class="icon">
              <Message/>
            </el-icon>
            <input style="width: 80%;" placeholder="请输入邮箱验证码" type="text" class="input_field"
                   v-model="formData.emailVerificationCode">
            <el-button class="el-button" size="large">发送验证码</el-button>
          </div>
        </div>
        <div class="input_container">
          <label class="input_label">邀请码</label>
          <el-icon :size="24" class="icon">
            <Position/>
          </el-icon>
          <input placeholder="邀请码(选填)" type="text" class="input_field" v-model="formData.invitationCode">
        </div>
        <div class="container">
          <el-checkbox v-model="formData.agreement" size="large" style="margin-right: 5px;"/>
          <el-text class="text-container">
            阅读并同意
            <el-link type="primary">服务条款和条件</el-link>
          </el-text>

        </div>
        <button title="Sign In" type="submit" class="sign-in_btn">
          <span>注册</span>
        </button>

        <div class="container1">
          <el-text type="info" class="text-container">
            已有账号
            <router-link :to="{ path: '/sign' }">
              <el-link href="#" type="info" style="font-weight: bold; margin-bottom: 2px;">立即登录</el-link>
            </router-link>
          </el-text>
        </div>


      </form>
    </div>
  </div>
</template>
<script>
import {Iphone, Key, Message, Phone, Position, User} from '@element-plus/icons-vue'

import {ElNotification} from 'element-plus'

import axios from 'axios'

export default {
  data() {
    return {
      ifshow: false,

      formData: {
        username: '',
        password: '',
        confirmPassword: '',
        phoneNumber: '',
        email: '',
        emailVerificationCode: '',
        invitationCode: '',

        agreement: false
      }
    };
  },
  components: {
    Key,
    User,
    Iphone,
    Message,
    Phone,
    Position
  },
  mounted() {
    const openFullScreen2 = () => {
      const loading = this.$loading({
        lock: true,
        text: 'Loading',
        background: 'rgb(255, 255, 255)',
      });

      setTimeout(() => {
        loading.close();
        this.ifshow = true
      }, 200);

    };
    openFullScreen2(); // 如果需要在组件挂载后立即执行该函数，可以调用它
  },
  methods: {
    submitForm() {
      if (
          this.formData.username === '' ||
          this.formData.password === '' ||
          this.formData.confirmPassword === '' ||
          this.formData.phoneNumber === '' ||
          this.formData.email === '' ||
          this.formData.emailVerificationCode === ''
      ) {
        ElNotification({
          title: '注册资料不完整',
          message: '用于注册的资料有空缺,请将数据填写完整',
          type: 'warning',
        });
        return
      }
      if (this.formData.agreement === false) {
        ElNotification({
          title: '阅读并同意服务条款和条件',
          message: '未勾选阅读并同意服务条款和条件,请保证阅读并且同意服务条款和条件后,进行勾选',
          type: 'warning',
        });
        return
      }

      axios
          .post("http://localhost:8080/user/register", {
            username: this.formData.username,
            password: this.formData.password,
            phone: this.formData.phoneNumber,
            mail: this.formData.email,
            email_verification_code: this.formData.emailVerificationCode,
            invitation_code: this.formData.invitationCode
          })
          .then(response => {
            const statusData = response.data;
            if (statusData === "注册成功") {
              ElNotification({
                title: '注册成功',
                message: '3秒后自动跳转到用户中心',
                type: 'success',
              });
            } else if (statusData === "邮箱验证码错误") {
              ElNotification({
                title: '邮箱验证码错误',
                message: '请输入正确的验证码',
                type: 'warning',
              });
            } else if (statusData === '邮箱已被注册') {
              ElNotification({
                title: '邮箱已被注册',
                message: '邮箱已被注册,如果不是您注册请联系管理员',
                type: 'warning',
              });
            } else if (statusData === "用户名已被注册") {
              ElNotification({
                title: '用户名已被注册',
                message: '用户名已被注册,请更换用户名再注册',
                type: 'warning',
              });
            } else {
              ElNotification({
                title: '错误',
                message: '出现了未知的错误，请联系网站管理员',
                type: 'warning',
              });
            }

          })
    }
  }
}

</script>
<style scoped>
.container {
  display: flex;
  align-items: center;
  margin-right: auto;
}

.container1 {
  display: flex;
  align-items: center;
}

.text-container {
  display: flex;
  align-items: center;
}


/* 保留原有的样式设置 */
.Con {
  display: flex;
  /* 使用 Flexbox 布局 */
  justify-content: center;
  /* 水平居中 */
  align-items: center;
  /* 垂直居中 */
  width: 100vw;
  /* 设置宽度为视口宽度 */
  height: 100vh;
  /* 设置高度为视口高度 */
  background-color: rgb(244, 244, 244);
}

.Img1 {
  height: 400px;
}

.form_container {
  width: 100%;
  height: fit-content;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 15px;
  padding: 50px 40px 20px 40px;
  background-color: #ffffff;
  box-shadow: 0 106px 42px rgba(0, 0, 0, 0.01),
  0 59px 36px rgba(0, 0, 0, 0.01), 0 26px 26px rgba(0, 0, 0, 0.01),
  0 7px 15px rgba(0, 0, 0, 0.01), 0 0 0 rgba(0, 0, 0, 0.01);
  border-radius: 11px;
  font-family: "Inter", sans-serif;

}

.title_container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 10px;
}

.title {
  margin: 0;
  font-size: 1.5rem;
  font-weight: 700;
  color: #212121;
}

.input_container {
  width: 100%;
  height: fit-content;
  position: relative;
  display: flex;
  flex-direction: column;
  gap: 5px;

}

.icon {
  width: 20px;
  position: absolute;
  z-index: 99;
  left: 12px;
  bottom: 9px;
}

.input_label {
  font-size: 0.75rem;
  color: #8B8E98;
  font-weight: 600;
}

.input_field {
  width: auto;
  height: 40px;
  padding: 0 0 0 40px;
  border-radius: 7px;
  outline: none;
  border: 1px solid #e5e5e5;
  filter: drop-shadow(0px 1px 0px #efefef) drop-shadow(0px 1px 1px rgba(239, 239, 239, 0.5));
  transition: all 0.3s cubic-bezier(0.15, 0.83, 0.66, 1);
}

.input_field:focus {
  border: 1px solid transparent;
  box-shadow: 0 0 0 2px #242424;
  background-color: transparent;
}

.input_with_button {
  display: flex;
  align-items: center;
  /* Align items vertically */
  position: relative;
}


.el-button {
  margin-left: auto;
  /* Push the button to the right */

}


.sign-in_btn {
  width: 100%;
  height: 40px;
  border: 0;
  background: #115DFC;
  border-radius: 7px;
  outline: none;
  color: #ffffff;
  cursor: pointer;
  margin-bottom: 10px;
}

</style>
