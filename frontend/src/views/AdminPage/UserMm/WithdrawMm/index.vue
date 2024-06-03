<!--用于管理员管理-用户管理-提现管理-->
<script setup lang="ts">

import topNavText from "@/views/AdminPage/components/topNavText.vue";


import {ref} from "vue";
import {ElTable, TabsPaneContext} from "element-plus";

const activeName = ref('first')

const handleClick = (tab: TabsPaneContext, event: Event) => {
  console.log(tab, event);
};


//输入框测试
const input = ref()

interface info {
  ID: string
  storeId: string
  withdrawalMethod: string
  withdrawalAmount: string
  status: string
  initiatedTime: string
  remarks: string
}

const tableData = ref<info[]>(
    [
      {
        ID: "123",
        storeId: "string",
        withdrawalMethod: "string",
        withdrawalAmount: "string",
        status: "string",
        initiatedTime: "123",
        remarks: "string",
      }
    ]
);

const multipleTableRef = ref<InstanceType<typeof ElTable>>()
const multipleSelection = ref<info[]>([])

const handleSelectionChange = (val: info[]) => {
  multipleSelection.value = val
}

</script>

<template>

  <topNavText :text="'用户管理-提现管理'"/>
  <el-card>

    <el-tabs v-model="activeName" class="demo-tabs" @tab-click="handleClick">
      <el-tab-pane label="未处理" name="first">

        <div style="margin-bottom: 10px">
          <el-input v-model="input" style="width: 240px;margin-right: 15px" placeholder="请输入搜索的提现id" clearable/>
          <el-button type="primary">搜索</el-button>
          <el-button type="primary">批量同意</el-button>
          <el-button type="primary">批量拒绝</el-button>
        </div>


        <el-table
            :data="tableData"
            stripe
            style="width: 100%"
            @selection-change="handleSelectionChange"
        >

          <el-table-column type="selection" width="55"/>

          <el-table-column prop="ID" label="提现id"/>

          <el-table-column prop="storeId" label="商户id"/>

          <el-table-column prop="withdrawalMethod" label="提现方式"/>

          <el-table-column prop="withdrawalAmount" label="提现金额"/>

          <el-table-column prop="status" label="状态"/>

          <el-table-column prop="initiatedTime" label="发起时间"/>

          <el-table-column prop="remarks" label="备注"/>

          <!--          <el-table-column prop="address" label="操作"/>-->


        </el-table>

      </el-tab-pane>

      <el-tab-pane label="已处理" name="second">
        Config
      </el-tab-pane>

      <el-tab-pane label="全部" name="third">
        Role
      </el-tab-pane>
    </el-tabs>

  </el-card>
</template>

<style scoped>
.demo-tabs > .el-tabs__content {
  padding: 32px;
  color: #6b778c;
  font-size: 32px;
  font-weight: 600;
}
</style>
