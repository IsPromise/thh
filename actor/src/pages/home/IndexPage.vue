<script setup>
import {NList, NListItem, useMessage} from 'naive-ui';
import {ref} from "vue";
import {addDays, isYesterday} from "date-fns";


const message = useMessage()


function success() {
  message.success(
      `还挺大`
  )
}


const value = ref(addDays(Date.now(), 1).valueOf())

function handleUpdateValue(_, {year, month, date}) {
  message.success(`${year}-${month}-${date}`)
}

function isDateDisabled(timestamp) {
  return isYesterday(timestamp);
}
</script>
<template>
  <n-list class="thh-index" :show-divider="false">
    <n-list-item>
      <div class="greetings" style=" text-align: center;">
        <h1 class="green">thh</h1>
        <p>候鸟归来时 羽翼扑满尘</p>
        <p>容我再等你一个转身</p>
        <p>斜阳对黄昏 谁的容颜老三分</p>
        <p>巷子里酒香 越传越深</p>
        <br>
        <p>莫等到白发催人泪 徒留悔恨增伤悲</p>
        <p>莫使金樽空对月 今朝有酒今朝醉</p>
        <p>莫怪青山换不回 念天地悠悠轮回</p>
        <p>弹指一挥这一生 也无悔</p>
        <br>
        <p>候鸟归来时 羽翼扑满尘</p>
        <p>容我再等你一个转身</p>
        <p>斜阳对黄昏 谁的容颜老三分</p>
        <p>巷子里酒香 越传越深</p>
        <br>
        <p>莫等到白发催人泪 徒留悔恨增伤悲</p>
        <p>莫使金樽空对月 今朝有酒今朝醉</p>
        <p>莫怪青山换不回 念天地悠悠轮回</p>
        <p>弹指一挥这一生 也无悔</p>
        <br>
        <p>莫等到山河都踏碎 徒留行囊换愚昧</p>
        <p>莫使年华付流水 敢爱敢恨把你追</p>
        <p>莫怪烛光不够美 千古诗篇来赞美</p>
        <p>我许下山盟海誓 永相随</p>
      </div>
    </n-list-item>
  </n-list>
</template>

<style scoped>
h1 {
  font-weight: 500;
  font-size: 2.6rem;
  top: -10px;
}

h3 {
  font-size: 1.2rem;
}

.greetings {
  max-width: 1280px;
  margin: 0 auto;
}

.greetings h1,
.greetings h3 {
  text-align: center;
}

/*@media (min-width: 1024px) {*/
/*  .greetings h1,*/
/*  .greetings h3 {*/
/*    text-align: left;*/
/*  }*/
/*}*/
.thh-index .n-list-item {
  padding-left: 20px;
  padding-right: 20px;
}

.green {
  text-decoration: none;
  color: hsla(160, 100%, 37%, 1);
  transition: 0.4s;
}

</style>