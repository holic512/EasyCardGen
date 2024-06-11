<!--用于后台管理-订单管理-售后管理-->
<script setup lang="ts">

import topNavText from "@/views/AdminPage/components/topNavText.vue";
import {ref} from "vue";
import {ElTable} from "element-plus";

interface after_sales_info {
  after_sales_id: string,
  order_id: string,
  after_sales_type: string,
  created_at: string,
  updated_at: string,
  rejection_reason: string,
  status: string,
}

const after_sales_table = ref<after_sales_info[]>([
  {
    after_sales_id: 'string',
    order_id: 'string',
    after_sales_type: 'string',
    created_at: 'string',
    updated_at: 'string',
    rejection_reason: 'string',
    status: 'string',
  }
])

// input 用于存储 搜索栏 中的信息
const input = ref('')

//用于 table 的 选择 处理
const multipleTableRef = ref<InstanceType<typeof ElTable>>()
const multipleSelection = ref<after_sales_info[]>([])

const handleSelectionChange = (val: after_sales_info[]) => {
  multipleSelection.value = val
}

// getTagType 用于 处理 tag 样式
function getTagType(message: string) {
  switch (message) {
    case '待支付':
      return 'info'
    case '待发货':
      return 'primary'
    case '待收货':
      return 'danger'
    case '已完成':
      return 'success'
    case '售后中':
      return 'warning'
    default:
      return
  }
}

</script>

<template>


  <topNavText :text="'订单管理-售后管理'"/>

  <el-card>

    <div style="margin-bottom: 10px">
      <el-input v-model="input" style="width: 240px;margin-right: 15px" placeholder="请输入搜索的提现id" clearable/>
      <el-button type="primary">搜索</el-button>
      <el-button type="success">批量同意</el-button>
      <el-button type="danger">批量驳回</el-button>
    </div>

    <el-table
        :data="after_sales_table"
        stripe
        style="width: 100%"
        @selection-change="handleSelectionChange"
    >

      <el-table-column type="selection" width="55"/>

      <el-table-column prop="after_sales_id" label="售后id"/>

      <el-table-column prop="order_id" label="订单id"/>

      <el-table-column prop="after_sales_type" label="售后类型"/>

      <el-table-column prop="created_at" label="创建时间"/>

      <el-table-column prop="updated_at" label="最后处理时间"/>

      <el-table-column prop="status" label="售后状态"/>

      <el-table-column label="更多操作">
        <el-button type="danger" size="small">
          封禁
        </el-button>
        <el-button size="small">
          详情
        </el-button>

      </el-table-column>


    </el-table>

  </el-card>

</template>

<style scoped>

</style>