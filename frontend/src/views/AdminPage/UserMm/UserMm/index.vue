<!--用于管理员管理-用户管理-所有用户-->
<script setup>
//添加用户页面
import {ref} from "vue";
import topNavText from "@/views/AdminPage/components/topNavText.vue";
const dialogVisible = ref(false)
function openAddForm() {
  dialogVisible.value = true
}

</script>

<template>
  <div class="sys">

    <topNavText :text="'用户管理-所有用户'"/>

    <el-card>
      <div>
        <el-table
            ref="multipleTableRef"
            :data="tableData"
            :table-layout="tableLayout"
            style="width: 100%"
        >
          <el-table-column prop="ID" label="ID"/>
          <el-table-column prop="username" label="用户名"/>
          <el-table-column prop="phone" label="电话"/>
          <el-table-column prop="email" label="邮箱"/>
          <el-table-column prop="user_type" label="权限"/>
          <el-table-column prop="account_status" label="状态"/>
          <el-table-column fixed="right">
            <template #header>
              <!--添加用户-->

              <el-button color="#626aef" @click="openAddForm">添加用户</el-button>

            </template>
            <template #default="scope">
              <el-button size="small" @click="handleEdit(scope.row)">编辑</el-button>
              <el-button size="small" type="danger" @click="handleDelete(scope.row)">删除</el-button>
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
