import {h, resolveComponent} from "vue";
import {LogoTwitter, LogoWebComponent} from '@vicons/ionicons5'

import sun from "@/pages/HomePage.vue";
import moon from "@/pages/AllManager.vue";
import twitterTool from "@/pages/manager/twitter/TwitterRoutes.js"

let home = () => import("@/pages/home/MainPage.vue")
let about = () => import("@/pages/home/AboutPage.vue")
let bbs = () => import("@/pages/home/BBSPage.vue")
let articlesPage = () => import("@/pages/home/ArticlesPage.vue")
let gridPage = () => import("@/pages/manager/GridPageDemo.vue")
let allTool = () => import("@/pages/manager/AllTool.vue")
let traefikManager = () => import("@/pages/manager/TraefikRouterManager.vue")
let imDemo = () => import("@/pages/manager/ImDemo.vue")
// let markdown = ()=>  import("@/pages/manager/MarkdownDemo.vue")
let index = () => import("@/pages/home/IndexPage.vue")
let twitterManager = () => import("@/pages/manager/TwitterManager.vue")
let sysInfo = () => import("@/pages/manager/SysInfo.vue")
let login = () => import("@/pages/Login.vue")


export default [
    {
        path: '/:catchAll(.*)*', name: '', redirect: '/manager/'
    },
    {
        path: '/login', component: login
    },
    {
        path: '/home', component: sun, children: [
            {name: '', path: '', component: home},
            {name: 'index', path: 'index', component: index},
            {name: 'home', path: 'home', component: home},
            {name: 'about', path: 'about', component: about},
            {name: 'bbs', path: 'bbs', component: bbs},
            {name: 'articlesPage', path: 'articlesPage', component: articlesPage},
        ]
    },
    {
        belongMenu: true,
        path: '/manager', component: moon, children: [
            {showName: '', path: '', component: allTool, belongMenu: false},
            {showName: 'all tool', path: 'allTool', component: allTool, belongMenu: true},
            {showName: 'Im', path: 'im', component: imDemo, belongMenu: true},
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