import axios from "axios"
import {useUserStore} from "@/modules/user";

const instanceAxios = axios.create({
    baseURL: import.meta.env.VITE_DEV_API_HOST,
    timeout: 10 * 1000,
    headers: {}
})

const remoteService = {}
instanceAxios.interceptors.request.use(config => {
    const userStore = useUserStore()
    config.headers = {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ' + userStore.token,
        ...config.headers
    }
    return config;
});

remoteService.getUserInfo = function (){
    return instanceAxios.get("get-user-info-v4")
}

remoteService.login = function (username, password) {
    return instanceAxios.post("/login", {
        username: username,
        password: password
    })
}

remoteService.reg = function (email, username, password) {
    return instanceAxios.post("/reg", {
        email: email,
        username: username,
        password: password
    })
}

remoteService.getTList = function (searchList) {
    return instanceAxios.post('t-list', {
        SearchList: searchList
    })
}
remoteService.getTwitterUserList = function (page = 1, pageSize = 10, search = "") {
    return instanceAxios.get('get-twitter-user-list', {
        params: {
            page: page,
            pageSize: pageSize,
            search: search
        }
    })
}

remoteService.getTwitterTweetList = function (page = 1, pageSize = 10, search = "") {
    return instanceAxios.get('get-twitter-tweet-list', {
        params: {
            page: page,
            pageSize: pageSize,
            search: search
        }
    })
}


remoteService.runTSpiderMaster = function () {
    return instanceAxios.get('run-tspider-master')
}

remoteService.getQueueLen = function () {
    return instanceAxios.get('get-queue-len')
}

remoteService.getTSpiderHis = function (page = 1, pageSize = 10) {
    return instanceAxios.get('get-tspider-his', {
        params: {
            page: page,
            pageSize: pageSize,
        }
    })
}

remoteService.getArticles = function (maxId) {
    return instanceAxios.post('get-articles', {
        maxId: maxId,
        pageSize: 10,
    })
}


remoteService.getArticlesDetail = function (id, maxCommentId) {
    return instanceAxios.post('get-articles-detail', {
        maxCommentId: maxCommentId,
        id: parseInt(id),
        pageSize: 10,
    })
}

remoteService.getSysInfo = function () {
    return instanceAxios.get("/sys-info")
}

export {remoteService}