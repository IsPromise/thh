<template>
  <div class="chat-box">
    <div v-for="(message, index) in messages" :key="index" class="message-row">
      <div class="avatar">
        <n-avatar :src="message.avatar"></n-avatar>
      </div>
      <div class="message-content" :class="{ 'my-message': message.from === currentUser }">
        <p v-html="formatMessage(message.content)"></p>
      </div>
    </div>
    <div class="input-box">
      <n-input
          placeholder="请输入消息..."
          v-model:value.trim="currentMessage"
      ></n-input>
      <n-button @click="sendMessage">发送</n-button>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { NAvatar, NInput, NButton } from "naive-ui";

const messages = ref([
  {
    content: "你好，欢迎来到聊天室！",
    from: "system",
    avatar: "",
  },
]);

const currentUser = ref("me");
const currentUserAvatar = ref("");

const currentMessage = ref("");

function sendMessage() {
  if (!currentMessage.value.trim()) return;
  messages.value.push({
    content: currentMessage.value,
    from: currentUser.value,
    avatar: currentUserAvatar.value,
  });
  currentMessage.value = "";
}

function addLineBreak(event) {
  if (event.keyCode === 13 && !event.ctrlKey) {
    event.preventDefault();
    sendMessage();
  }
  if (event.keyCode === 13 && event.ctrlKey) {
    currentMessage.value += "\n";
  }
}

function formatMessage(message) {
  return message.replace(/\n/g, "<br>");
}
</script>

<style scoped>
.chat-box {
  display: flex;
  flex-direction: column;
  height: 100%;
  padding: 16px;
}

.message-row {
  display: flex;
  align-items: center;
  margin-bottom: 8px;
}

.avatar {
  margin-right: 8px;
}

.message-content {
  padding: 0 6px;
  border-radius: 50px;
  background-color: #eeeeee;
  max-width: 70%;
}

.my-message {
  background-color: #0084ff;
  color: white;
}

.input-box {
  margin-top: auto;
  display: flex;
  align-items: center;
  margin-bottom: 16px;
}
</style>
