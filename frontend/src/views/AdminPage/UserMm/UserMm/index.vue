<!--用于管理员管理-用户管理-所有用户-->
<script setup lang="ts">

//添加用户页面
import {onMounted, reactive, ref, watch} from "vue";
import topNavText from "@/views/AdminPage/components/topNavText.vue";


import statusTag from "@/views/AdminPage/components/statusTag.vue";


// 引入 用户添加 页面
import addUser from "./components/addUser.vue"
// 分页组件
import Pagination from "@/views/AdminPage/components/pagination.vue";
import {ElTable} from "element-plus";
import axios from "axios";

// addUserVisible 控制 添加页面显示
let addUserVisible = ref(false)


// 分页管理
let currentPage = ref(1)
let itemCount = ref(100)

//用于获取用户数目
const getUserCount = () => {
  axios.get('http://localhost:8080/api/admin/getUserCount')
      .then(function (res) {
        if (res.status === 200) {
        }
        itemCount.value = res.data.count
      })
      .catch(function (error) {
        console.log(error)
      })
}
onMounted(() => {
  getUserCount()
})

// 监听 currentPage 的变化
watch(currentPage, (newValue) => {
  console.log('Current Page changed from', 'to', newValue);
});

//根据页数 获取用户数据
let tableData = reactive([]);
const getUserInfo = () => {
  axios.get('http://localhost:8080/api/admin/getUserInfo', {
    params: {
      currentPage: currentPage.value
    }
  })
      .then(function (res) {
        if (res.status === 200) {

          // 使用扩展运算符更新数据
          tableData.splice(0, tableData.length, ...res.data.userData);
        }
      })
      .catch(function (error) {
        console.log(error)
      })

}
onMounted(() => {
  getUserInfo()
})

const handleSelectionChange = (val: withdrawPending[]) => {
  multipleSelection.value = val
}

const getTypeMsg = (msg: string) => {
  switch (msg) {
    case 'user':
      return '用户'

    case 'merchant':
      return '商户'

    case 'admin':
      return '管理员'
  }
}

const getType = (msg: string) => {
  switch (msg) {
    case '用户':
      return 'success'

    case '商户':
      return ''

    case '管理员':
      return 'danger'
  }
}

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
            @selection-change="handleSelectionChange"
        >
          <el-table-column type="selection" width="55"/>
          <el-table-column prop="ID" label="ID" width="55"/>
          <el-table-column prop="Username" label="用户名" width="100"/>

          <el-table-column prop="Phone" label="电话" width="140"/>
          <el-table-column prop="Email" label="邮箱" width="220"/>
          <el-table-column prop="UserType" label="权限" width="120">
            <template #default="scope">
              <el-tag
                  :type="getType(getTypeMsg(scope.row.UserType) )"
                  disable-transitions
              >{{ getTypeMsg(scope.row.UserType) }}
              </el-tag
              >
            </template>
          </el-table-column>


          <el-table-column prop="State" label="状态" width="120">
            <template #default="scope">
              <statusTag :message="scope.row.State"/>
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
