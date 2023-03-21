<script setup>
import {
  NButton,
  NCard,
  NCarousel,
  NDivider,
  NEmpty,
  NGrid,
  NGridItem,
  NIcon,
  NImage,
  NImageGroup,
  NList,
  NListItem,
  NResult,
  NSpace,
  NStatistic,
  NTag,
  NThing
} from 'naive-ui'
import {ref} from 'vue'
import {useIsMobile, useIsSmallDesktop, useIsTablet} from "@/utils/composables";

const listData = ref([
  {
    title: "title1",
    tag: ["tag1", "tag2"],
    body: "bodybodybodybodybodybodybody"
  }
])

function sorlly() {
  let scrollTop = document.documentElement.scrollTop || document.body.scrollTop;
  //  可视区域
  let clientHeight = document.documentElement.clientHeight;
  // 页面的高度
  let scrollHeight = document.documentElement.scrollHeight;
  //如果触底就让index++
  if (scrollTop + clientHeight >= scrollHeight) {
    listData.value.push({
      title: "title1",
      tag: ["tag1", "tag2"],
      body: "bodybodybodybodybodybodybody"
    })
  }
}

// onMounted(() => {
// window.addEventListener('scroll',sorlly)
// })
const isMobileRef = useIsMobile()
const isTabletRef = useIsTablet()
const isSmallDesktop = useIsSmallDesktop()

const isMobile = useIsMobile()


let sessionStorageData = ref("")
let localStorageData = ref("")


function set() {
  let localTmp = localStorage.getItem("tmp")
  localStorage.setItem("tmp", (Number(localTmp) + 1).toString())
  let sessionTmp = sessionStorage.getItem("tmp")
  sessionStorage.setItem("tmp", (Number(sessionTmp) + 1).toString())
  sessionStorageData.value = JSON.stringify(sessionStorage)
  localStorageData.value = JSON.stringify(localStorage)
}

function showNew() {
  sessionStorageData.value = JSON.stringify(sessionStorage)
  localStorageData.value = JSON.stringify(localStorage)
}
</script>
<template>
  <n-space vertical>
    <n-card>
      <n-statistic label="SessionStorage">
        {{ sessionStorageData }}
      </n-statistic>
      <n-statistic label="LocalStorage">
        {{ localStorageData }}
      </n-statistic>
      <span>设置 sessionStorage/localStorage 如果存在+1 </span>
      <br>
      <n-button @click="set"> 设置</n-button>
      <n-button @click="showNew"> 展示最新</n-button>
    </n-card>
    <n-card>
      Mobile:<span>{{ isMobile }}</span>
      |Tablet:<span>{{ isTabletRef }}</span>
      |SmallDesktop:<span>{{ isSmallDesktop }}</span>
    </n-card>
    <n-card title="url拼接">
      <n-grid>
        <n-grid-item :span="24">
          <n-empty size="large" description="可以是大的">
            <template #extra>
              <n-button size="small">看看别的</n-button>
            </template>
          </n-empty>
        </n-grid-item>
      </n-grid>
    </n-card>


    <n-card title="统计数据">
      <n-grid>
        <n-grid-item :span="12">
          <n-statistic label="统计数据" :value="99">
            <template #prefix>
              <n-icon>

              </n-icon>
            </template>
            <template #suffix>/ 100</template>
          </n-statistic>
        </n-grid-item>
        <n-grid-item :span="12">
          <n-statistic label="活跃用户">1,234,123</n-statistic>
        </n-grid-item>
      </n-grid>
    </n-card>

    <n-card>
      <n-grid>
        <n-grid-item :span="24">
          <n-image-group>
            <n-space>
              <n-image
                  width="100"
                  src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg"
              />
              <n-image
                  width="100"
                  src="https://gw.alipayobjects.com/zos/antfincdn/aPkFc8Sj7n/method-draw-image.svg"
              />
            </n-space>
          </n-image-group>
        </n-grid-item>
      </n-grid>
    </n-card>

    <n-card>
      <n-grid>
        <n-grid-item :span="24">
          <n-result status="404" title="！" description="这么大" size="huge">
            <template #footer>
              <n-button>哦</n-button>
            </template>
          </n-result>
        </n-grid-item>
      </n-grid>
      <n-divider>神奇分割符</n-divider>
      <n-grid>
        <n-grid-item :span="24">
          <n-carousel autoplay>
            <img
                class="carousel-img"
                src="https://s.anw.red/fav/1623979004.jpg!/fw/600/quality/77/ignore-error/true"
            />
            <img
                class="carousel-img"
                src="https://s.anw.red/news/1623372884.jpg!/both/800x450/quality/78/progressive/true/ignore-error/true"
            />
            <img
                class="carousel-img"
                src="https://s.anw.red/news/1623177220.jpg!/both/800x450/quality/78/progressive/true/ignore-error/true"
            />
            <img
                class="carousel-img"
                src="https://s.anw.red/news/1623152423.jpg!/both/800x450/quality/78/progressive/true/ignore-error/true"
            />
          </n-carousel>
        </n-grid-item>
      </n-grid>
    </n-card>
    <n-card>
      <n-list>
        <n-list-item>
          <n-thing title="相见恨晚" content-style="margin-top: 10px;">
            <template #description>
              <n-space size="small" style="margin-top: 4px">
                <n-tag :bordered="false" type="info" size="small">
                  暑夜
                </n-tag>
                <n-tag :bordered="false" type="info" size="small">
                  晚春
                </n-tag>
              </n-space>
            </template>
            奋勇呀然后休息呀<br>
            完成你伟大的人生
          </n-thing>
        </n-list-item>

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
  </n-space>
</template>
<style>
.carousel-img {
  width: 100%;
  height: 240px;
  object-fit: cover;
}
</style>