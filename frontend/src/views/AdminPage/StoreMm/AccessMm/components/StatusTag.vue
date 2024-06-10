<script>
export default {
  props: {
    value: {
      type: String,
      required: true,
    }
  },
  computed: {
    tagType() {
      return this.value === 'true' ? 'success' : 'info';
    },
    displayText() {
      return this.value === 'true' ? '启用' : '停用';
    }
  }
}
</script>

<template>
  <el-tag :type="tagType" disable-transitions>
    {{ displayText }}
  </el-tag>
</template>

<style scoped>

</style>