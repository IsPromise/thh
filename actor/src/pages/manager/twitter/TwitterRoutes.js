let twitterTable = () => import("@/pages/manager/twitter/MixAllList.vue")
let twitterUser = () => import("@/pages/manager/twitter/UserList.vue")
let twitterTweet = () => import("@/pages/manager/twitter/TweetList.vue")
let twitterHis = () => import('@/pages/manager/twitter/SpiderHis.vue')

export default [
    {showName: '', path: '', component: twitterTable, belongMenu: false},
    {showName: '综合', path: 'Table', component: twitterTable, belongMenu: true},
    {showName: 'User', path: 'User', component: twitterUser, belongMenu: true},
    {showName: 'Tweet', path: 'Tweet', component: twitterTweet, belongMenu: true},
    {showName: 'twitterHis', path: 'twitterHis', component: twitterHis, belongMenu: true},
]