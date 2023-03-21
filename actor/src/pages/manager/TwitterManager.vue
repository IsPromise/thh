<script setup lang="ts">
import {NButton, NLayout, NLayoutContent, NLayoutHeader, NMenu, useMessage} from "naive-ui";
import TwitterRoutes from "@/pages/manager/twitter/TwitterRoutes";
import {RouterLink} from "vue-router";
import {h, ref} from "vue";
import {remoteService} from "@/service/remote";

const message = useMessage()
const menuOptions = TwitterRoutes.filter(item => {
  return item.showName !== ""
}).map(item => {
  return {
    label: () => h(RouterLink, {to: {path: '/manager/twitterManager/' + item.path,}},
        {default: () => item.showName}
    ),
    key: item.path,
    children: null,
  }
})
const activeKey = ref<string>("")

function newSpider(e) {
  remoteService.runTSpiderMaster().then(r => {
    message.success(r.data.data.message);
  }).catch(e => {
    console.log(e)
    message.success("error");
  })
}

function getQueueLen(e) {
  remoteService.getQueueLen().then(r => {
    message.success(r.data.data.message);
  }).catch(e => {
    console.log(e)
    message.success("error");
  })
}

</script>

<template>
  <n-layout position="absolute">
    <n-layout-header style="height: 64px; padding: 18px;" bordered
                     position="absolute"
    >
      <n-menu v-model:value="activeKey" mode="horizontal" :options="menuOptions"/>
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