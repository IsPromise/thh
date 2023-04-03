<script setup>
import {wsInfo} from "@/service/remote";
import {onMounted} from "vue";

let websocket;

async function start() {
  let wsInfoData = await wsInfo()
  let wsLink = 'ws://' + document.domain + ':' + wsInfoData.data.result.ws + '/ws-vc'
  websocket = new WebSocket(wsLink);
  console.log("success")

  websocket.onmessage = function (event) {
    console.log(event.data)
    playAudio(event.data);
  }
}


function startRecording() {
  navigator.mediaDevices.getUserMedia({audio: true, video: false})
      .then(function (stream) {
        let mediaRecorder = new MediaRecorder(stream);

        mediaRecorder.start();

        let chunks = [];

        mediaRecorder.ondataavailable = function (e) {
          chunks.push(e.data);
        }

        mediaRecorder.onstop = function (e) {
          console.log('recording stopped');

          let blob = new Blob(chunks, {type: 'audio/ogg; codecs=opus'});

          let reader = new FileReader();
          reader.readAsDataURL(blob);
          reader.onloadend = function () {
            let base64data = reader.result;//.split(',')[1];
            console.log(base64data)
            playAudio(base64data)
            // 将音频数据发送到服务器
            console.log("发送")
            // websocket.send(base64data);
          }
        };

        // 添加停止录音的代码
        setTimeout(function() {
          mediaRecorder.stop();
        }, 3000); // 录制5秒钟

      })
      .catch(function (err) {
        console.log('getUserMedia error: ' + err);
      });
}

function playAudio(base64data) {
  let audioBlob = new Blob([atob(base64data)], {type: 'audio/ogg'});
  let audioUrl = URL.createObjectURL(audioBlob);

  let audioElem = document.getElementById('audio');
  audioElem.src = audioUrl;
  audioElem.play();
}

onMounted(() => {
  start();
})

// onUnmounted(() => {
//   close()
// })
</script>

<template>
  <button @click="startRecording()">Start Recording</button>
  <audio id="audio" muted="muted" controls></audio>
</template>