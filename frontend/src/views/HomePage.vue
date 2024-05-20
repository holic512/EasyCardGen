<script setup>

import {ChatSquare, CircleCheck} from "@element-plus/icons-vue";
</script>

<template>

  <div class="common-layout">
    <el-scrollbar height="100vh" style="background-color: rgb(244,244,244);">
      <el-container>
        <el-header style="padding: 0;">

          <div class="Nav">
            <el-menu :default-active="activeIndex" mode="horizontal" :ellipsis="false"
                     @select="handleSelect">

              <div style="margin-right: 5px;"/>

              <img style="width: 100px ;margin-left: 36px" src="../assets/images/HomePage/logo.png" alt="logo"/>

              <!-- 起一个 隔断 的 作用 -->
              <div class="flex-grow"/>

              <el-menu-item index="1" class="menuItem">网站首页</el-menu-item>
              <el-menu-item index="2" class="menuItem">卡密查询</el-menu-item>
              <el-menu-item index="3" class="menuItem">投诉进度</el-menu-item>
              <el-menu-item index="4" class="menuItem">联系我们</el-menu-item>
              <el-menu-item index="5" class="menuItem">常见问题</el-menu-item>

              <div class="flex-grow"/>
              <div style="display: flex;justify-content: center;align-content: center">

                <button class="custom-button" @click="Sign">登录/注册</button>

              </div>


            </el-menu>
          </div>
        </el-header>
        <el-main>
          <div class="page1">
            <!--左半部-->
            <div class="pageCon1" style="margin-right: -10%;">
              <div class="tagCon1">
                <el-tag round style="height: 35px;">
                  <el-tag effect="dark" type="danger" round class="tag1">
                    v1.0
                  </el-tag>
                  全平台Go开发,0.1ms响应
                </el-tag>

                <div class="text1">
                  <el-text tag="b" style="font-size: 250%; --el-text-color: #303133 ">支持</el-text>
                  <el-text tag="b" style="font-size: 200%; --el-text-color: #303133">多</el-text>
                  <el-text type="primary" tag="b" style="font-size: 250%;">商户</el-text>
                  <el-text tag="b" style="font-size: 200%; --el-text-color: #303133">的</el-text>
                  <el-text type="primary" tag="b" style="font-size: 250%;">发卡平台</el-text>
                  <br><br>
                  <el-text tag="b" style="font-size: 250%;">EasyCardGen</el-text>

                </div>
                <div style="margin-top: 5%;margin-bottom: 5%;">
                  <el-text style="font-size: large; ">
                    致力于虚拟商品的快捷销售服务-安全,可靠
                  </el-text>

                </div>
                <div>
                  <el-button color="#626aef" round size="large" style="font-weight: bold;">快速入住 成为商户</el-button>
                </div>
              </div>
            </div>

            <div class="pageCon1">
              <img src="../assets/images/HomePage/homePage1.png" alt="" class="pageImg1">
            </div>
          </div>


          <div style="display: flex; justify-content: center; /* 水平居中 */">
            <div>
              <div style="display: flex; justify-content: center; /* 水平居中 */">
                <el-text tag="b" style="font-size: 250%; --el-text-color: #303133 ">Easy CardGen</el-text>

              </div>
              <br><br>
              <div style="display: flex; justify-content: center; /* 水平居中 */">
                <el-text style="font-size: small; ">
                  欢迎使用EasyCardGen,一款安全,可靠的多商户的发卡平台,为您的购买与销售,提供更好的选择.
                </el-text>
              </div>

            </div>
          </div>

          <!--page1-->
          <div style=" display: flex; justify-content: center; align-items: center;margin-top: 80px ;">
            <div
                style="display: flex; justify-content: center; align-items: center;background-color: white; width: 90vw; max-width: 1000px; border-radius: 10px;box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);">
              <div style="background-color: white ;width: 80%;height: 80%;">
                <div class="page2">
                  <div style="margin-right: 2%;">
                    <img src="../assets/images/HomePage/homePage2.png" alt="" class="pageImg2">
                  </div>
                  <div>
                    <el-text tag="b" style="font-size: 200%;">开心发卡网</el-text>
                    <br>
                    <el-text type="primary" tag="b" style="font-size: 200%;">虚拟产品</el-text>

                    <el-text tag="b" style="font-size: 200%;">销售服务</el-text>
                    <br>
                    <div style="margin-top: 5%;margin-bottom: 5%;">
                      <el-text style="font-size: large; ">
                        注册成为商家,将卡密托管在平台,(软件注册码、游戏帐号、虚拟商品卡密等)
                        为商家提供安全、便捷、绿色、稳定虚拟商品卡密自动交易服务
                      </el-text>
                    </div>
                    <el-button color="#626aef" size="large" style="font-weight: bold;"> 立即体验 ></el-button>
                  </div>
                </div>

              </div>

            </div>
          </div>

          <div style=" display: flex; justify-content: center; align-items: center;margin-top: 80px ;">
            <div
                style="display: flex; justify-content: center; align-items: center;background-color: white; width: 90vw; max-width: 1000px; border-radius: 10px;box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);">
              <!--page1-->
              <div style="margin-right: 3%;">
                <el-text tag="b" style="font-size: 200%;">企业级发卡项目</el-text>
                <br>
                <el-text type="primary" tag="b" style="font-size: 200%;">响应时间低至0.1ms</el-text>
                <div style="margin-top: 5%;margin-bottom: 5%;">
                  <el-text style="font-size: large; ">
                    致力于为用户提供最好的虚拟产品交易服务
                  </el-text>
                </div>
                <div>
                  <el-text>
                    <el-icon>
                      <CircleCheck/>
                    </el-icon>
                    全平台Vue+Go开发
                  </el-text>
                  <br>
                  <el-text>
                    <el-icon>
                      <CircleCheck/>
                    </el-icon>
                    为商家提供一站式虚拟商品在线购买及全自动发货
                  </el-text>
                  <br>
                  <el-text>
                    <el-icon>
                      <CircleCheck/>
                    </el-icon>
                    平台系统功能齐全，支付即发卡，数万用户的信赖之选
                  </el-text>
                </div>

              </div>
              <div>
                <img src="../assets/images/HomePage/homePage3.png" alt="" class="pageImg3">
              </div>
            </div>

          </div>


          <div style=" display: flex; justify-content: center; align-items: center;margin-top: 80px ;">
            <div
                style="display: flex; justify-content: center; align-items: center;background-color: white; width: 90vw; max-width: 1000px; border-radius: 10px;box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);">
              <div style="background-color: white ;width: 80%;height: 80%;">
                <div class="page4">
                  <el-row style="width: 500px; margin-left: 5%;">
                    <el-col :span="6">
                      <el-statistic title="总注册人数" :value="2685"/>
                    </el-col>
                    <el-col :span="6">
                      <el-statistic title="总商户数" :value="78"/>
                    </el-col>
                    <el-col :span="6">
                      <el-statistic title="总商品数" :value="1720"/>
                    </el-col>
                    <el-col :span="6">
                      <el-statistic title="平台成交总数" :value="55"/>
                    </el-col>
                  </el-row>
                </div>
              </div>
            </div>
          </div>

          <div style=" display: flex; justify-content: center; align-items: center;margin-top: 80px ;">
            <div
                style="display: flex; justify-content: center; align-items: center;background-color: white; width: 90vw; max-width: 1000px; border-radius: 10px;box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);">
              <div style="background-color: white ;width: 80%;height: 80%;">
                <!-- page5 -->
                <div class="page5">
                  <div>
                    <img src="../assets/images/HomePage/homePage5.png" alt="" class="pageImg5">
                  </div>
                  <div>
                    <el-text tag="b" style="font-size: 200%;">开心发卡网</el-text>
                    <br>
                    <el-text type="primary" tag="b" style="font-size: 200%;">平台公告</el-text>
                    <br>
                    <div style="margin-top: 5%;margin-bottom: 5%;">
                      <el-text style="font-size: large; ">
                        我们将一些重要信息公布于此，您可以再此了解最新动态
                      </el-text>
                    </div>
                    <el-divider/>
                    <div class="pageCon5">
                      <div style="margin-top: 8px;margin-right: 10px;">
                        <el-icon :size="50">
                          <ChatSquare/>
                        </el-icon>
                      </div>
                      <div>
                        掌握平台最新咨询
                        <br>
                        <el-link href="#" type="primary">即刻进入,了解详情</el-link>

                      </div>

                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>

        </el-main>


        <el-footer>
          <!-- page-end -->
          <div class="pageEnd">
            Copyright © 2023 开心发卡网 All rights reserved. 版权所有
          </div>
        </el-footer>


      </el-container>
    </el-scrollbar>
  </div>
</template>

<style scoped>

.custom-button {
  --color: rgba(86, 11, 173, 0.99);
  font-family: inherit;
  display: inline-block;
  width: 8em;
  height: 2.6em;
  line-height: 2.3em;
  position: relative;
  overflow: hidden;
  border: 2px solid var(--color);
  transition: color .5s;
  z-index: 1;
  font-size: 17px;
  border-radius: 6px;
  font-weight: 500;
  color: var(--color);
  margin: 8px 30px 8px 8px;
}

.custom-button:before {
  content: "";
  position: absolute;
  z-index: -1;
  background: var(--color);
  height: 150px;
  width: 200px;
  border-radius: 50%;
}

.custom-button:hover {
  color: #fff;
}

.custom-button:before {
  top: 100%;
  left: 100%;
  transition: all .7s;
}

.custom-button:hover:before {
  top: -30px;
  left: -30px;
}

.custom-button:active:before {
  background: #ffffff;
  transition: background 0s;
}


.Nav {
  /* margin-bottom: 50px; */
  position: fixed;
  top: 0;
  width: 100%;

}

.flex-grow {
  flex-grow: 1;
}

.menuItem {
  font-size: 16px;
  font-weight: bold;
}

.page1 {
  display: flex;
  width: auto;
  height: 700px;
  justify-content: center;
  align-items: center;

}

.pageCon1 {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;


}

.pageImg1 {
  height: 380px;
  width: 500px;
  justify-content: center;
  align-items: center;
  margin-left: -10%;
}

.tag1 {
  margin-right: 5px;
}

.tagCon1 {

  margin-right: -10%;
}

.page2 {
  display: flex;
  width: auto;
  justify-content: center;
  align-items: center;
}

.pageImg2 {
  width: 400px;
  height: 400px;
}

.pageImg3 {
  width: 500px;
  height: 500px;
}

.page4 {
  display: flex;

  justify-content: center;
  /* 水平居中 */
  align-items: center;
  /* 垂直居中 */
  height: 100px;

}

.pageEnd {
  display: flex;
  justify-content: center;
  /* 水平居中 */
  align-items: center;
  /* 垂直居中 */
  height: 50px;
  background-color: rgb(32, 41, 66);
  font-size: 14px;
  color: rgb(177, 186, 195);

  margin-top: 40px;
}

.page5 {
  display: flex;
  /* 使用 Flexbox 布局 */
  /* 左右分布 */
  width: auto;
  height: 700px;
  justify-content: center;
  /* 水平居中 */
  align-items: center;
  /* 垂直居中 */

}

.pageImg5 {
  width: 400px;
  height: auto;
}

.pageCon5 {
  display: flex;
  /* 使用 Flexbox 布局 */

  justify-content: flex-start;
  /* 水平居中 */
  align-items: center;
  /* 垂直居中 */

}

.el-carousel__item h3 {
  color: #475669;
  opacity: 0.75;
  line-height: 200px;
  margin: 0;
  text-align: center;
}

</style>