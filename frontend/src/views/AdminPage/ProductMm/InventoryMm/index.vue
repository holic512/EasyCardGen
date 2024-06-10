<!--用于后台管理-商品管理-库存管理-->
<script setup lang="ts">

import topNavText from "@/views/AdminPage/components/topNavText.vue";
import {ref} from "vue";
import {ElTable} from "element-plus";

//todo 库存管理的 逻辑是 单独的 表用来存储导入库存的 记录 同样单独的 表存储 每条库存

//用于 商品 库存管理的 接口
interface inventory_info {
  inventory_id: string,
  product_id: string,
  quantity: string,
  created_at: string,
  state: string,
  weight: string,
}

//用于 库存管理表格的 样板
const inventory_table = ref<inventory_info[]>([
  {
    inventory_id: '0001',
    product_id: '0001',
    quantity: '100',
    created_at: '2024-01-01',
    state: '正常',
    weight: '0',
  },
  {
    inventory_id: '0001',
    product_id: '0001',
    quantity: '100',
    created_at: '2024-01-01',
    state: '下架',
    weight: '0',
  },
  {
    inventory_id: '0001',
    product_id: '0001',
    quantity: '100',
    created_at: '2024-01-01',
    state: '封禁',
    weight: '0',
  },


])

//todo 用于处理搜索 商店的 功能
const input = ref('')

//todo 用于选择 处理
const multipleTableRef = ref<InstanceType<typeof ElTable>>()
const multipleSelection = ref<inventory_info[]>([])

const handleSelectionChange = (val: inventory_info[]) => {
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
  <topNavText :text="'商品管理-库存管理'"/>

  <el-card>
    <div style="margin-bottom: 10px">
      <el-input v-model="input" style="width: 240px;margin-right: 15px" placeholder="请输入搜索的提现id" clearable/>
      <el-button type="primary">搜索</el-button>
      <el-button type="success">添加库存</el-button>
      <el-button type="danger">批量封禁</el-button>
    </div>


    <el-table
        :data="inventory_table"
        stripe
        style="width: 100%"
        @selection-change="handleSelectionChange"
    >

      <el-table-column type="selection" width="55"/>

      <el-table-column prop="inventory_id" label="库存id"/>

      <el-table-column prop="product_id" label="商品id"/>

      <el-table-column prop="quantity" label="数量" width="100"/>

      <el-table-column prop="created_at" label="导入时间"/>

      <el-table-column prop="state" label="状态" width="100">

        <template #default="scope">
          <el-tag
              :type="getTagType(scope.row.state)"
              disable-transitions
          >
            {{ scope.row.state }}
          </el-tag>
        </template>

      </el-table-column>


      <el-table-column prop="quantity" label="权重" width="80"/>


      <el-table-column label="库存内容">
        <el-button>
          详情
        </el-button>

      </el-table-column>

      <el-table-column label="更多操作">
        <el-button type="danger">
          封禁
        </el-button>
        <el-button>
          编辑
        </el-button>

      </el-table-column>


    </el-table>
  </el-card>

</template>

<style scoped>

</style>