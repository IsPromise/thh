<script setup>

import {getGitStatus} from "@/service/remote";
import {onMounted, ref} from "vue";
import {
    NList,
    NListItem,
    NCard,
    NButton, useMessage, NTag, NSpace
} from 'naive-ui'

let commentList = ref([]);
onMounted(getGitInfo)
const message = useMessage();

async function getGitInfo() {
    let data = await getGitStatus()
    commentList.value = data.data.result
    message.success("刷新成功")
}

</script>
<template>
    <n-button @click="getGitInfo"> 刷新</n-button>
    <n-card>
        <n-list>
            <n-list-item v-for="item in commentList">

                <n-space>
                    {{ item.path }}:
                    <n-tag :bordered="false" :type="item.hasCommits? 'error':'info'" size="small">
                        {{ item.hasCommits ? 'unpushed' : 'pushed' }}
                    </n-tag>

                    <n-tag :bordered="false" v-if="item.hasChanges" type="warning" size="small">
                        hasChanges
                    </n-tag>
                </n-space>
            </n-list-item>
        </n-list>
    </n-card>
</template>