<script setup lang="ts">
import {onMounted, reactive, ref} from "vue";
import axios from "axios";
// addStoreVisible 用来控制 dialog显示
let addStoreVisible = ref(false)

// 获取用户列表
let userList = reactive([])
const getUserAdd = () => {
  axios.get("http://localhost:8080/api/admin/getUserList")
      .then((response) => {
        userList = response.data.userList;
        console.log(userList);
      })
      .catch((err) => {
        console.log(err)
      })
}

// 获取权限列表
let accessList = reactive([])
const getAccessList = () => {
  axios.get("http://localhost:8080/api/admin/accessList")
      .then((response) => {
        accessList = response.data.accessList;
        console.log(accessList);
      })
      .catch((err) => {
        console.log(err)
      })
}

// 钩子获取数据权限
onMounted(() => {
  getUserAdd()
  getAccessList()
})

</script>

<template>
  <el-dialog v-model="addStoreVisible" title="添加商店">
    <el-form>
      <el-form-item>
        <el-select @focus="">

        </el-select>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>

<style scoped>

</style>