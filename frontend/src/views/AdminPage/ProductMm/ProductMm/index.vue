<!--用于后台管理-商品管理-所有商品-->
<script setup lang="ts">

import topNavText from "@/views/AdminPage/components/topNavText.vue";
import {ref} from "vue";
import {ElTable} from "element-plus";

//todo 用于 商品管理的接口
interface products_info {
  products_id: string,
  category_id: string,
  products_name: string,
  products_state: string,
  products_price: string,
  quantity: string,
  created_at: string,
  delivery_method: string,
  weight: string,
}

//todo 用于 商品表格数据样板
const products_table = ref<products_info[]>([
  {
    products_id: '0001',
    category_id: '0001',
    products_name: '0001',
    products_state: '正常',
    products_price: '1000',
    quantity: '100',
    created_at: '2023-3-21',
    delivery_method: '自动发货',
    weight: '0',
  },
  {
    products_id: '0002',
    category_id: '0002',
    products_name: '0002',
    products_state: '封禁',
    products_price: '1000',
    quantity: '100',
    created_at: '2023-3-21',
    delivery_method: '手动发货',
    weight: '0',
  },
  {
    products_id: '0002',
    category_id: '0002',
    products_name: '0002',
    products_state: '下架',
    products_price: '1000',
    quantity: '100',
    created_at: '2023-3-21',
    delivery_method: '手动发货',
    weight: '0',
  },
])

//todo 用于处理搜索 商店的 功能
const input = ref('')

//todo 用于选择 处理
const multipleTableRef = ref<InstanceType<typeof ElTable>>()
const multipleSelection = ref<products_info[]>([])

const handleSelectionChange = (val: products_info[]) => {
  multipleSelection.value = val
}

// 处理tag
function getTagType(message: string) {
  switch (message) {
    case '正常':
      return 'success'
    case '下架':
      return ''
    case '封禁':
      return 'danger'
    default:
      return
  }

}

</script>

<template>
  <topNavText :text="'商品管理-所有商品'"/>

  <el-card>

    <div style="margin-bottom: 10px">
      <el-input v-model="input" style="width: 240px;margin-right: 15px" placeholder="请输入搜索的提现id" clearable/>
      <el-button type="primary">搜索</el-button>
      <el-button type="success">添加商品</el-button>
      <el-button type="danger">批量封禁</el-button>
    </div>


    <el-table
        :data="products_table"
        stripe
        style="width: 100%"
        @selection-change="handleSelectionChange"
    >

      <el-table-column type="selection" width="55"/>

      <el-table-column prop="products_id" label="商品id"/>

      <el-table-column prop="category_id" label="分类id"/>

      <el-table-column prop="products_name" label="名称"/>


      <el-table-column prop="products_state" label="状态">

        <template #default="scope">
          <el-tag
              :type="getTagType(scope.row.products_state)"
              disable-transitions
          >
            {{ scope.row.products_state }}
          </el-tag>
        </template>

      </el-table-column>

      <el-table-column prop="products_price" label="价格"/>

      <el-table-column prop="quantity" label="剩余库存"/>

      <el-table-column prop="created_at" label="创建时间"/>

      <el-table-column prop="delivery_method" label="发货方式">

        <template #default="scope">
          <el-tag
              :type="scope.row.delivery_method === '自动发货'?'success':''"
              disable-transitions
          >
            {{ scope.row.delivery_method }}
          </el-tag>
        </template>

      </el-table-column>

      <el-table-column prop="weight" label="权重"/>


      <el-table-column label="更多信息">
        详情

      </el-table-column>

      <el-table-column label="更多操作">
        封禁 编辑
      </el-table-column>


    </el-table>

  </el-card>


</template>

<style scoped>

</style>