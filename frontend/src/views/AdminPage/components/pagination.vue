<script setup lang="ts">
//用于处理翻页
const currentPage = defineModel('currentPage')
const itemCount = defineModel('itemCount')


</script>

<template>
  <div class="endPagination">
    <el-pagination
        background
        v-model:current-page="currentPage"
        :total="itemCount"
        :page-size="10"
        layout="total, prev, pager, next,jumper"

    />
  </div>
</template>

<style scoped>
.endPagination {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 60px;
  background-color: white;
}
</style>