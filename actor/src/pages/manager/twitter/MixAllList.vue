<script setup>
import {h, ref} from 'vue'
import {NButton, NDataTable, NDynamicTags, NSpace, useMessage} from 'naive-ui'
import {getTList} from "@/service/remote";


const message = useMessage()
const dataRef = ref([])
const formRef = ref([])


const Search = (searchList) => {
  message.info("开始")
  getTList(searchList).then(r => {
    if (r.data.result !== undefined && r.data.result) {
      dataRef.value = r.data.result
      message.success("成功")
    } else {
      message.warning("无数据")
    }

  }).catch((error) => {
    console.log(error)
    message.error("失败")
  })
}

const columns = [
  {
    title: 'ScreenName',
    key: 'ScreenName',
    width: "120px"
  },
  {
    title: 'Name',
    key: 'Name',
    width: "120px"
  },
  {
    title: 'CreateTime',
    key: 'CreateTime',
    width: "120px"
  },
  {
    title: 'Desc',
    key: 'Desc',
    width: "360px"
  },
  {
    title: 'Url',
    key: 'Url',
    fixed: "right",
    width: "120px",
    render(row) {
      return h(
          NButton,
          {
            size: 'small',
            onClick: () => {
              window.open(row.Url)
            }
          },
          {default: () => 'open'}
      )
    }
  },
]
const pagination = {
  pageSize: 10
}

const value = ref(["特朗普"])
const handleValidateClick = function () {
  Search(value.value)
}
</script>
<template>

  <n-space style="padding: 20px">
    <n-dynamic-tags v-model:value="value" :max="10"/>
    <n-button attr-type="button" @click="handleValidateClick" size="small">
      搜索
    </n-button>
  </n-space>

  <n-data-table :columns="columns" :data="dataRef"/>
</template>
