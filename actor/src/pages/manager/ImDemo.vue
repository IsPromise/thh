<script setup>
import {NButton, NInput, NLayout, NLayoutContent, NLayoutFooter} from "naive-ui";
import {onMounted, onUnmounted, ref} from "vue";

let ws;
const message = ref("")
const myMessage = ref("")

function print(data) {
  message.value = data
  console.log(data)
}

function open() {
  if (ws) {
    return false;
  }
  ws = new WebSocket("ws://localhost:90/ws");
  ws.onopen = function (evt) {
    print("连接websocket");
  }
  ws.onclose = function (evt) {
    print("CLOSE");
    ws = null;
  }
  ws.onmessage = function (evt) {
    let msg = JSON.parse(evt.data)
    print("收到消息: " + msg.message);
  }
  ws.onerror = function (evt) {
    print("ERROR: " + evt.data);
  }
  return false;
}

function send() {
  if (!ws) {
    return false;
  }

  print("发送消息: " + new Date() + myMessage.value);
  messageList.value.push(message.value);

  // contentRef.value.scrollTo({ top: 0, behavior: 'smooth' })
  console.log(contentRef.value)
  contentRef.value.scrollTo({top: 99999, behavior: 'smooth'})
  ws.send(myMessage.value);

  return false;
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

onUnmounted(()=>{
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