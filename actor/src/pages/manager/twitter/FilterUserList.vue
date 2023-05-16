<script setup>
import {h, onMounted, ref} from "vue";
import {NDataTable, NButton} from "naive-ui";
import {getFilterUser, deleteFilterUser} from "@/service/remote";

const columnsRefNew = [
    {title: 'ScreenName', key: 'screenName', width: "240px", ellipsis: true},
    {
        title: '解除', key: 'jiechu', render(row) {
            return h(NButton, {
                size: 'small', onClick: () => {
                    deleteFilterUser(row.screenName).then(r => {
                        getData()
                    })
                }
            }, "解除")
        }
    }
]
const dataRef = ref([])

onMounted(() => {
    getData()
})

function getData() {
    getFilterUser().then(r => {
        dataRef.value = r.data.result
    }).catch(e => {
        console.error(e)
    })
}
</script>
<template>
    <n-data-table :columns="columnsRefNew" :data="dataRef"/>
</template>