<script setup>
import {NButton, NCard, NList, NListItem, NSpace, NTag, NThing} from 'naive-ui'
import {onMounted, ref} from "vue";
import {getArticlesApi, remoteService} from "@/service/remote";

const listData = ref([])

let maxId = 0

function getArticlesAction() {

    console.log(maxId)
    getArticlesApi(maxId).then(r => {
        console.log(r)
        let newList = r.data.data.list.map(function (item) {
            maxId = item.id
            return {
                id: item.id,
                title: "title1" + item.id,
                tag: ["tag1", "tag2"],
                desc: item.content,
                lastUpdate: "2022-12-28 01:01:01",

                body: item.content
            }
        })
        listData.value.push(...newList)
    }).catch(e => {
        console.log(e)
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
                            </n-space>
                        </template>
                    </n-thing>
                </router-link>
            </n-list-item>
            <n-list-item>
                <n-thing>
                    <n-button @click="more">
                        more
                    </n-button>
                </n-thing>
            </n-list-item>

        </n-list>
    </n-card>
</template>
<style>
a {
    text-decoration: none
}
</style>

