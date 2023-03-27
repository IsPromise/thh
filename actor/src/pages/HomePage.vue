<script setup>
import {
    NGrid,
    NGi,
    NCard,
    NIcon,
    NInput,
    NLayout,
    NLayoutFooter,
    NLayoutContent,
    NLayoutHeader,
    NMenu,
    NAvatar
} from 'naive-ui';
import {useIsMobile, useIsSmallDesktop, useIsTablet} from '@/utils/composables';

import {
    FishOutline,
    FitnessOutline,
    InfiniteOutline,
    FastFoodOutline as CashIcon,
    GameController
} from '@vicons/ionicons5'
import {h, ref} from "vue";
import {RouterLink} from "vue-router";

function renderIcon(icon) {
    return () => h(NIcon, null, {default: () => h(icon)});
}

const isMobileRef = useIsMobile()
const isTabletRef = useIsTablet()
const isSmallDesktop = useIsSmallDesktop()
const isMobile = isMobileRef
const activeKey = ref('home')
const menuOptions = [
    {
        label: () => h(RouterLink, {to: {path: "/home/index",}},
            {
                default: () => "index"
            }
        ),
        key: "index",
        icon: renderIcon(FishOutline),
        children: null,
    }, {
        label: () => h(RouterLink, {to: {path: "/home/bbs",}},
            {
                default: () => "bbs"
            }
        ),
        key: "bbs",
        icon: renderIcon(InfiniteOutline),
        children: null,
    }, {
        label: () => h(RouterLink, {to: {path: "/home/about",}},
            {
                default: () => "about"
            }
        ),
        key: "about",
        icon: renderIcon(CashIcon),
        children: null,
    }
]
</script>
<template>
    <n-layout position="absolute">

        <n-layout-header class="l-header" position="absolute"
                         bordered
        >
            <n-grid cols="24" item-responsive>
                <n-gi span="8">
                    <n-menu v-model:value="activeKey" mode="horizontal" :options="menuOptions"/>
                </n-gi>
                <n-gi span="0 800:8">
                    <!--                    <n-input type="text">asdasd</n-input>-->
                </n-gi>
                <n-gi span="0 800:8">
                    <n-gi style="float:right">
                        <n-icon size="40" color="#0e7a0d">
                            <game-controller/>
                        </n-icon>
                        <n-avatar
                                round
                                size="medium"
                                src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg"
                        />
                        <n-avatar
                                round
                                size="medium"
                                src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg"
                        >
                            旺财
                        </n-avatar>
                    </n-gi>

                </n-gi>
            </n-grid>


        </n-layout-header>

        <n-layout-content position="absolute" style="top: 64px;">
            <router-view></router-view>
            <n-card>
              有花堪折直须折,莫待无花空折枝
            </n-card>
        </n-layout-content>
    </n-layout>
</template>

<style>
.l-header {
    padding: 12px 12px;
}
</style>