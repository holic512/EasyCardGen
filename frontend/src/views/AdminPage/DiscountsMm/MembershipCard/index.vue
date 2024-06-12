<!--用于后台管理-折扣管理-会员卡-->
<script setup lang="ts">

import topNavText from "@/views/AdminPage/components/topNavText.vue";
import {ref} from "vue";
import {ElTable} from "element-plus";

// memberCards 用于 会员卡 信息的接口
interface memberCards{
  member_card_id:string,
  user_id:string,
  store_id:string,
  balance:string,
  issue_date:string,
  status:string,
}

// memberCardsTable 用于 会员卡表格 数据样板
const memberCardsTable = ref<memberCards[]>([
  {
    member_card_id:'0001',
    user_id:'0001',
    store_id:'0001',
    balance:'100.5',
    issue_date:'2023-5-21',
    status:'正常',
  },

])

// input 用于存储 搜索栏 中的信息
const input = ref('')

//用于 table 的 选择 处理
const multipleTableRef = ref<InstanceType<typeof ElTable>>()
const multipleSelection = ref<memberCards[]>([])

const handleSelectionChange = (val: memberCards[]) => {
  multipleSelection.value = val
}

// statusTag 用于 状态组件
import statusTag from "@/views/AdminPage/components/statusTag.vue";

</script>

<template>
  <topNavText :text="'折扣管理-会员卡'"/>

  <el-card>
    <div style="margin-bottom: 10px">
      <el-input v-model="input" style="width: 240px;margin-right: 15px" placeholder="请输入搜索的提现id" clearable/>
      <el-button type="primary">搜索</el-button>
      <el-button type="success">添加会员卡</el-button>
      <el-button type="danger">批量禁用</el-button>
    </div>

    <el-table
        :data="memberCardsTable"
        stripe
        style="width: 100%"
        @selection-change="handleSelectionChange"
    >

      <el-table-column type="selection" width="55"/>

      <el-table-column prop="member_card_id" label="会员卡id"/>

      <el-table-column prop="user_id" label="用户id"/>

      <el-table-column prop="store_id" label="商店id"/>

      <el-table-column prop="balance" label="余额"/>

      <el-table-column prop="issue_date" label="开卡时间"/>

      <el-table-column prop="status" label="状态">
        <template #default="scope">
          <statusTag :message="scope.row.status" />
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