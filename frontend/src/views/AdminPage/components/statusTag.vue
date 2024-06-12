<script lang="ts" setup>

// defineProps 用于 接受 信息
const props = defineProps<{
  message?: String,
}>()


function getStatusTagType(message: string) {
  switch (message) {
    case '正常':
      return 'success'
    case '停用':
      return ''
    case '封禁':
      return 'danger'
    default:
      return

  }


}
</script>

<template>

  <el-tag
      :type="getStatusTagType(props.message)"
      disable-transitions
  >
    {{ props.message }}

  </el-tag>
</template>
