<!--用于商店管理-所有商店-->
<script setup lang="ts">

import topNavText from "@/views/AdminPage/components/topNavText.vue";
import {ref} from "vue";
import {ElTable} from "element-plus";

// statusTag 用于 状态组件
import statusTag from "@/views/AdminPage/components/statusTag.vue";

//todo 规定商店数据类型
interface store_info {
  store_id: string,
  user_id: string,
  store_name: string,
  store_type: string,
  store_state: string,
  store_product_count: string,
  store_balance: string,
  store_rating: string,
}

//todo 用来模拟数据插入
const store_info = ref<store_info[]>([
      {
        store_id: '0001',
        user_id: '0001',
        store_name: '开心发卡',
        store_type: '基础权限',
        store_state: '正常',
        store_product_count: '',
        store_balance: '',
        store_rating: '',
      },
      {
        store_id: '0001',
        user_id: '0001',
        store_name: '开心发卡',
        store_type: '中级权限',
        store_state: '停用',
        store_product_count: '',
        store_balance: '',
        store_rating: '',
      },
      {
        store_id: '0001',
        user_id: '0001',
        store_name: '开心发卡',
        store_type: '高级权限',
        store_state: '封禁',
        store_product_count: '100',
        store_balance: '34.5',
        store_rating: '4.5',
      }
    ]
)


//todo 用于处理搜索 商店的 功能
const input = ref('')


//todo 用于选择 处理
const multipleTableRef = ref<InstanceType<typeof ElTable>>()
const multipleSelection = ref<store_info[]>([])

const handleSelectionChange = (val: store_info[]) => {
  multipleSelection.value = val
}
</script>

<template>

  <topNavText :text="'商店管理-所有商店'"/>

  <el-card>
    <!--上操作栏-->
    <div style="margin-bottom: 10px">
      <el-input v-model="input" style="width: 240px;margin-right: 15px" placeholder="请输入搜索的提现id" clearable/>
      <el-button type="primary">搜索</el-button>
      <el-button type="success">创建商店</el-button>
      <el-button type="danger">批量封禁</el-button>
    </div>

    <el-table
        :data="store_info"
        stripe
        style="width: 100%"
        @selection-change="handleSelectionChange"
    >

      <el-table-column type="selection" width="55"/>

      <el-table-column prop="store_id" label="商店ID"/>

      <el-table-column prop="user_id" label="用户ID"/>

      <el-table-column prop="store_name" label="商店名称"/>

      <el-table-column prop="store_type" label="权限等级"/>

      <el-table-column prop="store_state" label="状态">

        <template #default="scope">
          <statusTag :message="scope.row.store_state"/>
        </template>

      </el-table-column>

      <el-table-column prop="store_product_count" label="商品数目"/>

      <el-table-column prop="store_balance" label="余额"/>

      <el-table-column prop="store_rating" label="评分"/>

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