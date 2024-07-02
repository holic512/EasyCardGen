<script setup lang="ts">
import {reactive, ref} from "vue";
import {FormInstance} from "element-plus";
import {FormRules} from "element-plus";
// addStoreVisible 用于控制 dialog 显示
let addStoreVisible = ref(false)

// storeAccess 用于存储表单信息
let storeAccess = ref({
  AccessName: '',
  Type: 'primary',
  MaxClassCount: '',
  MaxProductCount: '',
  Weight: '',
  WithdrawalFee: '',
  DiscountCode: false,
  MembershipCard: false,
  KeywordReply: false,
  PlatformService: false,
})

// 存储样式滑块
const TypeOptions = [
  {
    label: '初级',
    value: 'primary'
  },
  {
    label: '中级',
    value: 'success'
  },
  {
    label: '高级',
    value: 'warning'
  },
  {
    label: '终极',
    value: 'danger'
  },
]

// ruleFormRef 用于存储 表格 总规则
const ruleFormRef = ref<FormInstance>()

// validateAccessName 用于验证 权限名称
const validateAccessName = (rule: any, value: any, callback: any) => {
  if (value === '') {
    callback(new Error('权限名称不能为空'))
  } else {
    if (value.length < 4 || value.length > 12) {
      callback(new Error('权限名称要至少4位且不大于12位'))
    } else callback()
  }
}

// 用于实现 表单规则
const rules = reactive<FormRules>({
  AccessName: [{validator: validateAccessName, trigger: 'blur'}],
})

</script>

<template>
  <el-dialog v-model="addStoreVisible" title="添加新权限">
    <el-form
        :model="storeAccess"
        :label-position="'right'"
        label-width="auto"
        :inline="true"

        ref="ruleFormRef"
        :rules="rules"

    >
      <el-form-item
          label="权限名称"
          prop="AccessName"
      >
        <el-input
            v-model="storeAccess.AccessName"
            placeholder="请输入添加的权限名称"
        />
      </el-form-item>
      <br>

      <el-form-item
          label="颜色样式"
          prop="Type"
      >
        <el-segmented v-model="storeAccess.Type" :options="TypeOptions">
          <template #default="{ item }">
            <div>{{ item.label }}</div>
          </template>
        </el-segmented>

      </el-form-item>
      <br>

      <el-form-item
          label="最大分类数"
          prop="MaxClassCount"
      >
        <el-input-number v-model="storeAccess.MaxClassCount"/>

      </el-form-item>

      <el-form-item
          label="最大商品数"
          prop="MaxProductCount"
      >
        <el-input-number v-model="storeAccess.MaxProductCount"/>
      </el-form-item>
      <br>

      <el-form-item label="推荐位权重" prop="Weight">
        <el-input-number v-model="storeAccess.Weight"/>
      </el-form-item>

      <el-form-item label="提现手续费" prop="WithdrawalFee">
        <el-input-number v-model="storeAccess.WithdrawalFee"/>
      </el-form-item>

      <br>
      <el-form-item label="折扣码功能" prop="DiscountCode">
        <el-switch v-model="storeAccess.DiscountCode"/>
      </el-form-item>

      <el-form-item label="会员卡功能" prop="MembershipCard">
        <el-switch v-model="storeAccess.MembershipCard"/>
      </el-form-item>

      <el-form-item label="关键词回复" prop="KeywordReply">
        <el-switch v-model="storeAccess.KeywordReply"/>
      </el-form-item>

      <el-form-item label="平台客服" prop="PlatformService">
        <el-switch v-model="storeAccess.PlatformService"/>
      </el-form-item>
      <br>
      <el-form-item>
        <el-button
            type="primary"
            @click="submitForm(ruleFormRef)"
        >
          提交
        </el-button>

        <el-button @click="resetForm(ruleFormRef)">重置</el-button>

      </el-form-item>

    </el-form>
  </el-dialog>
</template>

<style scoped>
</style>