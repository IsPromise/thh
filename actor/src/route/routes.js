import {h, resolveComponent} from "vue";
import {LogoTwitter, LogoWebComponent} from '@vicons/ionicons5'

import sun from "@/pages/HomePage.vue";
import moon from "@/pages/Manager.vue";
import twitterTool from "@/pages/manager/twitter/TwitterRoutes.js"

let about = () => import("@/pages/home/AboutPage.vue")
let gridPage = () => import("@/pages/manager/GridPageDemo.vue")
let allTool = () => import("@/pages/manager/AllTool.vue")
let mainTool = () => import("@/pages/manager/MainTool.vue")
let traefikManager = () => import("@/pages/manager/TraefikRouterManager.vue")
let imDemo = () => import("@/pages/manager/ImDemo.vue")
let index = () => import("@/pages/home/IndexPage.vue")
let twitterManager = () => import("@/pages/manager/TwitterManager.vue")
let sysInfo = () => import("@/pages/manager/SysInfo.vue")
let login = () => import("@/pages/Login.vue")
let managerAbout = () => import("@/pages/manager/about/about.vue")
let info = () => import("@/pages/manager/about/info.vue")

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
            {name: 'about', path: 'about', component: about},

        ]
    },
    {
        belongMenu: true,
        path: '/manager', component: moon, children: [
            {showName: '', path: '', component: allTool, belongMenu: false},
            {showName: 'mainTool', path: 'mainTool', component: mainTool, belongMenu: true},
            {showName: 'all tool', path: 'allTool', component: allTool, belongMenu: true},
            {showName: 'Im', path: 'im', component: imDemo, belongMenu: true},
            {showName: 'grid demo', path: 'gridPage', component: gridPage, belongMenu: true},
            {
                showName: '网关管理',
                path: 'traefikManager',
                icon: LogoWebComponent,
                component: traefikManager,
                belongMenu: true
            },
            {
                showName: 'twitter',
                path: 'twitter',
                icon: LogoTwitter,
                component: twitterManager,
                belongMenu: true,
                children: twitterTool.map(item => {
                    return {showName: item.showName, path: item.path, component: item.component, belongMenu: false}
                })
            },
            {
                name: 'about',
                path: 'about',
                belongMenu: true,
                component: {render: () => h(resolveComponent("router-view"))},
                children: [
                    {showName: 'about', path: 'about', component: managerAbout, belongMenu: true},
                    {showName: 'info', path: 'info', component: info, belongMenu: true},
                ]
            },
            {showName: 'sysInfo', path: 'sysInfo', icon: LogoWebComponent, component: sysInfo, belongMenu: true},
        ]
    },
]