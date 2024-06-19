<!--用于管理员管理-用户管理-所有用户-->
<script setup lang="ts">

//添加用户页面
import {ref, watch} from "vue";
import topNavText from "@/views/AdminPage/components/topNavText.vue";


import statusTag from "@/views/AdminPage/components/statusTag.vue";


// 引入 用户添加 页面
import addUser from "./components/addUser.vue"
// 分页组件
import Pagination from "@/views/AdminPage/components/pagination.vue";

// addUserVisible 控制 添加页面显示
const addUserVisible = ref(false)


// 获取 用户表格信息


//todo 当 user_type 为用户 或者商户的 时候  tag有不同显示
const tableData = [
  {
    ID: 1,
    username: 'name1',
    phone: '12345678978',
    email: '123123112@qq.com',
    user_type: '用户',
    account_status: '正常'
  },
  {
    ID: 2,
    username: 'name2',
    phone: '12345678978',
    email: '123123112@qq.com',
    user_type: '商户',
    account_status: '封禁'
  },
]

// 分页管理

const currentPage = ref(1)
const itemCount = ref(100)

// 监听 currentPage 的变化
watch(currentPage, (newValue) => {
  console.log('Current Page changed from', 'to', newValue);
});


</script>

<template>
  <div class="sys">

    <topNavText :text="'用户管理-所有用户'"/>

    <el-card>
      <div style="margin-bottom: 10px">
        <el-input v-model="input" style="width: 240px;margin-right: 15px" placeholder="请输入搜索的用户ID" clearable/>
        <el-button type="primary">搜索</el-button>
        <el-button type="success" @click="addUserVisible = true">添加用户</el-button>
        <el-button type="danger">批量封禁</el-button>
      </div>

      <div>
        <el-table
            ref="multipleTableRef"
            :data="tableData"
            stripe
            style="width: 100%"

        >
          <el-table-column prop="ID" label="ID" width="55"/>
          <el-table-column prop="username" label="用户名" width="120"/>

          <el-table-column prop="phone" label="电话"/>
          <el-table-column prop="email" label="邮箱"/>
          <el-table-column prop="user_type" label="权限" width="120">
            <template #default="scope">
              <el-tag
                  :type="scope.row.user_type === '商户' ? '' : 'success'"
                  disable-transitions
              >{{ scope.row.user_type }}
              </el-tag
              >
            </template>
          </el-table-column>


          <el-table-column prop="account_status" label="状态" width="120">
            <template #default="scope">
              <statusTag :message="scope.row.account_status"/>
            </template>
          </el-table-column>


          <el-table-column fixed="right" label="更多操作">
            <template #default="scope">
              <el-button size="small" @click="handleEdit(scope.row)">编辑</el-button>
              <el-button size="small" type="primary" @click="handleEdit(scope.row)">解封</el-button>
              <el-button size="small" type="danger" @click="handleDelete(scope.row)">封禁</el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>

      <!-- 页数管理 -->

      <Pagination v-model:currentPage="currentPage" v-model:itemCount="itemCount"/>
    </el-card>

  </div>


  <addUser v-model="addUserVisible"/>

</template>

<style scoped>
.sys {
  height: auto;
}


</style>
