<!-- 此项目用来 后台管理->商品管理->商品管理 -->
<template>
  <div class="sys">
    <div class="navbar">
      <a
          style="
          color: black;
          margin-top: 6px;
          margin-bottom: 6px;
          font-weight: bold;
          font-size: 18px;
        "
      >用户管理</a
      >
    </div>
    <div class="topButtonBox">
      <el-button
          color="#626aef"
          @click="
          dialogFormVisible = true;
          resetFormDataAndImages();
        "
      >添加
      </el-button
      >
      <el-button type="danger" @click="deleteSelectionChange">删除</el-button>
    </div>

    <div>
      <el-table
          ref="multipleTableRef"
          :data="tableData"
          :table-layout="tableLayout"
          style="width: 100%"
          @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="40"/>
        <el-table-column prop="id" label="ID"/>
        <el-table-column prop="username" label="用户名"/>
        <el-table-column prop="password" label="密码"/>
        <el-table-column prop="status" label="状态"/>
        <el-table-column prop="balance" label="余额"/>
        <el-table-column prop="merchantId" label="商户id"/>
        <el-table-column prop="phone" label="电话"/>
        <el-table-column prop="email" label="邮箱"/>
        <el-table-column prop="inviteCode" label="邀请码"/>
        <el-table-column prop="inviter" label="邀请人"/>
        <el-table-column prop="lastLogin" label="上次登录"/>
        <!--          <el-table-column fixed="right" label="状态">-->
        <!--            <template #default="scope">-->
        <!--              <el-tag v-if="scope.row.tag === 'true'" class="ml-2" type="success"-->
        <!--              >运行-->
        <!--              </el-tag-->
        <!--              >-->
        <!--              <el-tag v-else-if="scope.row.tag === 'false'" class="ml-2" type="danger"-->
        <!--              >停止-->
        <!--              </el-tag-->
        <!--              >-->
        <!--            </template>-->
        <!--          </el-table-column>-->
        <el-table-column fixed="right" label="操作">
          <template #default="scope">
            <el-button size="small" @click="handleEdit(scope.row)">编辑</el-button>
            <el-button size="small" type="danger" @click="handleDelete(scope.row)"
            >删除
            </el-button
            >
          </template>
        </el-table-column>
      </el-table>
    </div>
    <!-- 页数管理 -->
    <div class="endPagination">
      <el-pagination
          background
          v-model="currentPage"
          @current-change="getNowPageData"
          layout="prev, pager, next"
          :total="tablepage * 10"
      />
    </div>
  </div>

  <!-- 用与添加分类的 from  -->
  <el-dialog v-model="dialogFormVisible" title="添加新的商品">
    <el-form ref="formData" :model="formData">
      <el-form-item label="商品图标" :label-width="formLabelWidth">
        <el-upload
            v-if="this.filesIcon === null"
            class="upload"
            drag
            :before-upload="beforeUpload"
        >
          <div class="el-upload__text">请将图标拖拽到此处 <em>点击上传</em></div>
          <template #tip>
            <div class="el-upload__tip">支持jpg格式,64*64,50kb</div>
          </template>
        </el-upload>
        <!-- 当检测到 成功上传图片时 则关闭 上传框 打开 展示框 -->
        <!-- <img v-else :src="filesIcon" alt="icon" /> -->
        <el-avatar v-else shape="square" :size="100" :fit="cover" :src="filesIcon"/>
      </el-form-item>

      <el-form-item label="商品名称" :label-width="formLabelWidth">
        <el-input v-model="formData.name" autocomplete="off"/>
      </el-form-item>

      <el-form-item label="商品说明" :label-width="formLabelWidth">
        <el-input v-model="formData.info" autocomplete="off"/>
      </el-form-item>

      <el-form-item label="商品分类" :label-width="formLabelWidth">
        <el-select v-model="formData.classId" placeholder="请选择">
          <el-option
              v-for="option in this.class"
              :label="option.name"
              :key="option.id"
              :value="option.id"
          >
            <template #default>
              <!-- 使 列中 展示 的 内容  分类名称靠左  分类状态 靠右 -->
              <div style="display: flex; justify-content: space-between">
                <div>
                  <el-text class="test-1">{{ option.name }}</el-text>
                </div>
                <div>
                  <el-tag v-if="option.tag === 'true'" class="ml-2" type="success"
                  >运行
                  </el-tag
                  >
                  <el-tag v-else-if="option.tag === 'false'" class="ml-2" type="danger"
                  >停止
                  </el-tag
                  >
                </div>
              </div>
            </template>
          </el-option>
        </el-select>
      </el-form-item>

      <el-form-item label="成本价" :label-width="formLabelWidth">
        <el-input v-model.number.trim="formData.costPrice" autocomplete="off"/>
      </el-form-item>

      <el-form-item label="商品单价" :label-width="formLabelWidth">
        <el-input v-model.number.trim="formData.price" autocomplete="off"/>
      </el-form-item>

      <el-form-item label="会员价" :label-width="formLabelWidth">
        <el-input v-model.number.trim="formData.priceVip" autocomplete="off"/>
      </el-form-item>

      <el-form-item label="发货方式" :label-width="formLabelWidth">
        <el-select v-model="formData.deliveryMethod" placeholder="请选择">
          <el-option label="自动发货" value="true"/>
          <el-option label="手动发货" value="false"/>
        </el-select>
      </el-form-item>

      <el-form-item
          v-if="formData.deliveryMethod === 'true'"
          label="发货留言"
          :label-width="formLabelWidth"
      >
        <el-input v-model="formData.deliveryMessage" autocomplete="off"/>
      </el-form-item>

      <el-form-item label="卡密排序" :label-width="formLabelWidth">
        <el-select v-model="formData.deliverySort" placeholder="请选择">
          <el-option label="优先发旧" value="0"/>
          <el-option label="随机发送" value="1"/>
          <el-option label="优先发新" value="2"/>
        </el-select>
      </el-form-item>

      <el-form-item label="权重" :label-width="formLabelWidth">
        <el-input-number v-model="formData.sort" :min="0" :max="10"/>
      </el-form-item>

      <el-form-item label="状态" :label-width="formLabelWidth">
        <el-select v-model="formData.tag" placeholder="请选择状态">
          <el-option label="展示" value="true"/>
          <el-option label="隐藏" value="false"/>
        </el-select>
      </el-form-item>

      <el-form-item label="联系方式" :label-width="formLabelWidth">
        <el-select v-model="formData.contact" placeholder="请选择">
          <el-option label="邮箱发货" value="0"/>
          <el-option label="短信发货" value="1"/>
        </el-select>
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="this.dialogFormVisible = false">取消</el-button>
        <el-button type="primary" @click="addNewCommodity"> 确认 </el-button>
      </span>
    </template>
  </el-dialog>
</template>
<script>
import {ElNotification} from "element-plus";
import {ref} from "vue";
// 拖拽上传组件

import axios from "axios";

export default {
  components: {},
  data() {
    return {
      // 用于判断 添加窗口是否展示
      dialogFormVisible: ref(true),
      formLabelWidth: "80px",

      //  存储 图片 文件
      files: null,
      filesIcon: null,

      // 用来储存 添加 表单 数据 样式
      form: {
        id: 0,
        classId: "",
        name: "",
        info: "",
        costPrice: "",
        price: "",
        priceVip: "",
        deliveryMethod: "",
        deliveryMessage: "",
        deliverySort: null,
        sort: 0,
        tag: "",
        contact: null,
      },
      // 用于储存 表格内容
      formData: {},

      // 用于存储 商品信息
      commodity: [],

      // 用于存储分类信息
      class: [],
    };
  },
  methods: {
    // 钩子函数 在 组件构建前 执行
    created() {
      // 使 formData 初始化 为 form
      this.formData = {...this.form};
    },

    // 图片上传 前 进行判断
    beforeUpload(file) {
      const isJPG = file.type === "image/jpeg";
      if (!isJPG) {
        ElNotification({
          title: "文件格式错误",
          message: "仅支持类型为jpg的文件",
          type: "error",
        });

        return false;
      }

      // 若符合数据类型，将选中的文件存储在组件数据中
      this.files = file;
      this.filesIcon = URL.createObjectURL(file);
      console.log(this.files);

      return false; // 阻止自动上传
    },

    // 添加 新的 商品
    addNewCommodity() {
      if (
          this.files != null &&
          this.formData.classId !== "" &&
          this.formData.name !== "" &&
          this.formData.info !== "" &&
          this.formData.costPrice !== "" &&
          this.formData.price !== "" &&
          this.formData.priceVip !== "" &&
          this.formData.deliveryMethod !== "" &&
          (this.formData.deliveryMethod === "true"
              ? this.formData.deliveryMessage !== ""
              : true) &&
          this.formData.deliverySort !== null &&
          this.formData.sort >= 0 &&
          this.formData.tag !== "" &&
          this.formData.contact !== null
      ) {
        const formData = new FormData();
        formData.append("files", this.files);

        // 使用JSON.stringify将对象转为字符串形式
        const commodityData = JSON.stringify(this.formData);
        formData.append("commodity", commodityData);

        // formData.append("commodity", this.formData);

        axios
            .post("http://localhost:8085/api/admin/addCommodity", formData)
            .then((response) => {
            });
      } else {
        ElNotification({
          title: "请填写完整数据",
          message: "请填写完整数据",
          type: "warning",
        });
      }
    },

    // 重置 formData,图片
    resetFormDataAndImages() {
      this.formData = {...this.form};
      this.files = null;
      this.filesIcon = null;
    },
  },
  mounted() {
    axios.get("http://localhost:8085/api/admin/commodityGetClass").then((response) => {
      this.class = response.data;
      console.log(this.class);
    });
  },
};
</script>
<style scoped>
.sys {
  height: auto;

}

.navbar {
  display: flex;
  justify-content: flex-start;
}

.topButtonBox {
  margin-top: 1%;
  display: flex;
  /* 靠右 */
  justify-content: flex-end;
  align-items: center;
  height: 40px;
  padding-right: 5%;
  padding-top: 1%;
}

.endPagination {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 60px;
  background-color: white;
}

.upload {
  height: 100%;
  width: 25%;
}
</style>
