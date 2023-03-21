<script setup>
import {dateZhCN, NButton, NConfigProvider, NDropdown, NIcon, NLayout, NMessageProvider, zhCN} from 'naive-ui'
import {h} from 'vue'
import {SparklesOutline} from '@vicons/ionicons5'
import {RouterLink} from "vue-router";
// import HelloWorld from './components/HelloWorld.vue'
// This starter template is using Vue 3 experimental <script setup> SFCs
// Check out https://github.com/vuejs/rfcs/blob/master/active-rfcs/0040-script-setup.md


const options = [
  {
    label: 'super menu',
    key: 'marina bay sands',
    disabled: false
  },
  {
    label: () => h(
        RouterLink,
        {
          to: {
            path: '/manager/',
          }
        },
        {default: () => 'manager'}),
    key: "manager"
  },

  {
    label: () => h(
        RouterLink,
        {
          to: {
            path: '/home/',
          }
        },
        {default: () => 'home'}),
    key: "home"
  },
]

function handleSelect(key) {
  console.log(key)
}
</script>

<template>
  <!--  <img alt="Vue logo" src="./assets/logo.png" />-->
  <!--  <HelloWorld msg="Hello Vue 3 + Vite" />-->
  <n-message-provider>
    <n-config-provider :locale="zhCN" :date-locale="dateZhCN">
      <n-layout position="absolute">
        <n-layout position="absolute">
          <router-view></router-view>
        </n-layout>

        <div
            style="right: 60px;bottom: 100px;position:absolute"
        >
          <n-dropdown trigger="hover" :options="options" @select="handleSelect">
            <n-button circle size="large">
              <template #icon>
                <n-icon>
                  <sparkles-outline/>
                </n-icon>
              </template>
            </n-button>
          </n-dropdown>

        </div>
      </n-layout>
    </n-config-provider>
  </n-message-provider>
</template>

<style>
@import "app.css";
</style>
