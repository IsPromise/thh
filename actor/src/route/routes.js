import {h, resolveComponent} from "vue";
import {
    AlbumsOutline,
    ChatbubblesOutline,
    ConstructOutline,
    GitCompareOutline,
    GridOutline,
    HappyOutline,
    LogoTwitter,
    LogoWebComponent,
} from '@vicons/ionicons5'

import sun from "@/pages/HomePage.vue";
import moon from "@/pages/Manager.vue";
import twitterTool from "@/pages/manager/twitter/TwitterRoutes.js"

let about = () => import("@/pages/home/AboutPage.vue")
let gridPage = () => import("@/pages/manager/GridPageDemo.vue")
let allTool = () => import("@/pages/manager/AllTool.vue")
let todoList = () => import("@/pages/manager/TodoList.vue")
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
            {showName: 'All Tool', path: 'allTool', component: allTool, icon: ConstructOutline, belongMenu: true},
            {showName: 'Git Status', path: 'mainTool', component: mainTool, icon: GitCompareOutline, belongMenu: true},
            {showName: 'Todo List', path: 'todoList', component: todoList, icon: AlbumsOutline, belongMenu: true},
            {showName: 'Im', path: 'im', component: imDemo, icon: ChatbubblesOutline, belongMenu: true},
            {showName: 'Grid Demo', path: 'gridPage', component: gridPage, icon: GridOutline, belongMenu: true},
            {
                showName: '网关管理',
                path: 'TraefikManager',
                icon: LogoWebComponent,
                component: traefikManager,
                belongMenu: true
            },
            {
                showName: 'TwitterTool',
                path: 'twitter',
                icon: LogoTwitter,
                component: twitterManager,
                belongMenu: true,
                children: twitterTool.map(item => {
                    return {showName: item.showName, path: item.path, component: item.component, belongMenu: false}
                })
            },
            {
                name: 'About',
                path: 'about',
                belongMenu: true,
                icon: HappyOutline,
                component: {render: () => h(resolveComponent("router-view"))},
                children: [
                    {showName: 'about', path: 'about', component: managerAbout, belongMenu: true},
                    {showName: 'info', path: 'info', component: info, belongMenu: true},
                ]
            },
            {showName: 'SysInfo', path: 'sysInfo', icon: LogoWebComponent, component: sysInfo, belongMenu: true},
        ]
    },
]