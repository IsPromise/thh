<script setup>
import {h, defineComponent, reactive, ref, onMounted} from "vue";
import {
  NTag,
  NSelect,
  NDataTable,
  NButton,
  NDatePicker,
  NInputNumber,
  useMessage,
  NForm,
  NSwitch,
  NFormItem,
  NInput,
  NSpace,
  NText
} from "naive-ui";
import {createTodoTaskList, getTodoStatusList, getTodoTaskList, updateTodoTaskList} from "@/service/remote";


const message = useMessage();
const size = ref("medium")
const data = ref([]);
let pagination = false
let createData = ref([])
let todoStatusList = ref([
  {id: '0', value: '创建'}]
)
let todoStatusMap = ref();
let searchParams = ref({
  status: []
})
let rules = ref({
      taskName: {
        required: true,
        message: "输入任务名",
        trigger: ["input"]
      }
    }
)
let submitData = reactive({
  taskName: "",
  description: "",
  deadline: "2023-02-03 02:03:04",
  weight: 0,
  paused: 0
})

let columns = [
  {align: "center", title: "id", key: "taskId"},
  {align: "center", title: "任务", key: "taskName"},
  {
    align: "center", title: "详情", key: "taskDescription",
    render(row) {
      if (row.isEditing) {
        return h(NInput, {
          value: row.taskDescription,
          onUpdateValue: val => {
            row.taskDescription = val;
          }
        });
      } else {
        return h(NText, {}, {default: () => row.taskDescription});
      }
    }
  },
  {
    align: "center", title: "状态", key: "status", render(row) {
      if (row.isEditing) {
        return h(NSelect, {
          value: row.status,
          options: todoStatusList.value,
          valueField: "id",
          labelField: "value",
          onUpdateValue: val => {
            row.status = val;
          }
        });
      } else {
        return h(NTag, {}, {default: () => todoStatusMap.value.get(row.status)});
      }
    }
  },
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
      return h(NSpace,
          {vertical: true, align: "center"},
          () => showList)
    }
  },
  {align: "center", title: "权重", key: "weight"},
  // {
  //   align: "center", title: "暂停", key: "paused", slot: 'paused', render(row) {
  //     return h(
  //         NSwitch,
  //         {
  //           checkedValue: 1,
  //           uncheckedValue: 0,
  //           value: row.paused,
  //           onUpdateValue: item => {
  //             row.paused = item
  //           }
  //         },
  //     );
  //   }
  // },
  {
    align: "center", title: "操作", key: "opt", render(row) {
      return [h(
          NButton,
          {strong: true, tertiary: true, size: "small", onClick: () => updateTask(row)},
          {default: () => "更新"}
      ),
        h(
            NButton,
            {strong: true, tertiary: true, size: "small", onClick: () => row.isEditing = !row.isEditing},
            {default: () => "编辑"}
        ), h(
            NButton,
            {strong: true, tertiary: true, size: "small", onClick: () => deleteTodoTask(row)},
            {default: () => "删除"}
        )
      ];
    }
  }
];


async function updateTask(row) {
  await updateTodoTaskList(
      row.taskId,
      row.taskName,
      row.taskDescription,
      row.status,
      row.deadline,
      row.weight,
      row.paused,
  )
  row.isEditing = false
  flashTodoTaskList()
}

function createTodoTask() {
  let result = createTodoTaskList(
      submitData.taskName,
      submitData.description,
      submitData.deadline,
      submitData.weight,
  ).then(r => {
    submitData.taskName = ""
    submitData.description = ""
    submitData.deadline = "2023-02-03 02:03:04"
    submitData.weight = 0
    submitData.paused = 0
    f5()
  })

}

function renderLabel(option) {
  return [
    option.value,
  ];
}

async function f5() {
  let todoTaskReq = await getTodoStatusList()
  todoStatusList.value = todoTaskReq.data.result
  todoStatusMap.value = todoStatusList.value.reduce((acc, obj) => {
    acc.set(obj.id, obj.value);
    return acc;
  }, new Map())
  flashTodoTaskList()
}

function flashTodoTaskList() {
  getTodoTaskList().then(result => {
    data.value = result.data.result.map(item => ({...item, isEditing: false}));
  });
}

onMounted(() => {
  f5()
})


</script>

<template>
  <n-form>
    <n-form-item>
      <n-select
          :options="todoStatusList"
          multiple
          :render-label="renderLabel"
          value-field="id"
          label-field="value"
      >
      </n-select>

      <n-button attr-type="button" @click="f5">
        刷新
      </n-button>
    </n-form-item>
  </n-form>
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
    </n-form-item>
  </n-form>
  <n-data-table
      :columns="columns"
      :data="data"
      :pagination="pagination"
      :bordered="false"
  />
</template>