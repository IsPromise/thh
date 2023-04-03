import {h, resolveComponent} from "vue";
import {LogoTwitter, LogoWebComponent} from '@vicons/ionicons5'

import sun from "@/pages/HomePage.vue";
import moon from "@/pages/AllManager.vue";
import twitterTool from "@/pages/manager/twitter/TwitterRoutes.js"

let about = () => import("@/pages/home/AboutPage.vue")
let gridPage = () => import("@/pages/manager/GridPageDemo.vue")
let allTool = () => import("@/pages/manager/AllTool.vue")
let traefikManager = () => import("@/pages/manager/TraefikRouterManager.vue")
let imDemo = () => import("@/pages/manager/ImDemo.vue")
let voiceChat = () => import("@/pages/manager/VoiceChat.vue")
// let markdown = ()=>  import("@/pages/manager/MarkdownDemo.vue")
let index = () => import("@/pages/home/IndexPage.vue")
let twitterManager = () => import("@/pages/manager/TwitterManager.vue")
let sysInfo = () => import("@/pages/manager/SysInfo.vue")
let login = () => import("@/pages/Login.vue")


let bbs = () => import("@/pages/home/bbs/BBSIndex.vue")
let bbsPage = () => import("@/pages/home/bbs/BBSPage.vue")
let articlesPage = () => import("@/pages/home/bbs/ArticlesPage.vue")
let articlesEdit = () => import("@/pages/home/bbs/ArticlesEdit.vue")


export default [
    {
        path: '/:catchAll(.*)*', name: '', redirect: '/manager/'
    },
    {
        path: '/login', component: login
    },
    {
        path: '/home', component: sun, children: [
            {name: '', path: '', redirect: '/home/index'},
            {name: 'index', path: 'index', component: index},
            {
                path: 'bbs', component: bbs, children: [
                    {name: '', path: '', redirect: '/home/bbs/bbs'},
                    {name: 'bbs', path: 'bbs', component: bbsPage},
                    {name: 'articlesPage', path: 'articlesPage', component: articlesPage},
                    {name: 'articlesEdit', path: 'articlesEdit', component: articlesEdit},
                ]
            },
            {name: 'about', path: 'about', component: about},

        ]
    },
    {
        belongMenu: true,
        path: '/manager', component: moon, children: [
            {showName: '', path: '', component: allTool, belongMenu: false},
            {showName: 'all tool', path: 'allTool', component: allTool, belongMenu: true},
            {showName: 'Im', path: 'im', component: imDemo, belongMenu: true},
            {showName: 'voiceChat', path: 'voiceChat', component: voiceChat, belongMenu: true},
            // {showName: 'markdown', path: 'markdown', component: markdown, belongMenu: true},
            {showName: 'grid demo', path: 'gridPage', component: gridPage, belongMenu: true},
            {showName: 'sysInfo', path: 'sysInfo', component: sysInfo, belongMenu: true},
            {
                showName: '网关管理',
                path: 'traefikManager',
                icon: LogoWebComponent,
                component: traefikManager,
                belongMenu: true
            },
            {
                showName: 'twitterManager',
                path: 'twitterManager',
                component: twitterManager,
                belongMenu: true,
                children: twitterTool.map(item => {
                    return {showName: item.showName, path: item.path, component: item.component, belongMenu: false}
                })
            },
            {
                name: 'twitter',
                path: 'twitter',
                icon: LogoTwitter,
                belongMenu: true,
                component: {render: () => h(resolveComponent("router-view"))},
                children: twitterTool
            }
        ]
    },
]