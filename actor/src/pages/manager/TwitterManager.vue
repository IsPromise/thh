<script setup>
import {NButton, NDropdown, NLayout, NLayoutContent, NLayoutHeader, NMenu, useMessage} from "naive-ui";
import {useIsMobile} from "@/utils/composables";
import TwitterRoutes from "@/pages/manager/twitter/TwitterRoutes";
import {RouterLink} from "vue-router";
import {h, ref} from "vue";
import {getQueueLenApi, runTSpiderMaster} from "@/service/remote.js";

const message = useMessage()
const menuOptions = TwitterRoutes.filter(item => {
    return item.showName !== ""
}).map(item => {
    return {
        label: () => h(RouterLink, {to: {path: '/manager/twitter/' + item.path,}},
            {default: () => item.showName}
        ),
        key: item.path,
    }
})
const activeKey = ref("")

function newSpider(e) {
    runTSpiderMaster().then(r => {
        message.success(r.data.result.message);
    }).catch(e => {
        console.log(e)
        message.success("error");
    })
}

function getQueueLen(e) {
    getQueueLenApi().then(r => {
        message.success(r.data.result.message);
    }).catch(e => {
        console.log(e)
        message.success("error");
    })
}

const isMobileRef = useIsMobile()

</script>

<template>
    <n-layout position="absolute">
        <n-layout-header style="height: 64px; padding: 18px;" bordered
                         position="absolute"
        >

            <n-dropdown v-if="isMobileRef" trigger="hover" :options="menuOptions">
                <n-button>目录</n-button>
            </n-dropdown>
            <n-menu v-else v-model:value="activeKey" mode="horizontal" :options="menuOptions"/>

            <n-button attr-type="button" @click="newSpider">
                新的抓取
            </n-button>

            <n-button attr-type="button" @click="getQueueLen">
                当前队列长度
            </n-button>

        </n-layout-header>
        <n-layout-content position="absolute" style="top: 64px;">
            <router-view></router-view>
        </n-layout-content>
    </n-layout>
</template>