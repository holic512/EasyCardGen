<!--用于管理员管理-用户管理-提现管理-->
<script setup lang="ts">

import topNavText from "@/views/AdminPage/components/topNavText.vue";


import {ref} from "vue";
import type {TabsPaneContext} from "element-plus";

const activeName = ref('first')

const handleClick = (tab: TabsPaneContext, event: Event) => {
  console.log(tab, event);
};

</script>

<template>

  <topNavText :text="'用户管理-提现管理'"/>
  <el-card>
    <el-tabs v-model="activeName" class="demo-tabs" @tab-click="handleClick">
      <el-tab-pane label="未处理" name="first">

        <el-table :data="tableData" stripe style="width: 100%">
          <el-table-column prop="date" label="Date" width="180" />
          <el-table-column prop="name" label="Name" width="180" />
          <el-table-column prop="address" label="Address" />
        </el-table>

      </el-tab-pane>

      <el-tab-pane label="已处理" name="second">
        Config
      </el-tab-pane>

      <el-tab-pane label="全部" name="third">
        Role
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
