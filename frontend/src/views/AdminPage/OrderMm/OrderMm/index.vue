<!--用于后台管理-订单管理-所有订单-->
<script setup lang="ts">

import topNavText from "@/views/AdminPage/components/topNavText.vue";
import {ref} from "vue";
import {ElTable} from "element-plus";

// order_info 用于 订单信息 的 接口
interface order_info {
  order_id: string,
  user_id: string,
  product_id: string,
  payment_method: string,
  original_price: string,
  purchase_quantity: string,
  actual_payment: string,
  discount_code: string,
  order_date: string,
  shipping_method: string,
  order_status: string,
}

// order_table 用于 模拟 数据表
const order_table = ref<order_info[]>([
  {
    order_id: '0001',
    user_id: '0001',
    product_id: '0001',
    payment_method: '微信支付',
    original_price: '100',
    purchase_quantity: '5',
    actual_payment: '500',
    discount_code: '',
    order_date: '2024-3-15',
    shipping_method: '自动发货',
    order_status: '待支付',
  },
  {
    order_id: '0001',
    user_id: '0001',
    product_id: '0001',
    payment_method: '微信支付',
    original_price: '100',
    purchase_quantity: '5',
    actual_payment: '500',
    discount_code: '',
    order_date: '2024-3-15',
    shipping_method: '自动发货',
    order_status: '待发货',
  },
  {
    order_id: '0001',
    user_id: '0001',
    product_id: '0001',
    payment_method: '微信支付',
    original_price: '100',
    purchase_quantity: '5',
    actual_payment: '500',
    discount_code: '',
    order_date: '2024-3-15',
    shipping_method: '自动发货',
    order_status: '待收货',
  },
  {
    order_id: '0001',
    user_id: '0001',
    product_id: '0001',
    payment_method: '微信支付',
    original_price: '100',
    purchase_quantity: '5',
    actual_payment: '500',
    discount_code: '',
    order_date: '2024-3-15',
    shipping_method: '自动发货',
    order_status: '已完成',
  },
  {
    order_id: '0001',
    user_id: '0001',
    product_id: '0001',
    payment_method: '微信支付',
    original_price: '100',
    purchase_quantity: '5',
    actual_payment: '500',
    discount_code: '',
    order_date: '2024-3-15',
    shipping_method: '自动发货',
    order_status: '售后中',
  },
])

// input 用于存储 搜索栏 中的信息
const input = ref('')

//用于 table 的 选择 处理
const multipleTableRef = ref<InstanceType<typeof ElTable>>()
const multipleSelection = ref<order_info[]>([])

const handleSelectionChange = (val: order_info[]) => {
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
  <topNavText :text="'订单管理-所有订单'"/>

  <el-card>
    <div style="margin-bottom: 10px">
      <el-input v-model="input" style="width: 240px;margin-right: 15px" placeholder="请输入搜索的提现id" clearable/>
      <el-button type="primary">搜索</el-button>
      <el-button type="success">添加库存</el-button>
      <el-button type="danger">批量封禁</el-button>
    </div>

    <el-table
        :data="order_table"
        stripe
        style="width: 100%"
        @selection-change="handleSelectionChange"
    >

      <el-table-column type="selection" width="55"/>

      <el-table-column prop="order_id" label="订单id"/>

      <el-table-column prop="user_id" label="用户id"/>

      <el-table-column prop="product_id" label="商品id"/>

      <el-table-column prop="payment_method" label="支付方式"/>

      <el-table-column prop="original_price" label="原价"/>

      <el-table-column prop="purchase_quantity" label="购买数量"/>

      <el-table-column prop="actual_payment" label="实际付款"/>

      <el-table-column prop="discount_code" label="使用折扣"/>

      <el-table-column prop="order_date" label="订单日期"/>

      <el-table-column prop="shipping_method" label="发货方式" width="100">

        <template #default="scope">

          <el-tag
              :type="getTagType(scope.row.state)"
              disable-transitions
          >
            {{ scope.row.shipping_method }}
          </el-tag>

        </template>

      </el-table-column>


      <el-table-column prop="order_status" label="订单状态" width="80">
        <template #default="scope">

          <el-tag
              :type="getTagType(scope.row.order_status)"
              disable-transitions
          >
            {{ scope.row.order_status }}
          </el-tag>

        </template>
      </el-table-column>

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