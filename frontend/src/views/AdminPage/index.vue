<script setup>
import {ref} from 'vue'

// 图标引用
import {
  Location,
  Menu as IconMenu,
  More,
  Picture,
  Present,
  Document,
  Setting,
  Shop,
  ShoppingBag,
  User,
  Back
} from '@element-plus/icons-vue'


// 用于缩放左导航
const isCollapse = ref(true)

function toggleCollapse() {
  isCollapse.value = !isCollapse.value
  // 在这里可以添加其他你希望执行的逻辑
}

//用于跳转对应路由
import {useRouter} from 'vue-router'

const router = useRouter()

function routerIndex(index) {
  switch (index) {
    case '1':
      // 仪表盘
      router.push('/admin/')
      break

    case '2-1':
      // 用户管理
      router.push('/admin/userMm')
      break

    case '2-2':
      //权限管理
      router.push('/admin/Withdraw')
      break

  }

}

</script>
<template>
  <div class="common-layout">
    <el-container>

      <!--侧边栏-->
      <el-aside width="auto" style="height: 100vh;">
        <!--滚动条-->
        <el-scrollbar>
          <el-menu default-active="1" style="height: 95vh;" :collapse="isCollapse">
            <el-menu-item index="1" @click="routerIndex('1')">
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
              <el-menu-item index="2-1" @click="routerIndex('2-1')">所有用户</el-menu-item>
              <el-menu-item index="2-2" @click="routerIndex('2-2')">提现管理</el-menu-item>
            </el-sub-menu>

            <el-sub-menu index="3">
              <template #title>
                <el-icon>
                  <Shop/>
                </el-icon>
                <span>店铺管理</span>
              </template>
              <el-menu-item index="3-1" @click="routerIndex('3-1')">所有店铺</el-menu-item>
              <el-menu-item index="3-2" @click="routerIndex('3-2')">权限管理</el-menu-item>
            </el-sub-menu>

            <el-sub-menu index="4">
              <template #title>
                <el-icon>
                  <ShoppingBag/>
                </el-icon>
                <span>商品管理</span>
              </template>
              <el-menu-item index="4-1" @click="routerIndex('4-1')">分类管理</el-menu-item>
              <el-menu-item index="4-2" @click="routerIndex('4-2')">商品管理</el-menu-item>
              <el-menu-item index="4-3" @click="routerIndex('4-3')">库存管理</el-menu-item>
            </el-sub-menu>

            <el-sub-menu index="5">
              <template #title>
                <el-icon>
                  <Document/>
                </el-icon>
                <span>订单管理</span>
              </template>
              <el-menu-item index="5-1" @click="routerIndex('5-1')">全部订单</el-menu-item>
              <el-menu-item index="5-2" @click="routerIndex('5-2')">售后管理</el-menu-item>
            </el-sub-menu>


            <el-menu-item index="6" @click="routerIndex('6')">
              <el-icon>
                <Picture/>
              </el-icon>
              <template #title>广告管理</template>
            </el-menu-item>

            <el-sub-menu index="7">
              <template #title>
                <el-icon>
                  <Present/>
                </el-icon>
                <span>折扣管理</span>
              </template>
              <el-menu-item index="6-1" @click="routerIndex('7-1')">会员卡</el-menu-item>
              <el-menu-item index="6-2" @click="routerIndex('7-2')">折扣码</el-menu-item>
            </el-sub-menu>


            <el-sub-menu index="8">
              <template #title>
                <el-icon>
                  <location/>
                </el-icon>
                <span>工单管理</span>
              </template>
              <el-menu-item index="7-1" @click="routerIndex('8-1')">商户工单</el-menu-item>
              <el-menu-item index="7-2" @click="routerIndex('8-2')">用户投诉</el-menu-item>
            </el-sub-menu>

            <el-menu-item index="9" @click="routerIndex('9')">
              <el-icon>
                <More/>
              </el-icon>
              <template #title>插件配置</template>
            </el-menu-item>

            <el-menu-item index="10" @click="routerIndex('10')">
              <el-icon>
                <Setting/>
              </el-icon>
              <template #title>总站设置</template>
            </el-menu-item>

            <el-menu-item index="11" @click="routerIndex('11')">
              <el-icon>
                <Back/>
              </el-icon>
              <template #title>退出登录</template>
            </el-menu-item>
          </el-menu>
        </el-scrollbar>
      </el-aside>

      <!--缩放按钮-->
      <div style="background-color: rgb(244,244,244);">
        <div class="ScaleBar">
          <div @click="toggleCollapse">
            <div class="line"></div>
          </div>
        </div>
      </div>


      <!--用于显示主页面的-->
      <el-main style="  background-color: rgb(244,244,244);">
        <router-view></router-view>
      </el-main>


    </el-container>
  </div>
</template>


<style scoped>
.ScaleBar {
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
