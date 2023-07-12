import axios from "axios"
import {useUserStore} from "@/modules/user";
import {createDiscreteApi,} from "naive-ui";

const {message} = createDiscreteApi(
    ["message"],
);


const instanceAxios = axios.create({
    baseURL: import.meta.env.VITE_DEV_API_HOST,
    timeout: 10 * 1000,
    headers: {}
})

const userStore = useUserStore()
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
        userStore.updateToken(response.headers['new-token'])
    }
    const res = response.data
    if (res === undefined) {
        return response
    }
    if (res.code === fail) {
        message.error(res.msg ? res.msg : "响应异常")
    }
    return response
}, error => {
    // 处理错误
    if (error.response) {
        console.log('请求失败，HTTP 状态码：', error.response.status);
        console.log('错误信息：', error.response.data);
        const res = error.response.data
        if (res !== undefined && res.code === fail) {
            message.error(res.msg ? res.msg : "响应异常")
        }
    } else if (error.request) {
        console.log('请求发送失败：', error.request);
    } else {
        console.log('请求失败：', error.message);
    }
    return Promise.reject(error);
})

export function getUserInfo() {
    return instanceAxios.get("get-user-info-v4")
}

export function login(username, password) {
    return instanceAxios.post("/login", {
        username: username,
        password: password
    })
}

export function reg(email, username, password) {
    return instanceAxios.post("/reg", {
        email: email,
        username: username,
        password: password
    })
}

export function getTList(searchList) {
    return instanceAxios.post('twitter/get-mix-list', {
        SearchList: searchList
    })
}

export function getTwitterUserList(page = 1, pageSize = 10, search = "") {
    return instanceAxios.post('twitter/get-twitter-user-list', {
        page: page,
        pageSize: pageSize,
        search: search
    })
}

export function getTwitterTweetList(page = 1, pageSize = 10, search = "", useUserFilter = false) {
    return instanceAxios.post('twitter/get-twitter-tweet-list', {
        page: page,
        pageSize: pageSize,
        search: search,
        useUserFilter: useUserFilter
    })
}


export function runTSpiderMaster() {
    return instanceAxios.get('twitter/run-spider-twitter-master')
}

export function getQueueLenApi() {
    return instanceAxios.get('twitter/get-queue-len')
}

export function getTSpiderHis(page = 1, pageSize = 10) {
    return instanceAxios.post('twitter/get-spider-twitter-his', {
        page: page,
        pageSize: pageSize,
    })
}

export function setFilterUser(screenName = "") {
    return instanceAxios.post('twitter/set-filter-user', {
        screenName: screenName
    })
}

export function deleteFilterUser(screenName = "") {
    return instanceAxios.post('twitter/delete-filter-user', {
        screenName: screenName
    })
}

export function getFilterUser() {
    return instanceAxios.get('twitter/get-filter-user', {})
}


export function getGitStatus() {
    return instanceAxios.get('git-status-list')
}


export function wsInfo() {
    return instanceAxios.get('ws-info')
}

export function getSysInfo() {
    return instanceAxios.get("/sys-info")
}

export function getTodoStatusList(
) {
    return instanceAxios.get("todo-task/status-list",)
}

export function createTodoTaskList(
    taskName = "",
    description = "",
    deadline = "",
    weight = ""
) {
    return instanceAxios.post("todo-task/create", {
        taskName: taskName,
        description: description,
        deadline: deadline,
        weight: weight,
    })
}

export function updateTodoTaskList(
    taskId = "",
    taskName = "",
    description = "",
    status = "",
    deadline = "",
    weight = "",
    paused = ""
) {
    return instanceAxios.post("todo-task/update", {
        taskId: taskId,
        taskName: taskName,
        description: description,
        status: status,
        deadline: deadline,
        weight: weight,
        paused: paused,
    })
}

export function getTodoTaskList() {
    return instanceAxios.get("todo-task/list")
}