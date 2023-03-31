<script setup>
import {h, ref} from 'vue'
import {NButton, NCard, NDataTable, NDynamicTags, NSpace, useMessage} from 'naive-ui'
import {getTList, remoteService} from "@/service/remote";

const createColumns = () => {
  return [
    {
      title: 'ScreenName',
      key: 'ScreenName'
    },
    {
      title: 'Name',
      key: 'Name'
    },
    {
      title: 'CreateTime',
      key: 'CreateTime'
    },
    {
      title: 'Desc',
      key: 'Desc'
    },
    {
      title: 'Url',
      key: 'Url',
      render(row) {
        return h(
            NButton,
            {
              size: 'small',
              onClick: () => {
                // console.log(row.Url)
                window.open(row.Url)
              }
            },
            {default: () => 'open'}
        )
      }
    },
  ]
}


const message = useMessage()
const dataRef = ref([])
const formRef = ref([])


const Search = (searchList) => {
  message.info("开始")
  getTList(searchList).then(r => {
    dataRef.value = r.data.result
    console.log(r.data)
    message.success("成功")
  }).catch((error) => {
    console.log(error)
    message.error("失败")
  })
}

const columns = createColumns()
const pagination = {
  pageSize: 10
}

const value = ref(["特朗普"])
const handleValidateClick = function () {
  Search(value.value)
}
</script>
<template>
  <n-card>
    <n-space>
      <n-dynamic-tags v-model:value="value" :max="10"/>
      <n-button attr-type="button" @click="handleValidateClick" size="small">
        搜索
      </n-button>
    </n-space>
  </n-card>
  <n-data-table :columns="columns" :data="dataRef"/>
</template>
