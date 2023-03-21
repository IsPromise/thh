
<script setup>
import {h, reactive, ref} from 'vue'
import {
  NButton,
  NDataTable,
  NDatePicker,
  NForm,
  NFormItem,
  NGi,
  NGrid,
  NInput,
  NModal,
  NSpace,
  NTag,
  useMessage
} from 'naive-ui'

const showModal = ref(false);
const formBtnLoading = ref(false);
const formParams = reactive({
  name: '',
  address: '',
  date: null,
});
const rules = {
  name: {
    required: true,
    trigger: ['blur', 'input'],
    message: '请输入名称',
  },
  address: {
    required: true,
    trigger: ['blur', 'input'],
    message: '请输入地址',
  },
  date: {
    type: 'number',
    required: true,
    trigger: ['blur', 'change'],
    message: '请选择日期',
  },
};

const createColumns = ({sendMail}) => {
  return [
    {
      title: '',
      key: 'isUse',
      render(row) {
        if (!row.isUse) {
          return
        }
        return h(NTag, {
              style: {
                marginRight: '6px'
              },
              type: 'error'
            },
            {
              default: () => "使用"
            })
      }
    }, {
      title: '选择',
      key: "isUse",
      render(row) {
        return h(NButton, {
          size: 'small',
          onClick: () => {
            console.log()
          }
        }, {
          default: () => "选择"
        })
      }
    },
    {
      title: 'host',
      key: 'host'
    },
    {
      title: 'port',
      key: 'port'
    },
    {
      title: 'version',
      key: 'version'
    },
    {
      title: '作者',
      key: 'author',
    },
    {
      title: '创建时间',
      key: 'createTime',
    },
    {
      title: 'Tags',
      key: 'tags',
      render(row) {
        return row.tags.map((tagKey) => {
          return h(
              NTag,
              {
                style: {
                  marginRight: '6px'
                },
                type: 'info'
              },
              {
                default: () => tagKey
              }
          )
        })
      }
    },
    {
      title: 'Action',
      key: 'actions',
      render(row) {
        if (!row.isLocal) {
          return
        }
        return [h(
            NButton,
            {
              size: 'small',
              onClick: () => sendMail(row)
            },
            {default: () => '编辑'}
        )
        ]
      }
    }
  ]
}

const createData = () => [
  {
    serviceId: 'serviceId',
    isUse: true,
    host: '127.0.0.1',
    port: '8080',
    version: '4.45.67.13',
    tags: ['xxx', 'xxxxx'],
    author: 'auther_content',
    createTime: 'xxx-xxx-xxx',
    isLocal: false
  },
  {
    serviceId: 'serviceId',
    isUse: false,
    host: '127.0.0.1',
    port: '8080',
    version: '4.45.67.13',
    tags: ['xxx', 'xxxxx'],
    author: 'auther_content',
    createTime: 'xxx-xxx-xxx',
    isLocal: true
  },
]

const message = useMessage()


const data = ref(createData())
const columns = createColumns({
  sendMail(rowData) {
    message.info('send mail to ' + rowData.name)
  }
})
const pagination = {
  pageSize: 100
}
const addTable = function () {
  showModal.value = true;
}
const confirmForm = function (e) {
  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      window['$message'].success('新建成功');
      setTimeout(() => {
        showModal.value = false;
        reloadTable();
      });
    } else {
      window['$message'].error('请填写完整信息');
    }
    formBtnLoading.value = false;
  });
}
</script>
<template>
  <n-grid>
    <n-gi :span="24">
      <n-button attr-type="button" @click="addTable">快速录入</n-button>
    </n-gi>
    <n-gi :span="24">
      <n-data-table :columns="columns" :data="data" :pagination="pagination"/>
    </n-gi>
  </n-grid>
  <n-modal v-model:show="showModal" :show-icon="false" preset="dialog" title="新建">
    <n-form
        :model="formParams"
        :rules="rules"
        ref="formRef"
        label-placement="left"
        :label-width="80"
        class="py-4"
    >
      <n-form-item label="名称" path="name">
        <n-input placeholder="请输入名称" v-model:value="formParams.name"/>
      </n-form-item>
      <n-form-item label="地址" path="address">
        <n-input type="textarea" placeholder="请输入地址" v-model:value="formParams.address"/>
      </n-form-item>
      <n-form-item label="日期" path="date">
        <n-date-picker type="datetime" placeholder="请选择日期" v-model:value="formParams.date"/>
      </n-form-item>
    </n-form>

    <template #action>
      <n-space>
        <n-button @click="() => (showModal = false)">取消</n-button>
        <n-button type="info" :loading="formBtnLoading" @click="confirmForm">确定</n-button>
      </n-space>
    </template>
  </n-modal>
</template>
