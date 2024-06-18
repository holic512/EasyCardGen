<script setup lang="ts">

import {reactive, ref} from "vue";
import {FormInstance, FormRules} from "element-plus";
import axios from "axios";

//addUserVisible 用于 控制添加表单 显示
const addUserVisible = ref(false)

// userForm 用于存储表单信息
const userForm = ref({
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
const validateName = (rule: any, value: any, callback: any) => {
  if (value === '') {
    callback(new Error('用户名不能为空'))
  } else {
    if (value.length < 4 || value.length > 12) {
      callback(new Error('用户名要至少4位且不大于12位'))
    } else callback()
  }
}

// validatePass,validatePass2 用于 验证密码 规则
// @ts-ignore
const validatePass = (rule: any, value: any, callback: any) => {
  if (value === '') {
    callback(new Error('密码不能为空'))
  } else {
    if (value.length < 8) {
      callback(new Error('密码不能小于8位'))
      return;
    } else if (value.length > 16) {
      callback(new Error('密码不能大于16位'))
      return;
    }

    if (userForm.value.confirmPassword !== '') {
      if (!ruleFormRef.value) return
      ruleFormRef.value.validateField('confirmPassword')
    }
    callback()
  }
}
// @ts-ignore
const validatePass2 = (rule: any, value: any, callback: any) => {
  if (value === '') {
    callback(new Error('密码不能为空'))
  } else if (value !== userForm.value.password) {
    callback(new Error("两次输入的密码不一致"))
  } else {
    callback()
  }
}

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


// 用于实现 邮箱补全
const restaurants = [
  {value: '@email.com'},
  {value: '@qq.com'},
  {value: '@163.com'},
  {value: '@126.com'},
  {value: '@gogo.com'}
];

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


const submitForm = (formEl: FormInstance | undefined) => {
  if (!formEl) {
    console.log('1');
    return;
  }

  // 检测表单规范
  formEl.validate((valid) => {
    if (valid) {
      console.log('2');
      // 执行axios
      axios.post('http://localhost:8080/api/admin/addUser', {
        Username: userForm.value.username,
        Password: userForm.value.password,
        Email: userForm.value.email,
        Phone: userForm.value.phone,
        userType: userForm.value.userType,
        State: userForm.value.state,
      })
          .then(response => {
            console.log(response.data);
          })
          .catch(error => {
            console.error(error);
            // 处理登录失败的逻辑
          });
    } else {
      console.log('3');
      // 可以考虑在这里显示错误信息给用户
      // 例如：使用 Element Plus 的 Message 组件
      // ElementPlus.ElMessage.error('表单验证失败！');
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
        :model="userForm"
        :label-position="'right'"
        label-width="auto"
        :rules="rules"
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