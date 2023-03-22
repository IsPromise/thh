<script setup>
import {NButton, NInput, NLayout, NLayoutContent, NLayoutFooter} from "naive-ui";
import {onMounted, onUnmounted, ref} from "vue";

let ws;
const myMessage = ref("")


function open() {
    if (ws) {
        return;
    }
    let wsLink = 'ws://' + document.domain + import.meta.env.VITE_WS_HOST
    ws = new WebSocket(wsLink);
    ws.onopen = function (evt) {
        console.log("连接websocket");
    }
    ws.onclose = function (evt) {
        console.log("CLOSE");
        ws = null;
    }
    ws.onmessage = function (evt) {
        let msg = JSON.parse(evt.data)
        console.log("收到消息: " + msg.message);
        messageList.value.push(msg.message);
        contentRef.value.scrollTo({top: 99999, behavior: 'smooth'})
    }
    ws.onerror = function (evt) {
        console.log("ERROR: " + evt.data);
    }
}

function send() {
    if (!ws) {
        return;
    }
    console.log("发送消息: " + new Date() + myMessage.value);
    ws.send(myMessage.value);
}

function close() {
    if (!ws) {
        return false;
    }
    ws.close();
    return true;
}

onMounted(() => {
    open()
})

onUnmounted(() => {
    close()
})

const messageList = ref([])
const contentRef = ref(null)
</script>

<template>
    <n-layout position="absolute">
        <n-layout-content
                ref="contentRef"
                position="absolute"
                style="bottom: 256px;"
                :native-scrollbar="false"
                id="im"
        >
            <p v-for="message in messageList"> {{ message }}</p>
        </n-layout-content>
        <n-layout-footer position="absolute" style="height: 256px;">
            <n-input v-model:value="myMessage" type="text" placeholder="基本的 Input"/>
            <n-button @click="send">发送消息</n-button>
        </n-layout-footer>
    </n-layout>


</template>