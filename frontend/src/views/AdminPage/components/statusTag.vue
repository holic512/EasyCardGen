<script setup lang="ts">
import { defineProps, computed, ref } from 'vue';

const props = defineProps<{
  message?: String,
}>()

// 使用 ref 定义状态
const message1 = ref("");

// 计算属性来获取标签类型
const tagType = computed(() => {
  switch (props.message) {
    case '正常':
    case 'active':
      message1.value = '正常';
      return 'success';
    case '停用':
      return '';
    case '封禁':
      return 'danger';
    default:
      return '';
  }
});
</script>

<template>
  <el-tag
      :type="tagType"
      disable-transitions
  >
    {{ message1 }}
  </el-tag>
</template>
