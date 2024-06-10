<!--用于后台管理-商品管理-分类管理-->
<script setup lang="ts">
import topNavText from "@/views/AdminPage/components/topNavText.vue";
import {ref} from "vue";
import {ElTable} from "element-plus";

//todo 用于 处理分类信息的 约束
interface categories_info {
  category_id: string;
  store_id: string;
  state: string;
  category_name: string;
  weight: string;
  created_at: string;
}

//分类管理信息模板
const categories_table = ref<categories_info[]>([
  {
    category_id: '0001',
    store_id: 'string',
    state: '正常',
    category_name: 'string',
    weight: 'string',
    created_at: 'string',
  },
  {
    category_id: '0002',
    store_id: 'string',
    state: '停用',
    category_name: 'string',
    weight: 'string',
    created_at: 'string',
  },
  {
    category_id: '0003',
    store_id: 'string',
    state: '封禁',
    category_name: 'string',
    weight: 'string',
    created_at: 'string',
  },
]);


//todo 用于处理搜索 商店的 功能
const input = ref('')


//todo 用于选择 处理
const multipleTableRef = ref<InstanceType<typeof ElTable>>()
const multipleSelection = ref<categories_info[]>([])

const handleSelectionChange = (val: categories_info[]) => {
  multipleSelection.value = val
}

// 处理tag
function getTagType(message: string) {
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
  <topNavText :text="'商品管理-分类管理'"/>

  <el-card>

    <!--上操作栏-->
    <div style="margin-bottom: 10px">
      <el-input v-model="input" style="width: 240px;margin-right: 15px" placeholder="请输入搜索的提现id" clearable/>
      <el-button type="primary">搜索</el-button>
      <el-button type="success">添加分类</el-button>
      <el-button type="danger">批量封禁</el-button>
    </div>


    <el-table
        :data="categories_table"
        stripe
        style="width: 100%"
        @selection-change="handleSelectionChange"
    >

      <el-table-column type="selection" width="55"/>

      <el-table-column prop="category_id" label="分类id"/>

      <el-table-column prop="store_id" label="商店id"/>

      <el-table-column prop="category_name" label="名称"/>


      <el-table-column prop="state" label="状态">

        <template #default="scope">
          <el-tag
              :type="getTagType(scope.row.state)"
              disable-transitions
          >
            {{ scope.row.state }}
          </el-tag>
        </template>

      </el-table-column>

      <el-table-column prop="weight" label="权重"/>

      <el-table-column prop="created_at" label="创建时间"/>

      <el-table-column label="更多信息">
        详情

      </el-table-column>

      <el-table-column label="更多操作">
        封禁 编辑
      </el-table-column>


    </el-table>

  </el-card>
</template>

<style scoped>

</style>