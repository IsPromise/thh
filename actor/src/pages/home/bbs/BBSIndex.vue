<script setup lang="ts">
import {NButton, NH2, NCard, NLayout, NLayoutContent, NLayoutSider, NLayoutHeader, NMenu, useMessage} from "naive-ui";
import TwitterRoutes from "@/pages/manager/twitter/TwitterRoutes";
import {RouterLink} from "vue-router";
import {h, ref} from "vue";
import {getQueueLenApi, remoteService, runTSpiderMaster} from "@/service/remote.js";

const message = useMessage()


</script>

<template>
    <n-layout>
        <n-layout has-sider sider-placement="right"
                  class="content"
                  :style="{flex: 1, maxWidth: '1400px', margin: '0 auto', padding: '24px'}">
            <n-layout-content content-style="padding: 24px;">
                <router-view></router-view>
            </n-layout-content>
            <n-layout-sider
                    :width="360"
                    content-style="padding: 24px;"
                    bordered
            >
                <n-card>
                    <n-h2>海淀桥</n-h2>
                </n-card>

                <n-card>
                    <n-h2>海淀桥</n-h2>
                </n-card>

                <n-card>
                    <n-h2>海淀桥</n-h2>
                </n-card>
            </n-layout-sider>
        </n-layout>
    </n-layout>
</template>
<style>
@media (max-width: 1200px) {
    .content {
        max-width: 100%;
    }
}
</style>