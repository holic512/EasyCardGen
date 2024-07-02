<!--用于商店管理-权限管理-->
<script setup lang="ts">
// 标签文字显示
import topNavText from "@/views/AdminPage/components/topNavText.vue";
import {ref} from "vue";
import {ElTable} from "element-plus";

// 引用 状态 tag
import StatusTag from './components/StatusTag.vue';
import AddStoreAccess from "@/views/AdminPage/StoreMm/AccessMm/components/addStoreAccess.vue";

//用用于权限管理约束
interface access_info {
  access_id: string,
  access_name: string,
  max_class_count: string,
  max_product_count: string,
  recommendation_weight: string,
  withdrawal_fee: string,
  discount_code: string,
  membership_card: string,
  keyword_assistant_reply: string,
  exclusive_platform_customer_service: string,
}

//todo 商品权限信息模板
const access_table = ref<access_info[]>([
  {
    access_id: '0001',
    access_name: '普通商店',
    max_class_count: '100',
    max_product_count: '20',
    recommendation_weight: '3',
    withdrawal_fee: '0.3',
    discount_code: 'true',
    membership_card: 'true',
    keyword_assistant_reply: 'true',
    exclusive_platform_customer_service: 'true',
  },
  {
    access_id: '0002',
    access_name: '普通商店',
    max_class_count: '100',
    max_product_count: '20',
    recommendation_weight: '3',
    withdrawal_fee: '0.3',
    discount_code: 'false',
    membership_card: 'true',
    keyword_assistant_reply: 'true',
    exclusive_platform_customer_service: 'true',
  },
  {
    access_id: '0003',
    access_name: '普通商店',
    max_class_count: '100',
    max_product_count: '20',
    recommendation_weight: '3',
    withdrawal_fee: '0.3',
    discount_code: 'false',
    membership_card: 'true',
    keyword_assistant_reply: 'false',
    exclusive_platform_customer_service: 'true',
  },

])

// addStoreVisible 用于控制 dialog 显示
let addStoreVisible = ref(false)
</script>

<template>

  <topNavText :text="'商店管理-权限管理'"/>

  <el-card>

    <!--上操作栏-->
    <div style="margin-bottom: 10px">

      <el-input v-model="input" style="width: 240px;margin-right: 15px" placeholder="请输入搜索的提现id" clearable/>
      <el-button type="primary">搜索</el-button>
      <el-button type="success" @click="addStoreVisible = true">创建权限</el-button>
      <el-button type="danger">批量停用</el-button>

    </div>

    <el-table
        :data="access_table"
        stripe
        style="width: 100%"
    >

      <el-table-column prop="access_id" label="权限ID"/>

      <el-table-column prop="access_name" label="权限名称"/>

      <el-table-column prop="max_class_count" label="分类数" width="80"/>

      <el-table-column prop="max_product_count" label="商品数" width="80"/>

      <el-table-column prop="recommendation_weight" label="权重" width="55"/>

      <el-table-column prop="withdrawal_fee" label="提现费" width="80"/>

      <el-table-column prop="discount_code" label="折扣码" width="80">

        <template #default="scope">
          <StatusTag :value="scope.row.discount_code"/>
        </template>

      </el-table-column>

      <el-table-column prop="membership_card" label="会员卡" width="80">
        <template #default="scope">
          <StatusTag :value="scope.row.membership_card"/>
        </template>
      </el-table-column>

      <el-table-column prop="keyword_assistant_reply" label="自动回复" width="80">
        <template #default="scope">
          <StatusTag :value="scope.row.keyword_assistant_reply"/>
        </template>

      </el-table-column>
      <el-table-column prop="exclusive_platform_customer_service" label="平台客服" width="80">

        <template #default="scope">
          <StatusTag :value="scope.row.exclusive_platform_customer_service"/>
        </template>
      </el-table-column>

      <el-table-column label="更多操作">
        <el-button size="small">
          停用
        </el-button>
        <el-button size="small">
          编辑
        </el-button>
      </el-table-column>


    </el-table>
  </el-card>

  <AddStoreAccess v-model="addStoreVisible"/>

</template>

<style scoped>

</style>