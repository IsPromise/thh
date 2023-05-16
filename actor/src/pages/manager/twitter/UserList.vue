<script setup>
import {h, onMounted, reactive, ref} from 'vue'
import {getTwitterUserList} from "@/service/remote";
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
    NSpace,
    NThing, NTime,
    useMessage
} from "naive-ui"


const showModal = ref(false)
const testInfoList = ref([{key: "", value: ""}])
const columnsRefNew = [
    {
        title: '时间', key: 'CreateTime', width: "80px",
        align: "center",
        titleAlign: "center",
        render(row) {
            let showList = [];
            let times = new Date(Date.parse(row.CreateTime));
            showList.push(h(NTime, {time: times, type: "relative"}))
            return h(NSpace, {
                vertical: true,
                align: "center"
            }, () => showList)
        }
    },
    {title: 'ScreenName', key: 'ScreenName', width: "120px", ellipsis: true},
    {title: 'Name', key: 'Name', width: "120px"},
    {title: 'Desc', key: 'Desc',},
    {
        title: 'info',
        key: 'info',
        width: "100px",
        fixed: "right",
        align: "center",
        titleAlign: "center",
        render(row) {
            return h(NSpace, {
                vertical: true,
                align: "center"
            }, () => [h(
                NButton,
                {
                    size: 'small',
                    onClick: () => {
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
                },
                {default: () => '详情展示'}
            ), h(
                NButton,
                {
                    type: "primary",
                    size: 'small',
                    onClick: () => {
                        window.open(row.Url)
                    }
                },
                {default: () => 'open'}
            )])

        }
    }
]

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
    getTwitterUserList(current, paginationReactive.pageSize, paginationReactive.search).then(r => {
        dataRef.value = r.data.result.itemList
        paginationReactive.page = current
        paginationReactive.pageCount = parseInt(String(r.data.result.total / r.data.result.size))
        paginationReactive.itemCount = r.data.result.total
        message.success("success");
    }).catch(e => {
        console.log(e)
        message.error("error");
    })
}

onMounted(() => {
    searchPage(1)
})

const message = useMessage()

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

</script>
<template>
    <n-form
            ref="formRef"
            inline
            :label-width="80"
            :model="paginationReactive"
            :rules="rules"
            size="medium"
            style="padding: 0 20px "
    >
        <n-form-item>
            <n-space>
                <n-input v-model:value="paginationReactive.search" placeholder="搜索内容"/>
                <n-button attr-type="button" @click="handleValidateClick">
                    搜索
                </n-button>
            </n-space>
        </n-form-item>
    </n-form>
    <n-data-table
            remote
            :columns="columnsRefNew" :data="dataRef" :pagination="paginationReactive"
            @update:page="searchPage" flex-height :style="{ height: `600px` }" striped/>
    <n-modal v-model:show="showModal">
        <n-card style="width: 1000px;" title="详情" :bordered="false" size="huge">
            <n-list bordered>
                <n-list-item v-for="item in testInfoList">
                    <n-thing :title="item.key" :description="item.value"
                             :style="{'word-wrap': 'break-word','word-break': 'break-all'}"/>
                </n-list-item>
            </n-list>
        </n-card>
    </n-modal>
</template>
