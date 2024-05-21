<template>
  <div class="Con" v-show="ifshow">
    <div style="margin-right: 100px;">
      <img src="./public/登录.png" alt="" class="Img1">
    </div>
    <div>
      <form class="form_container" @submit.prevent="submitForm">
        <div class="title_container">
          <p class="title">商户登录</p>
          <span class="subtitle">开始使用我们的应用程序，只需创建一个帐户并享受体验.</span>
        </div>
        <br>
        <div class="input_container">
          <label class="input_label" for="email_field">用户名</label>
          <el-icon :size="24" class="icon">
            <User/>
          </el-icon>
          <input placeholder="请输入用户名" name="input-name" type="text" class="input_field"
                 id="email_field" v-model="this.username">
        </div>
        <div class="input_container">
          <label class="input_label" for="password_field">密码</label>
          <el-icon :size="24" class="icon">
            <Key/>
          </el-icon>
          <input placeholder="请输入密码" name="input-name" type="password" class="input_field"
                 id="password_field" v-model="this.password">
        </div>
        <div class="link_container">
          <el-link class="input_label" href="#" type="info">忘记密码?</el-link>
        </div>

        <button title="Sign In" type="submit" class="sign-in_btn">
          <span>登录</span>
        </button>

        <div class="separator">
          <hr class="line">
          <span>Or</span>
          <hr class="line">
        </div>

        <button title="Sign In" type="submit" class="sign-in_ggl">
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" class="w-5 h-5 fill-current" height="22px">
            <path
                d="M16.318 13.714v5.484h9.078c-0.37 2.354-2.745 6.901-9.078 6.901-5.458 0-9.917-4.521-9.917-10.099s4.458-10.099 9.917-10.099c3.109 0 5.193 1.318 6.38 2.464l4.339-4.182c-2.786-2.599-6.396-4.182-10.719-4.182-8.844 0-16 7.151-16 16s7.156 16 16 16c9.234 0 15.365-6.49 15.365-15.635 0-1.052-0.115-1.854-0.255-2.651z">
            </path>
          </svg>
          <span>使用Google登录</span>
        </button>

        <button title="Sign In" type="submit" class="sign-in_apl">
          <svg style="fill: #fff;" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" height="22px"
               width="24px">
            <path
                d="M16 0.396c-8.839 0-16 7.167-16 16 0 7.073 4.584 13.068 10.937 15.183 0.803 0.151 1.093-0.344 1.093-0.772 0-0.38-0.009-1.385-0.015-2.719-4.453 0.964-5.391-2.151-5.391-2.151-0.729-1.844-1.781-2.339-1.781-2.339-1.448-0.989 0.115-0.968 0.115-0.968 1.604 0.109 2.448 1.645 2.448 1.645 1.427 2.448 3.744 1.74 4.661 1.328 0.14-1.031 0.557-1.74 1.011-2.135-3.552-0.401-7.287-1.776-7.287-7.907 0-1.751 0.62-3.177 1.645-4.297-0.177-0.401-0.719-2.031 0.141-4.235 0 0 1.339-0.427 4.4 1.641 1.281-0.355 2.641-0.532 4-0.541 1.36 0.009 2.719 0.187 4 0.541 3.043-2.068 4.381-1.641 4.381-1.641 0.859 2.204 0.317 3.833 0.161 4.235 1.015 1.12 1.635 2.547 1.635 4.297 0 6.145-3.74 7.5-7.296 7.891 0.556 0.479 1.077 1.464 1.077 2.959 0 2.14-0.020 3.864-0.020 4.385 0 0.416 0.28 0.916 1.104 0.755 6.4-2.093 10.979-8.093 10.979-15.156 0-8.833-7.161-16-16-16z">
            </path>
          </svg>
          <span>使用Github登录</span>
        </button>
        <div class="container">
          <el-text type="info" class="text-container">
            没有账号？
            <router-link :to="{ path: '/sign/register' }">
              <el-link href="#" type="info" style="font-weight: bold; margin-bottom: 2px;">立即注册</el-link>
            </router-link>
          </el-text>
        </div>

        <p class="note">使用条款及细则</p>
      </form>
    </div>
  </div>
</template>
<script>
import {Key, User} from '@element-plus/icons-vue'
import {ElNotification} from "element-plus";
import axios from "axios";

export default {
  data() {
    return {
      ifshow: false,
      username: '',
      password: ''
    };
  },
  components: {
    Key,
    User
  },

  mounted() {

    // 加载样式
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
          this.username === '' ||
          this.password === ''
      ) {
        ElNotification({
          title: '注册资料不完整',
          message: '用于注册的资料有空缺,请将数据填写完整',
          type: 'warning',
        });
        return
      }
      axios
          .post("http://localhost:8080/user/login", {
            username: this.username,
            password: this.password,
          })
          .then(response => {
            const statusData = response.data;
            if (statusData === "登录成功") {
              ElNotification({
                title: '登录成功',
                message: '3秒后自动跳转到用户中心',
                type: 'success',
              });
            } else if (statusData === "登录失败") {
              ElNotification({
                title: '登录失败',
                message: '用户名不存在或者密码错误',
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

}

.text-container {
  display: flex;
  align-items: center;
}

/* 其他样式保持不变 */
.link_container {
  margin-left: auto;
  /* 将元素推到容器的最右侧 */
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
  box-shadow: 0px 106px 42px rgba(0, 0, 0, 0.01),
  0px 59px 36px rgba(0, 0, 0, 0.01), 0px 26px 26px rgba(0, 0, 0, 0.01),
  0px 7px 15px rgba(0, 0, 0, 0.01), 0px 0px 0px rgba(0, 0, 0, 0.01);
  border-radius: 11px;
  font-family: "Inter", sans-serif;

}

.logo_container {
  box-sizing: border-box;
  width: 80px;
  height: 80px;
  background: linear-gradient(180deg, rgba(248, 248, 248, 0) 50%, #F8F8F888 100%);
  border: 1px solid #F7F7F8;
  filter: drop-shadow(0px 0.5px 0.5px #EFEFEF) drop-shadow(0px 1px 0.5px rgba(239, 239, 239, 0.5));
  border-radius: 11px;
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

.subtitle {
  font-size: 0.725rem;
  max-width: 80%;
  text-align: center;
  line-height: 1.1rem;
  color: #8B8E98
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
  filter: drop-shadow(0px 1px 0px #efefef) drop-shadow(0px 1px 0.5px rgba(239, 239, 239, 0.5));
  transition: all 0.3s cubic-bezier(0.15, 0.83, 0.66, 1);
}

.input_field:focus {
  border: 1px solid transparent;
  box-shadow: 0px 0px 0px 2px #242424;
  background-color: transparent;
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
}

.sign-in_ggl {
  width: 100%;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  background: #ffffff;
  border-radius: 7px;
  outline: none;
  color: #242424;
  border: 1px solid #e5e5e5;
  filter: drop-shadow(0px 1px 0px #efefef) drop-shadow(0px 1px 0.5px rgba(239, 239, 239, 0.5));
  cursor: pointer;
}

.sign-in_apl {
  width: 100%;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  background: #212121;
  border-radius: 7px;
  outline: none;
  color: #ffffff;
  border: 1px solid #e5e5e5;
  filter: drop-shadow(0px 1px 0px #efefef) drop-shadow(0px 1px 0.5px rgba(239, 239, 239, 0.5));
  cursor: pointer;
}

.separator {
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 30px;
  color: #8B8E98;
}

.separator .line {
  display: block;
  width: 100%;
  height: 1px;
  border: 0;
  background-color: #e8e8e8;
}

.note {
  font-size: 0.75rem;
  color: #8B8E98;
  text-decoration: underline;
}
</style>
