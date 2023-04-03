<script setup>
import {onMounted, onUnmounted} from "vue";
import {wsInfo} from "@/service/remote";

let ws;
let localStream;
let mediaRecorder;

function init() {
    // 获取本地媒体流
    navigator.mediaDevices.getUserMedia({audio: true})
        .then(async function (stream) {
            localStream = stream;

            // 绑定本地媒体流到音频标签
            let audioElement = document.getElementById("localAudio");
            audioElement.srcObject = localStream;


            // 连接 ws 服务端
            let wsInfoData = await wsInfo()
            let wsLink = 'ws://' + document.domain + ':' + wsInfoData.data.result.ws + '/ws-vc'
            ws = new WebSocket(wsLink);

            // 监听连接事件
            ws.onopen = function (event) {
                console.log("连接成功");

                // 开始录制音频
                mediaRecorder = new MediaRecorder(localStream);
                mediaRecorder.ondataavailable = function (event) {
                    ws.send(event.data);
                }
                mediaRecorder.start();
            };

            // 监听消息事件
            ws.onmessage = function (event) {
                let data = event.data;
                console.log("收到消息：" + data);

                // TODO: 处理接收到的消息
            };

            // 监听错误事件
            ws.onerror = function (event) {
                console.log("连接出错");
            };

            // 监听断开连接事件
            ws.onclose = function (event) {
                console.log("连接已关闭");
            };
        })
        .catch(function (error) {
            console.log("获取本地媒体流失败：" + error);
        });
}

// 停止录音，停止音频播放
function stop() {
    let audioElement = document.getElementById("localAudio");
    audioElement.pause();
    mediaRecorder.stop();
}


function close() {
    if (!ws) {
        return false;
    }
    ws.close();
    return true;
}

onMounted(() => {
    init()
})

onUnmounted(() => {
    close()
    stop()
})
</script>

<template>
    <h1>Voice Chat Room</h1>

  <!-- 消息列表 -->
    <div id="messageList"></div>

  <!-- 本地音频播放器 -->
    <audio id="localAudio" autoplay>biubiubiu</audio>
</template>