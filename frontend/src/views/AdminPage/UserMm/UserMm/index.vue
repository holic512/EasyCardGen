<!--用于管理员管理-用户管理-所有用户-->
<script setup lang="ts">

//添加用户页面
import {ref} from "vue";
import topNavText from "@/views/AdminPage/components/topNavText.vue";


import statusTag from "@/views/AdminPage/components/statusTag.vue";
import AddUser from "@/views/AdminPage/UserMm/UserMm/components/addUser.vue";

// 引入 用户添加 页面
import addUser from "./components/addUser.vue"
const dialogVisible = ref(false)


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

const addUserVisible = ref(false)


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
      <div class="endPagination">
        <el-pagination
            background
            v-model="currentPage"
            @current-change="getNowPageData"
            layout="prev, pager, next"
            :total="tablepage * 10"
        />
      </div>
    </el-card>

  </div>


  <addUser v-model="addUserVisible"/>

</template>

<style scoped>
.sys {
  height: auto;
}

.endPagination {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 60px;
  background-color: white;
}
</style>
