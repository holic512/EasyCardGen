<script setup lang="ts">
import {ref} from "vue";
// addStoreVisible 用来控制 dialog显示
let addStoreVisible = ref(false)


</script>

<template>
  <el-dialog v-model="addStoreVisible" title="添加商店">

  </el-dialog>
</template>

<style scoped>

</style>