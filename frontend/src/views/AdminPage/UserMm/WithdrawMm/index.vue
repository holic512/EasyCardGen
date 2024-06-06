<!--用于管理员管理-用户管理-提现管理-->
<script setup lang="ts">

import topNavText from "@/views/AdminPage/components/topNavText.vue";
import {ref} from "vue";
import {TabsPaneContext} from "element-plus";

//未处理页面
import Pending from "../WithdrawMm/components/Pending.vue";

//已处理页面
import Completed from "../WithdrawMm/components/Completed.vue";

//全部页面
import All from "../WithdrawMm/components/All.vue";

const activeName = ref('first')

const handleClick = (tab: TabsPaneContext, event: Event) => {
  console.log(tab, event);
};

//todo 优化代码重复 更新操作按钮


</script>

<template>

  <topNavText :text="'用户管理-提现管理'"/>
  <el-card>

    <el-tabs v-model="activeName" class="demo-tabs" @tab-click="handleClick">
      <el-tab-pane label="未处理" name="first">

        <Pending/>

      </el-tab-pane>

      <el-tab-pane label="已处理" name="second">

        <Completed/>

      </el-tab-pane>

      <el-tab-pane label="全部" name="third">

        <All/>

      </el-tab-pane>
    </el-tabs>

  </el-card>
</template>

<style scoped>
.demo-tabs > .el-tabs__content {
  padding: 32px;
  color: #6b778c;
  font-size: 32px;
  font-weight: 600;
}
</style>
