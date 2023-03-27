<script setup>
import {NList,NListItem,NCalendar, NGrid, NGridItem, NTimeline, NTimelineItem, useMessage} from 'naive-ui';
import moment from "moment"
import {ref} from "vue";
import {addDays, isYesterday} from "date-fns";


const message = useMessage()
let dayInfoList = [];
let nowT = moment()
let t = moment(moment().format("YYYY-01-01"))
for (let i = 1; i < 12; i++) {
  t.add(1, "months")
  let type = 'warning'
  let lineType = 'dashed'
  if (parseInt(t.format('M')) > parseInt(nowT.format('M'))) {
    type = 'success'
    lineType = 'default'
  }
  let timeInfo = t.format('YYYY-MM-DD')
  dayInfoList.push({
    title: timeInfo,
    time: timeInfo,
    // content: timeInfo,
    type: type,
    lineType: lineType
  })
}
dayInfoList.sort(function (item1, item2) {
  return item1.time > item2.time ? -1 : 1
})

dayInfoList.push({title: "start"})
dayInfoList.unshift({title: "end", type: "success"});

function success() {
  message.success(
      `还挺大`
  )
}


const value = ref(addDays(Date.now(), 1).valueOf())

function handleUpdateValue(_, {year, month, date}) {
  message.success(`${year}-${month}-${date}`)
}

function isDateDisabled(timestamp) {
  return isYesterday(timestamp);
}
</script>
<template>
  <n-list  class="thh-index" :show-divider="false">
    <n-list-item>
      <div class="greetings">
        <h1 class="green">thh</h1>
        <h3>
          You’ve successfully created a project with
          <a href="https://vitejs.dev/" target="_blank" rel="noopener">Vite</a> +
          <a href="https://vuejs.org/" target="_blank" rel="noopener">Vue 3</a>.
        </h3>
      </div>
    </n-list-item>
    <n-list-item >
      <n-timeline :size="'large'" >
        <n-timeline-item v-for="timeInfo in dayInfoList" :type="timeInfo.type"
                         :title="timeInfo.title"
                         :content="timeInfo.content"
                         :time="timeInfo.time"
                         :line-type="timeInfo.lineType"
        />
      </n-timeline>
    </n-list-item>
    <n-list-item >
      <n-calendar
          @update:value="handleUpdateValue"
          #="{ year, month, date }"
          v-model:value="value"
          :is-date-disabled="isDateDisabled"
      >
        {{ year }}-{{ month }}-{{ date }}
      </n-calendar>
    </n-list-item>
  </n-list>
</template>

<style scoped>
h1 {
  font-weight: 500;
  font-size: 2.6rem;
  top: -10px;
}

h3 {
  font-size: 1.2rem;
}

.greetings{
  max-width: 1280px;
  margin: 0 auto;
}
.greetings h1,
.greetings h3 {
  text-align: center;
}

/*@media (min-width: 1024px) {*/
/*  .greetings h1,*/
/*  .greetings h3 {*/
/*    text-align: left;*/
/*  }*/
/*}*/
.thh-index .n-list-item{
  padding-left: 20px;
  padding-right: 20px;
}

.green {
  text-decoration: none;
  color: hsla(160, 100%, 37%, 1);
  transition: 0.4s;
}

</style>