<script setup>
import {NCalendar, NCard, useMessage} from 'naive-ui';
import {ref} from 'vue'
import {addDays, isYesterday} from 'date-fns'

const message = useMessage()


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
</template>

<style>

</style>