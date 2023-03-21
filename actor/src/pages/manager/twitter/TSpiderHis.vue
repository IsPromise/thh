<script setup>
import {onMounted, reactive, ref} from "vue";
import {remoteService} from "@/service/remote";
import {NButton, NDataTable, NForm, NFormItem, useMessage} from "naive-ui";

const columns = ref([{title: 'CreateTime', key: 'CreateTime', width: "120px"}])

const pagination = reactive({
  page: 1,
  pageCount: 1,
  pageSize: 40,
  itemCount: 0,
  search: "",
  prefix({itemCount}) {
    return `Total is ${itemCount}.`
  }
})
const data = ref([])
const message = useMessage()
const searchPage = function (current) {
  remoteService.getTSpiderHis(current, pagination.pageSize = 10).then(r => {
    data.value = r.data.data.itemList
    columns.value = r.data.data.keyList.map(function (item) {
      let tData =  {
        title: item, key: item, ellipsis: true
      }
      if(item==="id"){
        tData.width= "120px"
      }
      return tData
    })
    pagination.page = current
    pagination.pageCount = parseInt(String(r.data.data.total / r.data.data.size))
    pagination.itemCount = r.data.data.total
    message.success("success");
  }).catch(e => {
    console.log(e)
    message.error("error");
  })
}

// onMounted(() => {
//   searchPage(1)
// })
const showModal = ref(false)

</script>
<template>
  <n-form
      ref="formRef"
      inline
      :label-width="80"
      :model="pagination"
  >

    <n-form-item>
      <n-button attr-type="button" @click="searchPage(0)">
        搜索
      </n-button>
    </n-form-item>
  </n-form>
  <n-data-table
      remote
      :columns="columns" :data="data" :pagination="pagination"
      @update:page="searchPage" flex-height :style="{ height: `600px` }" striped/>
  <!--  <n-modal v-model:show="showModal">-->
  <!--    <n-card style="width: 1000px;" title="详情" :bordered="false" size="huge">-->
  <!--      <n-list bordered>-->
  <!--        <n-list-item v-for="item in itemList">-->
  <!--          <n-thing :title="item.key" :description="item.value"-->
  <!--                   :style="{'word-wrap': 'break-word','word-break': 'break-all'}"/>-->
  <!--        </n-list-item>-->
  <!--      </n-list>-->
  <!--    </n-card>-->
  <!--  </n-modal>-->
</template>