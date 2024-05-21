<!--此文件用于 总站设置-->
<template>
  <div>

    <div class="div1">
      <el-text class="text1">设置</el-text>
    </div>

    <div>
      <el-link type="info" class="link1" @click="scrollToComponent(0)">系统</el-link>
      <el-link type="info" class="link1" @click="scrollToComponent(1)">注册</el-link>
      <el-link type="info" class="link1" @click="scrollToComponent(2)">商品</el-link>
      <el-link type="info" class="link1" @click="scrollToComponent(3)">邮箱</el-link>
      <el-link type="info" class="link1" @click="scrollToComponent(4)">短信</el-link>
      <el-link type="info" class="link1" @click="scrollToComponent(5)">工单</el-link>
      <el-link type="info" class="link1" @click="scrollToComponent(6)">提现</el-link>
      <el-link type="info" class="link1" @click="scrollToComponent(7)">充值</el-link>
      <el-link type="info" class="link1" @click="scrollToComponent(8)">备份</el-link>
      <el-link type="info" class="link1" @click="scrollToComponent(9)">管理员</el-link>
    </div>
    <el-divider/>
    <el-scrollbar height="78vh" ref="scrollbarRef">

      <!--系统-->
      <div ref="targetComponent1">
        <el-container class="con1">
          <el-aside width="160px">
            <el-text style="font-size: 16px;font-weight: bold">系统</el-text>

          </el-aside>
          <el-main>
            <div style="margin-bottom: 10px">
              <el-form label-width="100px" style="max-width: 460px" label-position="right">
                <el-form-item label="总站域名">
                  <el-input v-model="formData.name" placeholder="填写解析到此网站的域名"/>
                </el-form-item>
                <el-form-item label="总站名称">
                  <el-input v-model="formData.activityZone" placeholder="网站名称"/>
                </el-form-item>
                <el-form-item label="总站标题">
                  <el-input v-model="formData.activityForm" placeholder="网站标题"/>
                </el-form-item>
                <el-form-item label="Seo">
                  <el-input v-model="formData.activityForm" placeholder="网站关键词"/>
                </el-form-item>
                <el-form-item label="Logo">
                  <el-image style="width: 100px; height: 50px;border-radius: 5px"
                            :src="'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg'"
                            :fit="'cover'"
                  />
                  <el-button type="primary" circle style="margin-left: 10px">
                    <el-icon>
                      <Plus/>
                    </el-icon>
                  </el-button>

                </el-form-item>
                <el-form-item label="主题选择">
                  <el-select v-model="value" class="m-2" placeholder="选择您主题">
                    <el-option
                        v-for="item in options"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value"
                    />
                  </el-select>
                </el-form-item>
                <el-form-item label="后台背景">
                  <el-button circle style="margin-left: 10px">
                    <el-icon>
                      <Edit/>
                    </el-icon>
                  </el-button>
                </el-form-item>
                <el-form-item label="维护模式">
                  <el-switch v-model="value1"/>
                </el-form-item>
                <el-form-item label="维护公告">
                  <el-input v-model="formData.activityForm" placeholder="维护公告"/>
                </el-form-item>
              </el-form>
              <div style="margin-left: 20px;margin-top: 30px">
                <el-button type="primary" plain>保存配置</el-button>
                <el-button type="primary" plain>恢复初始配置</el-button>
              </div>
            </div>

          </el-main>
        </el-container>
        <el-divider/>
      </div>


      <!--注册配置-->
      <div ref="targetComponent2">
        <el-container class="con1">
          <el-aside width="160px">
            <el-text style="font-size: 16px;font-weight: bold">注册</el-text>

          </el-aside>
          <el-main>
            <div style="margin-bottom: 10px">
              <el-form label-width="120px" style="max-width: 460px" label-position="right">

                <el-form-item label="禁止注册">
                  <el-switch v-model="value1"/>
                </el-form-item>

                <el-form-item label="Google登录">
                  <el-switch v-model="value1"/>
                  <el-button circle style="margin-left: 10px">
                    <el-icon>
                      <Edit/>
                    </el-icon>
                  </el-button>
                </el-form-item>

                <el-form-item label="GitHub登录">
                  <el-switch v-model="value1"/>
                  <el-button circle style="margin-left: 10px">
                    <el-icon>
                      <Edit/>
                    </el-icon>
                  </el-button>
                </el-form-item>

                <el-form-item label="强制邮箱验证">
                  <el-switch v-model="value1"/>
                </el-form-item>

                <el-form-item label="手机验证">
                  <el-switch v-model="value1"/>
                </el-form-item>

                <el-form-item label="开通商户费用(元)">
                  <el-input-number v-model="num" :min="1" :max="10" @change="handleChange"/>
                </el-form-item>

                <el-form-item label="用户名注册长度">
                  <el-input-number v-model="num" :min="1" :max="10" @change="handleChange"/>
                </el-form-item>

              </el-form>
              <div style="margin-left: 20px;margin-top: 30px">
                <el-button type="primary" plain>保存配置</el-button>
                <el-button type="primary" plain>恢复初始配置</el-button>
              </div>
            </div>
          </el-main>

        </el-container>
        <el-divider/>
      </div>
      <!--商品管理-->
      <div ref="targetComponent3">
        <el-container class="con1">
          <el-aside width="160px">
            <el-text style="font-size: 16px;font-weight: bold">商品</el-text>

          </el-aside>
          <el-main>
            <div style="margin-bottom: 10px">
              <el-form label-width="100px" style="max-width: 460px" label-position="right">
                <el-form-item label="违规词(;分隔)">
                  <el-input v-model="formData.name" placeholder="禁止出现的违规词"/>
                </el-form-item>
                <el-form-item label="零库存超售">
                  <el-switch v-model="value1"/>
                </el-form-item>
              </el-form>
              <div style="margin-left: 20px;margin-top: 30px">
                <el-button type="primary" plain>保存配置</el-button>
                <el-button type="primary" plain>恢复初始配置</el-button>
              </div>
            </div>

          </el-main>
        </el-container>
        <el-divider/>
      </div>

      <!--邮箱-->
      <div ref="targetComponent4">
        <el-container class="con1">
          <el-aside width="160px">
            <el-text style="font-size: 16px;font-weight: bold">邮箱</el-text>

          </el-aside>
          <el-main>
            <div style="margin-bottom: 10px">
              <el-form label-width="100px" style="max-width: 460px" label-position="right">

              </el-form>
              <div style="margin-left: 20px;margin-top: 30px">
                <el-button type="primary" plain>发送测试邮件</el-button>
                <el-button type="primary" plain>保存设置</el-button>
              </div>
            </div>

          </el-main>
        </el-container>
        <el-divider/>
      </div>
      <!--短信-->
      <div ref="targetComponent5">
        <el-container class="con1">
          <el-aside width="160px">
            <el-text style="font-size: 16px;font-weight: bold">短信</el-text>

          </el-aside>
          <el-main>
            <div style="margin-bottom: 10px">
              <el-form label-width="100px" style="max-width: 460px" label-position="right">

              </el-form>
              <div style="margin-left: 20px;margin-top: 30px">
                <el-button type="primary" plain>发送测试短信</el-button>
                <el-button type="primary" plain>保存设置</el-button>
              </div>
            </div>

          </el-main>
        </el-container>
        <el-divider/>
      </div>

      <!--工单管理-->
      <div ref="targetComponent6">
        <el-container class="con1">
          <el-aside width="160px">
            <el-text style="font-size: 16px;font-weight: bold">工单</el-text>

          </el-aside>
          <el-main>
            <div style="margin-bottom: 10px">
              <el-form label-width="140px" style="max-width: 460px" label-position="right">

                <el-form-item label="工单开关">
                  <el-switch v-model="value1"/>
                </el-form-item>

                <el-form-item label="自动发送到邮箱">
                  <el-switch v-model="value1"/>
                </el-form-item>

                <el-form-item label="投诉开关">
                  <el-switch v-model="value1"/>
                </el-form-item>

                <el-form-item label="自动发送到邮箱">
                  <el-switch v-model="value1"/>
                </el-form-item>

              </el-form>
              <div style="margin-left: 20px;margin-top: 30px">
                <el-button type="primary" plain>保存配置</el-button>
                <el-button type="primary" plain>恢复初始配置</el-button>
              </div>
            </div>

          </el-main>
        </el-container>
        <el-divider/>
      </div>

      <!--提现-->
      <div ref="targetComponent7">
        <el-container class="con1">
          <el-aside width="160px">
            <el-text style="font-size: 16px;font-weight: bold">提现</el-text>

          </el-aside>
          <el-main>
            <div style="margin-bottom: 10px">
              <el-form label-width="140px" style="max-width: 460px" label-position="right">

                <el-form-item label="提现开关">
                  <el-switch v-model="value1"/>
                </el-form-item>


                <el-form-item label="自动提现">
                  <el-switch v-model="value1"/>
                </el-form-item>

                <el-form-item label="提现信息发送邮箱">
                  <el-switch v-model="value1"/>
                </el-form-item>

                <el-form-item label="提现费率">
                  <el-input v-model="input1" placeholder="提现费率">
                    <template #append>%</template>
                  </el-input>
                </el-form-item>

              </el-form>
              <div style="margin-left: 20px;margin-top: 30px">
                <el-button type="primary" plain>保存配置</el-button>
                <el-button type="primary" plain>恢复初始配置</el-button>
              </div>
            </div>

          </el-main>
        </el-container>
        <el-divider/>
      </div>

      <!--充值-->
      <div ref="targetComponent8">
        <el-container class="con1">
          <el-aside width="160px">
            <el-text style="font-size: 16px;font-weight: bold">充值</el-text>

          </el-aside>
          <el-main>
            <div style="margin-bottom: 10px">
              <el-form label-width="100px" style="max-width: 460px" label-position="right">
                <el-form-item label="充值开关">
                  <el-switch v-model="value1"/>
                </el-form-item>
                <el-form-item label="支付接口">
                  <el-select v-model="value" class="m-2" placeholder="选择支付接口">
                    <el-option
                        v-for="item in options"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value"
                    />
                  </el-select>
                </el-form-item>


                <el-form-item label="储值赠送">

                  <el-input v-model="formData.name" placeholder="请填写赠送规则"/>
                  <el-text>例如100送10,200送50(100-110;200-250;)</el-text>
                </el-form-item>


              </el-form>
              <div style="margin-left: 20px;margin-top: 30px">
                <el-button type="primary" plain>保存配置</el-button>
                <el-button type="primary" plain>恢复初始配置</el-button>
              </div>
            </div>

          </el-main>
        </el-container>
        <el-divider/>
      </div>

      <!--备份-->
      <div ref="targetComponent9">
        <el-container class="con1">
          <el-aside width="160px">
            <el-text style="font-size: 16px;font-weight: bold">备份</el-text>

          </el-aside>
          <el-main>
            <div style="margin-bottom: 10px">
              <el-form label-width="100px" style="max-width: 460px" label-position="right">
                <el-form-item label="定时备份开关">
                  <el-switch v-model="value1"/>
                </el-form-item>
                <el-form-item label="备份周期">
                  <el-select v-model="value" class="m-2" placeholder="选择备份周期">
                    <el-option
                        v-for="item in options"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value"
                    />
                  </el-select>
                </el-form-item>

                <el-form-item label="手动备份">
                  <el-button>点击备份</el-button>
                </el-form-item>

                <el-form-item label="还原备份">
                  <el-button>点击还原</el-button>
                </el-form-item>

                <el-form-item label="初始化">
                  <el-button type="danger">重置全站</el-button>
                </el-form-item>

              </el-form>
              <div style="margin-left: 20px;margin-top: 30px">
                <el-button type="primary" plain>保存配置</el-button>
                <el-button type="primary" plain>恢复初始配置</el-button>
              </div>
            </div>

          </el-main>
        </el-container>
        <el-divider/>
      </div>
      <!--管理员-->
      <div ref="targetComponent10">
        <el-container class="con1">
          <el-aside width="160px">
            <el-text style="font-size: 16px;font-weight: bold">管理员</el-text>

          </el-aside>
          <el-main>
            <div style="margin-bottom: 10px">
              <el-form label-width="100px" style="max-width: 460px" label-position="right">

                <el-form-item label="用户名">
                  <el-input v-model="formData.name" placeholder="修改用户名"/>
                </el-form-item>

                <el-form-item label="密码">
                  <el-input v-model="formData.name" placeholder="修改密码"/>
                </el-form-item>

                <el-form-item label="密保问题">
                  <el-input v-model="formData.name" placeholder="密保问题"/>
                </el-form-item>

                <el-form-item label="密保答案">
                  <el-input v-model="formData.name" placeholder="密保答案"/>
                </el-form-item>

                <el-form-item label="邮箱">
                  <el-input v-model="formData.name" placeholder="邮箱"/>
                </el-form-item>

                <el-form-item label="电话">
                  <el-input v-model="formData.name" placeholder="电话"/>
                </el-form-item>
              </el-form>
              <div style="margin-left: 20px;margin-top: 30px">
                <el-button type="primary" plain>保存设置</el-button>
              </div>
            </div>

          </el-main>
        </el-container>

      </div>
    </el-scrollbar>
  </div>


</template>

<script>
import {CirclePlus, Edit, Plus} from '@element-plus/icons-vue'

export default {
  data() {
    return {
      formData: {}
    }
  },
  components: {
    Plus,
    Edit,
    CirclePlus
  },
  methods: {
    scrollToComponent(index) {
      const scrollbar = this.$refs.scrollbarRef;
      const target = [this.$refs.targetComponent1,
        this.$refs.targetComponent2,
        this.$refs.targetComponent3,
        this.$refs.targetComponent4,
        this.$refs.targetComponent5,
        this.$refs.targetComponent6,
        this.$refs.targetComponent7,
        this.$refs.targetComponent8,
        this.$refs.targetComponent9,
        this.$refs.targetComponent10,
      ];

      if (scrollbar && target) {
        const offset = target[index].offsetTop;
        scrollbar.scrollTo({top: offset, behavior: 'smooth'});
      }
    }
  }
}
</script>


<style scoped>
.div1 {
  margin-bottom: 15px;
}

.text1 {
  font-size: 25px;
  font-weight: bolder;
  margin-left: 5px;
}

.link1 {
  font-size: 16px;
  font-weight: bold;
  margin-right: 10px;
}

.con1 {
  margin-bottom: 30px;
  margin-left: 5px;
}

</style>