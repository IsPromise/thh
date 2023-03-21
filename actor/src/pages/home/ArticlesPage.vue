<script setup>
import {NCard, NList, NListItem, NSpace, NTag, NThing} from 'naive-ui'
import {onMounted, ref} from "vue";
import {useRouter} from "vue-router";
import {remoteService} from "@/service/remote";


const listData = ref([])
const commentList = ref([])

let id = 1
let maxCommentId = 0

function getArticlesDetail() {
  remoteService.getArticlesDetail(id, maxCommentId).then(r => {
    console.log(r.data.data)
    listData.value.push(
        {
          title: "title1",
          tag: ["tag1", "tag2"],
          desc: "bodybodybodybodybodybodybody",
          createDate: "2022-12-28 01:01:01",
          lastUpdateDate: "2022-12-28 01:01:01",
          body: r.data.data.articleContent
        })
    let commentData = r.data.data.commentList.map(function (item) {
      return {
        username: "" + item.userId,
        content: item.content,
      }
    })
    commentList.value.push(
        ...commentData
    )
  }).catch(e => {
    console.error(e)
  })
}

onMounted(() => {
  const router = useRouter();
  id = router.currentRoute.value.query.id
  getArticlesDetail()
})
</script>
<template>
  <n-card style="margin:0 auto">
    <n-thing>
      你好你好你好
    </n-thing>
  </n-card>
  <n-card style="margin:0 auto">
    <n-list>
      <n-list-item v-for="item in listData">

        <n-thing :title="item.title" content-style="margin-top: 10px;">
          <template #description>
            <n-space size="small" style="margin-top: 4px">
              <n-tag v-for="itemTag in item.tag" :bordered="false" type="info" size="small" v-text="itemTag">
              </n-tag>
            </n-space>
          </template>
          <span v-text="item.body"></span>
        </n-thing>
      </n-list-item>
    </n-list>
  </n-card>
  <n-card style="margin:0 auto">
    <n-list>
      <n-list-item v-for="item in commentList">

        <n-thing :title="item.userId" content-style="margin-top: 10px;">
          <span v-text="item.content"></span>
        </n-thing>
      </n-list-item>
    </n-list>
  </n-card>
</template>