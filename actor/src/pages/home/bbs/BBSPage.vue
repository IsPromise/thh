<script setup>
import {NButton, NCard, NList, NListItem, NSpace, NTag, NThing} from 'naive-ui'
import {onMounted, ref} from "vue";
import {getArticlesPageApi} from "@/service/remote";

const listData = ref([])

let maxId = 0

function getArticlesAction() {
    getArticlesPageApi(maxId).then(r => {
        let newList = r.data.result.list.map(function (item) {
            return {
                id: item.id,
                title: item.title,
                tag: ["tag1", "tag2"],
                desc: item.content,
                lastUpdateTime: item.lastUpdateTime,
                body: item.content
            }
        })
        listData.value.push(...newList)
        maxId += 1
    })
}

onMounted(() => {
    maxId = 0
    getArticlesAction()
})

function more() {
    getArticlesAction()
}


</script>
<template>
    <n-card style="margin:0 auto">
        <n-list>
            <n-list-item v-for="item in listData">
                <router-link :to="{path:'articlesPage',query:{title:item.title,id:item.id}}">
                    <n-thing>
                        <template #description>
                            <n-space size="small" style="padding-top: 4px">
                                {{ item.title }}
                                <n-tag v-for="itemTag in item.tag" :bordered="false" type="info" size="small"
                                       v-text="itemTag">
                                </n-tag>
                                {{ item.lastUpdateTime }}
                            </n-space>
                        </template>
                    </n-thing>
                </router-link>
            </n-list-item>
            <n-list-item>
                <n-button @click="more">
                    more
                </n-button>
            </n-list-item>

        </n-list>
    </n-card>
</template>
<style>
a {
    text-decoration: none
}
</style>

