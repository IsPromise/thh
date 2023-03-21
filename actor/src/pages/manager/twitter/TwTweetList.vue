

<script setup>
import {h, onMounted, reactive, ref} from 'vue'
import {remoteService} from "@/service/remote";
import {
  NButton,
  NCard,
  NDataTable,
  NForm,
  NFormItem,
  NInput,
  NList,
  NListItem,
  NModal,
  NThing,
  useMessage
} from "naive-ui"


const columnsList = ({action}) => {
  return [
    {title: 'CreateTime', key: 'CreateTime', width: "120px",},
    {title: 'ScreenName', key: 'ScreenName', width: "120px", ellipsis: true},
    {title: 'originScreenName', key: 'originScreenName', width: "120px", ellipsis: true},
    {title: 'Name', key: 'Name', width: "120px"},
    {title: 'Desc', key: 'Desc',},

    {
      title: 'Url',
      key: 'Url',
      width: "60px",
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
    {
      title: 'info',
      key: 'info',
      width: "100px",
      render(row) {
        return h(
            NButton,
            {
              size: 'small',
              onClick: () => {
                action(row)
              }
            },
            {default: () => '详情展示'}
        )
      }
    }
  ]
}


const showModal = ref(false)
const testInfoList = ref([{key: "", value: ""}])
const columnsRef = ref(columnsList({
  action(row) {
    showModal.value = true
    let tmpTestInfoList = []
    let notShowKey = ['conclusion', 'groupId']

    Object.keys(row).forEach(function (key) {
      if (notShowKey.indexOf(key) !== -1) {
        return
      }
      tmpTestInfoList.push({key: key, value: row[key].toString()});
    });

    testInfoList.value = tmpTestInfoList
  }
}))
const paginationReactive = reactive({
  page: 1,
  pageCount: 1,
  pageSize: 40,
  itemCount: 0,
  search: "",
  prefix({itemCount}) {
    return `Total is ${itemCount}.`
  }
})
const dataRef = ref([])
const formRef = ref(null);
const searchPage = function (current) {
  remoteService.getTwitterTweetList(current, paginationReactive.pageSize, paginationReactive.search).then(r => {
    dataRef.value = r.data.data.itemList
    paginationReactive.page = current
    paginationReactive.pageCount = parseInt(String(r.data.data.total / r.data.data.size))
    paginationReactive.itemCount = r.data.data.total
    message.success("success");
  }).catch(e => {
    console.log(e)
    message.success("error");
  })
}

onMounted(() => {
  searchPage(1)
})

const message = useMessage()


const size = ref("medium")
const rules = {
  search: {
    required: false,
    message: "输入搜索内容",
    trigger: ["input"]
  }
}

function handleValidateClick(e) {
  formRef.value?.validate((errors) => {
    if (!errors) {
      message.success("Valid");
      searchPage(1)
    } else {
      console.log(errors);
      message.error("Invalid");
    }
  });
}

function newSpider(e) {
  remoteService.runTSpiderMaster().then(r => {
    message.success(r.data.data.message);
  }).catch(e => {
    console.log(e)
    message.success("error");
  })
}

function getQueueLen(e) {
  remoteService.getQueueLen().then(r => {
    message.success(r.data.data.message);
  }).catch(e => {
    console.log(e)
    message.success("error");
  })
}

const itemList = testInfoList

const data = dataRef
const pagination = paginationReactive
const columns = columnsRef
const handlePageChange = searchPage

function showInfo() {
  this.showModal = true
}
</script>
<template>
  <n-form
      ref="formRef"
      inline
      :label-width="80"
      :model="pagination"
      :rules="rules"
      :size="size"
  >
    <n-form-item label="搜索内容" path="phone">
      <n-input v-model:value="pagination.search" placeholder="搜索内容"/>
    </n-form-item>
    <n-form-item>
      <n-button attr-type="button" @click="handleValidateClick">
        搜索
      </n-button>
      <n-button attr-type="button" @click="newSpider">
        新的抓取
      </n-button>

      <n-button attr-type="button" @click="getQueueLen">
        当前队列长度
      </n-button>
    </n-form-item>
  </n-form>
  <n-data-table
      remote
      :columns="columns" :data="data" :pagination="pagination"
      @update:page="handlePageChange" flex-height :style="{ height: `600px` }" striped/>
  <n-modal v-model:show="showModal">
    <n-card style="width: 1000px;" title="详情" :bordered="false" size="huge">
      <n-list bordered>
        <n-list-item v-for="item in itemList">
          <n-thing :title="item.key" :description="item.value"
                   :style="{'word-wrap': 'break-word','word-break': 'break-all'}"/>
        </n-list-item>
      </n-list>
    </n-card>
  </n-modal>
</template>