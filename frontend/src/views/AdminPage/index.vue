<template>
  <div class="common-layout">
    <el-container>

      <el-aside width="auto" style="height: 95vh;">
        <!--滚动条-->
        <el-scrollbar>
          <!--侧边栏-->
          <el-menu default-active="2" style="height: 95vh;" :collapse="isCollapse">
            <el-menu-item index="1" @click="index = 1">
              <el-icon>
                <icon-menu/>
              </el-icon>
              <template #title>仪表盘</template>
            </el-menu-item>

            <el-sub-menu index="2">
              <template #title>
                <el-icon>
                  <User/>
                </el-icon>
                <span>用户管理</span>
              </template>
              <el-menu-item index="2-1" @click="index = '2-1'">所有用户</el-menu-item>
              <el-menu-item index="2-2" @click="index = 2-2">所有商户</el-menu-item>
              <el-menu-item index="2-3" @click="index = 2-3">提现管理</el-menu-item>
            </el-sub-menu>

            <el-menu-item index="3" @click="index = 3">
              <el-icon>
                <Shop/>
              </el-icon>
              <template #title>所有店铺</template>
            </el-menu-item>

            <el-sub-menu index="4">
              <template #title>
                <el-icon>
                  <ShoppingBag/>
                </el-icon>
                <span>商品管理</span>
              </template>
              <el-menu-item index="4-1" @click="index = 4-1">所有分类</el-menu-item>
              <el-menu-item index="4-2" @click="index = 4-2">所有商品</el-menu-item>
              <el-menu-item index="4-3" @click="index = 4-3">所有卡密</el-menu-item>
            </el-sub-menu>

            <el-menu-item index="5" @click="index = 5">
              <el-icon>
                <Picture/>
              </el-icon>
              <template #title>广告管理</template>
            </el-menu-item>

            <el-menu-item index="6" @click="index = 6">
              <el-icon>
                <Present/>
              </el-icon>
              <template #title>折扣优惠</template>
            </el-menu-item>

            <el-sub-menu index="7">
              <template #title>
                <el-icon>
                  <location/>
                </el-icon>
                <span>工单管理</span>
              </template>
              <el-menu-item index="7-1" @click="index = 7-1">商户工单</el-menu-item>
              <el-menu-item index="7-2" @click="index = 7-2">用户投诉</el-menu-item>
            </el-sub-menu>

            <el-menu-item index="8" @click="index = 8">
              <el-icon>
                <More/>
              </el-icon>
              <template #title>插件配置</template>
            </el-menu-item>

            <el-menu-item index="9" @click="index = 9">
              <el-icon>
                <Promotion/>
              </el-icon>
              <template #title>支付接口</template>
            </el-menu-item>

            <el-menu-item index="10" @click="index = 10">
              <el-icon>
                <Setting/>
              </el-icon>
              <template #title>总站设置</template>
            </el-menu-item>

            <el-menu-item index="11" @click="index = 11">
              <el-icon>
                <Back/>
              </el-icon>
              <template #title>退出登录</template>
            </el-menu-item>
          </el-menu>
        </el-scrollbar>

      </el-aside>


      <!--缩放按钮-->
      <div class="Scalebar">
        <div @click="toggleCollapse">
          <div class="line"></div>
        </div>
      </div>


      <el-main>
        <div v-if="index === 1">123</div>
        <merchant-mm v-else-if="index === '2-1'"></merchant-mm>
        <AllSetting v-else-if="index === 10"/>
      </el-main>
    </el-container>
  </div>
</template>
<script>

import AllSetting from './AllSetting/index.vue'
import MerchantMm from "./UserMm/MerchantMm/index.vue";

import {
  Back,
  Document,
  Location,
  Menu as IconMenu,
  More,
  Picture,
  Present,
  Promotion,
  Setting,
  Shop,
  ShoppingBag,
  User
} from '@element-plus/icons-vue'

export default {
  data() {
    return {
      isCollapse: true,

      index: "1",
    };
  },
  components: {
    Document,
    Location,
    IconMenu,
    Setting,
    User,
    More,
    Present,
    ShoppingBag,
    Shop,
    Picture,
    Promotion,
    Back,


    AllSetting,
    MerchantMm

  },
  methods: {
    handleOpen(key, keyPath) {
      console.log(key, keyPath)
    },
    handleClose(key, keyPath) {
      console.log(key, keyPath)
    },


    toggleCollapse() {
      this.isCollapse = !this.isCollapse;
      // 在这里可以添加其他你希望执行的逻辑
    }
  }
}
</script>


<style>

.Scalebar {
  margin-top: 40vh; /* 将元素顶部与视口垂直中心对齐 */
}

.line {
  margin-left: 5px;
  height: 65px;
  width: 6px;
  border-radius: 9999px; /* 圆角 */
  background-color: rgb(137, 135, 135); /* 灰色 */
  cursor: pointer; /* 鼠标悬停时显示指针 */
}

</style>
