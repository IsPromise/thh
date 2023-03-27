<script setup lang="ts">
import {NButton, NLayout, NLayoutContent, NLayoutHeader, NMenu, useMessage} from "naive-ui";
import TwitterRoutes from "@/pages/manager/twitter/TwitterRoutes";
import {RouterLink} from "vue-router";
import {h, ref} from "vue";
import {getQueueLenApi, remoteService, runTSpiderMaster} from "@/service/remote.js";

const message = useMessage()


</script>

<template>
    <router-view></router-view>
</template>