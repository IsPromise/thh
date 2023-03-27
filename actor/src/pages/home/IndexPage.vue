<script setup>
import {NCalendar, NCard, NGrid, NGridItem, NTimeline, NTimelineItem, useMessage} from 'naive-ui';
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
    <n-grid>
        <n-grid-item :span="24">
            <n-card style="margin:0 auto;right: 30px;left: 30px">
                <n-timeline :size="'large'">
                    <n-timeline-item v-for="timeInfo in dayInfoList" :type="timeInfo.type"
                                     :title="timeInfo.title"
                                     :content="timeInfo.content"
                                     :time="timeInfo.time"
                                     :line-type="timeInfo.lineType"
                    />
                </n-timeline>
            </n-card>
        </n-grid-item>
        <n-grid-item :span="24">
            <n-card style="margin:0 auto">
                <n-calendar
                        @update:value="handleUpdateValue"
                        #="{ year, month, date }"
                        v-model:value="value"
                        :is-date-disabled="isDateDisabled"
                >
                    {{ year }}-{{ month }}-{{ date }}
                </n-calendar>
            </n-card>
        </n-grid-item>
    </n-grid>
</template>

<style>

</style>