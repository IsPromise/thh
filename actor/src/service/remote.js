import axios from "axios"
import {useUserStore} from "@/modules/user";
import {
    createDiscreteApi,
} from "naive-ui";

const {message} = createDiscreteApi(
    ["message"],
);


const instanceAxios = axios.create({
    baseURL: import.meta.env.VITE_DEV_API_HOST,
    timeout: 10 * 1000,
    headers: {}
})

const userStore = useUserStore()
const remoteService = {}
instanceAxios.interceptors.request.use(config => {
    config.headers = {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ' + userStore.token,
        ...config.headers
    }
    return config;
});

const success = 0
const fail = 1

instanceAxios.interceptors.response.use(response => {
    if (response.headers['new-token'] !== undefined) {
        userStore.token = response.headers['new-token']
    }
    const res = response.data
    if (res === undefined) {
        return response
    }
    if (res.code === fail) {
        message.error(res.msg ? res.msg : "响应异常")
    }
    return response
})

remoteService.getUserInfo = function () {
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
    return instanceAxios.post('bbs/get-articles', {
        maxId: maxId,
        pageSize: 10,
    })
}


remoteService.getArticlesDetail = function (id, maxCommentId) {
    return instanceAxios.post('bbs/get-articles-detail', {
        maxCommentId: maxCommentId,
        id: parseInt(id),
        pageSize: 10,
    })
}

remoteService.getSysInfo = function () {
    return instanceAxios.get("/sys-info")
}

export {remoteService}