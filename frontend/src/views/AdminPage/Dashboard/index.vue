<!--用于管理员管理-仪表盘-->
<script setup>
import {onMounted} from 'vue';


import card1 from './components/card1.vue';

//数据展示
import * as echarts from 'echarts';
// 店铺热度排行
onMounted(() => {
  const myChart = echarts.init(document.getElementById('p1'));
  myChart.setOption({
    title: {
      text: '店铺热度排行',
    },
    tooltip: {},
    xAxis: {
      data: ['1', '2', '3', '4', '5', '6'],
    },
    yAxis: {},
    series: [
      {
        name: '访问量',
        type: 'bar',
        data: [5, 10, 15, 20, 25, 30],
      },
    ],
  });
})

// 每周用户活跃量
onMounted(() => {
  const chart = echarts.init(document.getElementById('p2'));
  chart.setOption({
    title: {
      text: '每日用户活跃量',
    },
    xAxis: {
      type: 'category',
      data: ['123', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
    },
    yAxis: {
      type: 'value'
    },
    series: [
      {
        data: [150, 230, 224, 218, 135, 147, 260],
        type: 'line'
      }
    ]
  })
});


</script>

<template>
  <div>
    <!--  topNav-->
    <div style="margin-bottom: 24px">
      <el-text tag="b" size="large">仪表盘</el-text>
    </div>
    <el-scrollbar height="85vh">
      <el-space wrap style="width: 100%; margin-bottom: 24px" fill :fill-ratio="24" size="default">
        <card1/>
        <card1/>
        <card1/>
        <card1/>

        <el-card id="p1" style="width: 560px;height:360px; padding: 10px;max-width: 560px"></el-card>
        <el-card id="p2" style=" padding: 10px;width: 600px;height:400px;max-width: 600px"></el-card>

      </el-space>
    </el-scrollbar>
  </div>


</template>

<style scoped>

</style>