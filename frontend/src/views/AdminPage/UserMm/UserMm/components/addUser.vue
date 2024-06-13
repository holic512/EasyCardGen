<script setup lang="ts">
//addUserVisible 用于 控制添加表单 显示
import {ref} from "vue";

const addUserVisible = ref(false)

const userForm = ref({
  username: String,
  password: String,
  email: String,
  phone: String,
  userType: 'user',
  state: 'active'
})

</script>

<template>
  <!--添加用户 对话框-->
  <el-dialog v-model="addUserVisible" title="添加新用户">
    <el-form
        :label-position="'right'"
        label-width="auto"
    >
      <el-form-item label="用户名">
        <el-input>

        </el-input>
      </el-form-item>

      <el-form-item label="密码">
        <el-input>

        </el-input>
      </el-form-item>

      <el-form-item label="电话">
        <el-input>

        </el-input>
      </el-form-item>

      <el-form-item label="邮箱">
        <el-input>

        </el-input>
      </el-form-item>

      <el-form-item label="状态">
        <el-radio-group v-model="userForm.state">
          <el-radio value="active">正常</el-radio>
          <el-radio value="disabled">停用</el-radio>
          <el-radio value="banned">封禁</el-radio>
        </el-radio-group>
      </el-form-item>

      <el-form-item label="权限">
        <el-radio-group v-model="userForm.userType">
          <el-radio value="user">用户</el-radio>
          <el-radio value="merchant">商户</el-radio>
          <el-radio value="admin">管理员</el-radio>
        </el-radio-group>
      </el-form-item>

      <el-form-item>
        <el-button type="primary">提交</el-button>
        <el-button>重置</el-button>
      </el-form-item>

    </el-form>
  </el-dialog>
</template>

<style scoped>

</style>