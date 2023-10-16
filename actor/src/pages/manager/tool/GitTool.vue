<script setup>

import {getGitStatus} from "@/service/remote";
import {h, onMounted, ref} from "vue";
import {NButton, NDataTable, NSpace, NTag, useMessage,} from 'naive-ui'

let commentList = ref([]);
onMounted(getGitInfo)
const message = useMessage();

async function getGitInfo() {
  loadingRef.value = true
  let data = await getGitStatus()
  commentList.value = data.data.result
  message.success("刷新成功")
  loadingRef.value = false
}

let columns = [
  {
    align: "left", title: "Git Status", key: "taskId", render(row) {
      let dataList = []
      dataList.push(row.path)
      dataList.push(h(NTag,
          {type: row.hasCommits ? 'error' : 'info', size: 'small'},
          () => row.hasCommits ? 'unpushed' : 'pushed'))
      if (row.hasChanges) {
        dataList.push(h(NTag,
            {type: 'warning', size: 'small'},
            () => "hasChanges"))
      }
      return h(NSpace, () => dataList)

    }
  },
]
const loadingRef = ref(true)

</script>
<template>


  <n-space style="padding-bottom: 20px">
    <n-button @click="getGitInfo"> 刷新</n-button>
  </n-space>


  <n-data-table :columns="columns"
                :data="commentList"
                :pagination="false"
                :bordered="false"
                :loading="loadingRef"
  >
  </n-data-table>
</template>