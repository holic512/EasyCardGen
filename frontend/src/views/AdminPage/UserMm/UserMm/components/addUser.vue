<script setup lang="ts">

import {reactive, ref} from "vue";
import {ElMessage, FormInstance, FormRules} from "element-plus";
import axios from "axios";

//addUserVisible 用于 控制添加表单 显示
let addUserVisible = ref(false)

// userForm 用于存储表单信息
let userForm = ref({
  username: '',
  password: '',
  confirmPassword: '',
  email: '',
  phone: '',
  userType: 'user',
  state: 'active'
})

// ruleFormRef 用于存储 表格 总规则
const ruleFormRef = ref<FormInstance>()

// validateName
// @ts-ignore
const validateName = (rule: any, value: string, callback: Function) => {
  if (value === '') {
    // 用户名不能为空
    callback(new Error('用户名不能为空'));
  } else if (!/^[A-Za-z0-9_]+$/.test(value)) {
    // 用户名只能包含大小写英文字母、数字以及下划线
    callback(new Error('用户名只能包含大小写英文、数字以及_'));
  } else if (value.length < 4 || value.length > 12) {
    // 用户名要至少4位且不大于12位
    callback(new Error('用户名要至少4位且不大于12位'));
  } else {
    // 验证通过
    callback();
  }
};

// 验证密码规则
const validatePass = (rule: any, value: string, callback: Function) => {
  if (value === '') {
    callback(new Error('密码不能为空'));
  } else if (!/^[A-Za-z0-9_!@#$%^&*(),.?":{}|<>]+$/.test(value)) {
    // 这里添加了更多的允许字符，你可以根据需要调整正则表达式
    callback(new Error('密码只能包含大小写英文、数字以及允许的特殊字符（如!@#$%^&*(),.?":{}|<>）'));
  } else if (value.length < 8) {
    callback(new Error('密码不能小于8位'));
  } else if (value.length > 16) {
    callback(new Error('密码不能大于16位'));
  } else {
    // 如果需要，可以在这里添加对confirmPassword的验证，但通常应该在confirmPassword的验证中做
    // 不过，如果你确实想在这里触发confirmPassword的验证，确保ruleFormRef.value是有效的
    if (userForm.value.confirmPassword !== '') {
      if (!ruleFormRef.value) return;
      ruleFormRef.value.validateField('confirmPassword');
    }
    callback();
  }
};

// 验证确认密码规则
const validatePass2 = (rule: any, value: string, callback: Function) => {
  if (value === '') {
    callback(new Error('确认密码不能为空'));
  } else if (value !== userForm.value.password) {
    callback(new Error("两次输入的密码不一致"));
  } else if (!/^[A-Za-z0-9_!@#$%^&*(),.?":{}|<>]+$/.test(value)) {
    // 同样，这里添加了与密码相同的字符规则
    callback(new Error('确认密码只能包含大小写英文、数字以及允许的特殊字符（如!@#$%^&*(),.?":{}|<>）'));
  } else {
    callback();
  }
};

//validateEmail 用于 验证邮箱 规则
// @ts-ignore
const validateEmail = (rule: any, value: any, callback: any) => {
  if (value === '') {
    callback(new Error('邮箱不能为空'))
    return;
  }

  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  const isValidEmail = emailRegex.test(value);
  if (!isValidEmail) {
    callback(new Error('请输入正确的邮箱格式'))
    return
  }
  callback()

}

//validatePhone 用于 验证电话 规则
// @ts-ignore
const validatePhone = (rule: any, value: any, callback: any) => {
  if (value === '') {
    callback(new Error('电话不能为空'))
    return;
  }

  if (value.length != 11) {
    callback(new Error('请输入正确的电话格式'))
    return
  }
  callback()
}


// 用于实现 表单规则
const rules = reactive<FormRules>({
  username: [{validator: validateName, trigger: 'blur'}],
  password: [{validator: validatePass, trigger: 'blur'}],
  confirmPassword: [{validator: validatePass2, trigger: 'blur'}],
  phone: [{validator: validatePhone, trigger: 'blur'}],
  email: [{validator: validateEmail, trigger: 'blur'}],
})


// restaurants 用于实现 邮箱补全
const restaurants = [
  {value: '@qq.com'},
  {value: '@163.com'},
  {value: '@126.com'},
  {value: '@139.com'},
  {value: '@sina.com'},
  {value: '@gmail.com'},

];

// fetchSuggestions 用于 邮箱自动填充规则
const fetchSuggestions = (queryString: string, cb: any) => {

  if (queryString === '') {
    cb([])
    return
  }

  const hasAtSymbol = /@/.test(queryString);
  if (hasAtSymbol) {
    cb([])
    return
  }

  const suggestions = restaurants.map(item => {
        return {...item, value: queryString + item.value}
      }
  )
  cb(suggestions)
}

// stateOptions 用于状态选择框
const stateOptions = [
  {
    label: '正常',
    value: 'active',
  },
  {
    label: '停用',
    value: 'disabled',
  },
  {
    label: '封禁',
    value: 'banned',
  }
]

// typeOptions 用于 身份 选择框
const typeOptions = [
  {
    label: '用户',
    value: 'user',
  },
  {
    label: '商户',
    value: 'merchant',
    disabled: true
  },
  {
    label: '管理员',
    value: 'admin',
  }
]

// submitForm 用于提交表单
const submitForm = (formEl: FormInstance | undefined) => {
  if (!formEl) {
    return;
  }

  // 函数用于重置表单数据为初始值
  const resetForm = () => {
    userForm.value = {
      username: '',
      password: '',
      confirmPassword: '',
      email: '',
      phone: '',
      userType: 'user',
      state: 'active'
    };
  };


  // 检测表单规范
  formEl.validate((valid) => {
    if (valid) {
      // 执行axios
      axios.post('http://localhost:8080/api/admin/addUser', {
        Username: userForm.value.username,
        Password: userForm.value.password,
        Email: userForm.value.email,
        Phone: userForm.value.phone,
        UserType: userForm.value.userType,
        State: userForm.value.state,
      })
          .then(response => {
            if (response.status === 200) {
              ElMessage({
                message: '添加用户成功!',
                type: 'success',
              })
              // 初始化
              resetForm()
              addUserVisible.value = false
            }
          })
          .catch(error => {
            let errorMessage;
            switch (error.response.data.error_code) {
              case 'USERNAME_DUPLICATE_ERROR':
                errorMessage = '用户名已存在';
                break;
              case 'EMAIL_DUPLICATE_ERROR':
                errorMessage = '邮箱已存在';
                break;
              case 'PHONE_DUPLICATE_ERROR':
                errorMessage = '电话已存在';
                break;
              default:
                errorMessage = '添加用户失败';
            }
            // 处理登录失败的逻辑
            ElMessage({
              message: errorMessage,
              type: 'warning',
            })
          });
    } else {
      ElMessage.error('表单验证失败！');
    }
  });
}

// resetForm 用于 重置 form
const resetForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.resetFields()
}

</script>

<template>
  <!--添加用户 对话框-->
  <el-dialog v-model="addUserVisible" title="添加新用户">
    <el-form
        ref="ruleFormRef"
        :rules="rules"
        :model="userForm"
        :label-position="'right'"
        label-width="auto"
    >
      <el-form-item
          label="用户名"
          prop="username"
      >
        <el-input
            v-model="userForm.username"
            placeholder="请输入用户名"
        />

      </el-form-item>

      <el-form-item
          label="密码"
          prop="password"
      >
        <el-input
            type="password"
            v-model="userForm.password"
            autocomplete="off"
            show-password
            placeholder="请输入密码"
        />


      </el-form-item>

      <el-form-item
          label="确认密码"
          prop="confirmPassword"
      >
        <el-input
            type="password"
            v-model="userForm.confirmPassword"
            autocomplete="off"
            show-password
            placeholder="请再次输入密码"
        />


      </el-form-item>

      <el-form-item label="邮箱" prop="email">

        <el-autocomplete
            v-model="userForm.email"
            :fetch-suggestions="fetchSuggestions"
            placeholder="请输入用户邮箱"
            style="width: 100%"
        />

      </el-form-item>
      <el-form-item label="电话" prop="phone">

        <el-input v-model="userForm.phone" placeholder="请输入用户电话"/>

      </el-form-item>


      <el-form-item label="状态" prop="state">
        <el-segmented v-model="userForm.state" :options="stateOptions" size="default"/>
      </el-form-item>

      <el-form-item label="权限" prop="userType">
        <el-segmented v-model="userForm.userType" :options="typeOptions" size="default"/>
      </el-form-item>

      <el-form-item>
        <el-button
            type="primary"
            @click="submitForm(ruleFormRef)"
        >
          提交
        </el-button>

        <el-button @click="resetForm(ruleFormRef)">重置</el-button>

      </el-form-item>

    </el-form>
  </el-dialog>
</template>

<style scoped>

</style>