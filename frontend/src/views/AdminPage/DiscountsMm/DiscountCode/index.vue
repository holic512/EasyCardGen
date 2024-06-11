<!--用于后台管理-折扣管理-折扣码-->
<script setup lang='ts'>

import topNavText from "@/views/AdminPage/components/topNavText.vue";
import {ref} from "vue";
import {ElTable} from "element-plus";

// discountCodes 用于 折扣码信息的 约束
interface discountCodes {
  discount_id: string,
  code: string,
  store_id: string,
  duration_mode: string,
  stock: string,
  status: string,
}

// discountCodesTable 用于 折扣码表格 数据样板
const discountCodesTable = ref<discountCodes[]>([
  {
    discount_id: 'string',
    code: 'string',
    store_id: 'string',
    duration_mode: '限时',
    stock: '5',
    status: '正常',
  },
  {
    discount_id: 'string',
    code: 'string',
    store_id: 'string',
    duration_mode: '永久',
    stock: '5',
    status: '停用',
  },
  {
    discount_id: 'string',
    code: 'string',
    store_id: 'string',
    duration_mode: '限时',
    stock: '5',
    status: '禁用',
  }
])

// input 用于存储 搜索栏 中的信息
const input = ref('')

//用于 table 的 选择 处理
const multipleTableRef = ref<InstanceType<typeof ElTable>>()
const multipleSelection = ref<discountCodes[]>([])

const handleSelectionChange = (val: discountCodes[]) => {
  multipleSelection.value = val
}


</script>

<template>
  <topNavText :text="'折扣管理-折扣码'"/>


  <el-card>
    <div style="margin-bottom: 10px">
      <el-input v-model="input" style="width: 240px;margin-right: 15px" placeholder="请输入搜索的提现id" clearable/>
      <el-button type="primary">搜索</el-button>
      <el-button type="success">添加折扣码</el-button>
      <el-button type="danger">批量禁用</el-button>
    </div>

    <el-table
        :data="discountCodesTable"
        stripe
        style="width: 100%"
        @selection-change="handleSelectionChange"
    >

      <el-table-column type="selection" width="55"/>

      <el-table-column prop="discount_id" label="折扣id"/>

      <el-table-column prop="code" label="折扣码"/>

      <el-table-column prop="store_id" label="商店id"/>

      <el-table-column prop="duration_mode" label="使用模式"/>

      <el-table-column prop="stock" label="库存"/>

      <el-table-column prop="status" label="状态"/>


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