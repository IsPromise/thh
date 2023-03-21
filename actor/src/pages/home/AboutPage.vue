<script setup>
import {NCard, NCountdown, NGi, NGrid, NSpace} from 'naive-ui'
const date1 = new Date('2023/12/24')
const time1 = parseInt(date1-(new Date()))
function renderCountdown({  hours,
                           minutes,
                           seconds}){
  let day = parseInt(hours/24);
  let dayHours = hours%24
  return `至 2023/12/24 ${String(day).padStart(3, "0")} 天 ${String(dayHours).padStart(2, "0")} 小时 ${String(minutes).padStart(2, "0")} 分 ${String(seconds).padStart(2, "0")} 秒`;
}
</script>
<template>
  <n-card style="margin:0 auto">
    <n-space>
      <n-grid>
        <n-gi span="24">
          <p>青山不改绿水长流</p>
        </n-gi>

        <n-gi span="24">
          <n-countdown ref="countdown" :render="renderCountdown" :duration="time1" :active="true"/>
        </n-gi>
      </n-grid>
    </n-space>
  </n-card>
</template>


