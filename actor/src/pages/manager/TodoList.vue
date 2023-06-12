<script setup>
import {h, defineComponent, ref, onMounted} from "vue";
import {
  NDataTable,
  NButton,
  NDatePicker,
  NInputNumber,
  useMessage,
  NForm,
  NSwitch,
  NFormItem,
  NInput,
  NSpace
} from "naive-ui";
import {createTodoTaskList, getTodoTaskList} from "@/service/remote";

let columns = [
  {align: "center", title: "id", key: "taskId"},
  {align: "center", title: "任务", key: "taskName"},
  {align: "center", title: "详情", key: "taskDescription"},
  {align: "center", title: "状态", key: "status"},
  {
    align: "center", title: "创建～截止", key: "createTime", render(row) {
      let showList = [
        h(
            'span',
            {},
            {default: () => row.deadline}
        ),
        h(
            'span',
            {},
            {default: () => row.createTime}
        )
      ]
      return h(NSpace, {
        vertical: true,
        align: "center"
      }, () => showList)
    }
  },
  {align: "center", title: "权重", key: "weight"},
  {
    align: "center", title: "暂停", key: "paused", render(row) {
      console.log(row.paused)
      let action = ref(row.paused)
      return h(
          NSwitch,
          {
            modelValue: row.paused,
            "onUpdate:value": item => {
              console.log(item)
              console.log(row)
            }
          },
      );
    }
  },
  {
    align: "center", title: "操作", key: "opt", render(row) {
      return h(
          NButton,
          {
            strong: true,
            tertiary: true,
            size: "small",
            onClick: () => console.log(row)
          },
          {default: () => "更新"}
      );
    }
  },
];


const data = ref([]);


const message = useMessage();


let pagination = false


let createData = ref([])
let rules = ref({
      taskName: {
        required: true,
        message: "输入任务名",
        trigger: ["input"]
      }
    }
)
const size = ref("medium")
const submitData = ref({
  taskName: "",
  description: "",
  deadline: ref("2023-02-03 02:03:04"),
  weight: 0,
  paused: ""
});

let submitInitData = {
  taskName: "",
  description: "",
  deadline: ref("2023-02-03 02:03:04"),
  weight: 0,
  paused: ""
}

async function createTodoTask() {
  let result = await createTodoTaskList(
      submitData.value.taskName,
      submitData.value.description,
      submitData.value.deadline,
      submitData.value.weight,
  )
  submitData.value = submitInitData
  f5()
}

async function f5() {
  let result = await getTodoTaskList();
  data.value = result.data.result;
}

onMounted(() => {
  f5()
})


</script>

<template>
  <n-form
      ref="formRef"
      inline
      :label-width="80"
      :model="createData"
      :rules="rules"
      :size="size"
  >
    <n-form-item label="任务名">
      <n-input v-model:value="submitData.taskName" placeholder="任务名"/>
    </n-form-item>

    <n-form-item label="任务描述">
      <n-input v-model:value="submitData.description" placeholder="任务描述"/>
    </n-form-item>

    <n-form-item label="截止日期">
      <n-date-picker v-model:formatted-value="submitData.deadline" type="datetime" value-format="yyyy-MM-dd HH:mm:ss"
                     clearable/>
    </n-form-item>

    <n-form-item label="权重">
      <n-input-number v-model:value="submitData.weight" clearable/>
    </n-form-item>

    <n-form-item>
      <n-button attr-type="button" @click="createTodoTask">
        创建
      </n-button>

      <n-button attr-type="button" @click="f5">
        刷新
      </n-button>
    </n-form-item>
  </n-form>
  <n-data-table
      :columns="columns"
      :data="data"
      :pagination="pagination"
      :bordered="false"
  />
</template>