<!--用于管理员管理-用户管理-未处理-->
<script setup lang="ts">
import {ref} from "vue";
import {ElTable} from "element-plus";

//输入框测试
const input = ref()

// withdrawPending 用于约束 未处理 提现
interface withdrawPending {
  ID: string
  storeId: string
  withdrawalMethod: string
  withdrawalAmount: string
  status: string
  initiatedTime: string
  remarks: string
}

const withdrawPendingTableData = ref<withdrawPending[]>(
    [
      {
        ID: "0001",
        storeId: "0001",
        withdrawalMethod: "微信",
        withdrawalAmount: "100",
        status: "已拒绝",
        initiatedTime: "123",
        remarks: "100",
      },
      {
        ID: "0002",
        storeId: "0001",
        withdrawalMethod: "支付宝",
        withdrawalAmount: "100",
        status: "已完成",
        initiatedTime: "123",
        remarks: "1000",
      },
      {
        ID: "0003",
        storeId: "0001",
        withdrawalMethod: "其他",
        withdrawalAmount: "100",
        status: "待处理",
        initiatedTime: "123",
        remarks: "1000",
      }
    ]
);

const multipleTableRef = ref<InstanceType<typeof ElTable>>()
const multipleSelection = ref<withdrawPending[]>([])

const handleSelectionChange = (val: withdrawPending[]) => {
  multipleSelection.value = val
}

function getMethodType(message: string) {
  switch (message) {
    case '微信':
      return 'success'
    case '支付宝':
      return 'primary'
    case '其他':
      return 'info'
    default:
      return
  }
}

//用来处理 状态 的 tag标签type
function getTagType(message: string) {

  switch (message) {
    case '已完成':
      return 'success'
    case '待处理':
      return ''
    case '已拒绝':
      return 'danger'
    default:
      return
  }


}
</script>

<template>
  <div style="margin-bottom: 10px">
    <el-input v-model="input" style="width: 240px;margin-right: 15px" placeholder="请输入搜索的提现id" clearable/>
    <el-button type="primary">搜索</el-button>
    <el-button type="success">批量同意</el-button>
    <el-button type="danger">批量拒绝</el-button>
  </div>


  <el-table
      :data="withdrawPendingTableData"
      stripe
      style="width: 100%"
      @selection-change="handleSelectionChange"
  >

    <el-table-column type="selection" width="55"/>

    <el-table-column prop="ID" label="提现ID" />

    <el-table-column prop="storeId" label="商户ID"/>

    <el-table-column prop="withdrawalMethod" label="提现方式" width="120">

      <template #default="scope">
        <el-tag
            :type="getMethodType(scope.row.withdrawalMethod)"
            disable-transitions
        >
          {{ scope.row.withdrawalMethod }}
        </el-tag>
      </template>


    </el-table-column>

    <el-table-column prop="withdrawalAmount" label="提现金额"/>

    <el-table-column prop="status" label="处理状态" width="120">
      <template #default="scope">
        <el-tag
            :type="getTagType(scope.row.status)"
            disable-transitions
        >
          {{ scope.row.status }}
        </el-tag>
      </template>
    </el-table-column>

    <el-table-column prop="initiatedTime" label="发起时间"/>

    <el-table-column prop="remarks" label="备注"/>

    <el-table-column prop="address" label="其他操作">
      <template #default="scope">
        <el-button type="success" size="small">同意</el-button>
        <el-button type="danger" size="small">拒绝</el-button>
      </template>
    </el-table-column>


  </el-table>

</template>

<style scoped>


</style>