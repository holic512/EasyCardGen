<!--用于商店管理-权限管理-->
<script setup lang="ts">
// 标签文字显示
import topNavText from "@/views/AdminPage/components/topNavText.vue";
import {ref} from "vue";
import {ElTable} from "element-plus";

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


// 引用tag
import StatusTag from './components/StatusTag.vue';

</script>

<template>

  <topNavText :text="'商店管理-权限管理'"/>

  <el-card>
    <el-table
        :data="access_table"
        stripe
        style="width: 100%"
    >

      <el-table-column prop="access_id" label="权限id"/>

      <el-table-column prop="access_name" label="权限名称"/>

      <el-table-column prop="max_class_count" label="最大分类数"/>

      <el-table-column prop="max_product_count" label="最大商品数"/>

      <el-table-column prop="recommendation_weight" label="推荐位权重"/>

      <el-table-column prop="withdrawal_fee" label="提现手续费"/>

      <el-table-column prop="discount_code" label="折扣码">

        <template #default="scope">
          <StatusTag :value="scope.row.discount_code"/>
        </template>

      </el-table-column>

      <el-table-column prop="membership_card" label="会员卡">
        <template #default="scope">
          <StatusTag :value="scope.row.membership_card"/>
        </template>
      </el-table-column>

      <el-table-column prop="keyword_assistant_reply" label="关键词助手">
        <template #default="scope">
          <StatusTag :value="scope.row.keyword_assistant_reply"/>
        </template>

      </el-table-column>
      <el-table-column prop="exclusive_platform_customer_service" label="专属平台客服">

        <template #default="scope">
          <StatusTag :value="scope.row.exclusive_platform_customer_service"/>
        </template>
      </el-table-column>

      <el-table-column label="更多操作">
        封禁 编辑
      </el-table-column>


    </el-table>
  </el-card>

</template>

<style scoped>

</style>