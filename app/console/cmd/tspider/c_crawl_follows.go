package tspider

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"thh/app/service/twservice"
	"thh/bundles/logging"
	"time"

	"github.com/leancodebox/goose/jsonopt"

	"github.com/leancodebox/goose/preferences"

	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func init() {
	appendCommand(&cobra.Command{
		Use:   "t:spider:follow",
		Short: "follow 抓取",
		Run:   tFollow,
		//Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
	})
}

// tFollow 抓取关注的列表
func tFollow(_ *cobra.Command, _ []string) {
	var maxRoutineNum = 3
	rootPrefix = preferences.GetString("sprider.twitter.output", "./storage/tmp/")
	outputPrefix = rootPrefix + time.Now().Format("20060102_150405")
	queueKey = "twitter:screenName:list"
	downMedia := preferences.GetBool("sprider.twitter.downmedia", false)
	screenNamesFromEnv := preferences.GetString("sprider.twitter.screename", "")
	ch := make(chan int, maxRoutineNum)
	var wg4master sync.WaitGroup

	stdToolClient = newToolClient()

	dataList := strings.Split(screenNamesFromEnv, ",")

	if len(dataList) == 0 {
		fmt.Println("当前无配置")
		return
	}

	for _, jobScreenName := range dataList {
		wg4master.Add(1)
		ch <- 1
		go func(screenName string, ch chan int) {
			defer wg4master.Done()
			superFollow(superTConfig{
				screenName: screenName,
				usePush:    true,
				downMedia:  downMedia,
			})
			<-ch
		}(jobScreenName, ch)
	}
	wg4master.Wait()

	fmt.Println("抓取关注人列表完毕")
}

func superFollow(sConfig superTConfig) {
	//tScreenNameList := config.GetString("T_SCREENAME", "")
	screenName := sConfig.screenName
	// Create a Resty Client
	client := newTClient()

	r, err := client.getUserInfo(screenName)
	twservice.SaveTSpiderHis(userinfoType, screenName+"_userinfo_"+cast.ToString(time.Now().UnixMilli()), r, err)
	if ifErr(err) {
		return
	}
	userInfo := jsonopt.Decode[TUserInfo](r.String())
	restId := userInfo.Data.User.Result.RestID
	desc := userInfo.Data.User.Result.Legacy.Description
	name := userInfo.Data.User.Result.Legacy.Name

	if restId == "" {
		fmt.Println("信息获取失败")
	} else {
		twservice.SaveUserEntity(restId, screenName, desc, name)
	}

	var linkList []string
	for _, value := range userInfo.Data.User.Result.Legacy.Entities.URL.Urls {
		linkList = append(linkList, value.ExpandedURL)
	}

	cursor := ""

	// 谁 发了 什么
	// 谁 用户id 用户名 用户简介
	// 发了什么 内容 图片 视频地址
	i := 0
	pageCount := 20
	for {
		r, err = client.getFollowList(restId, pageCount, cursor)
		twservice.SaveTSpiderHis(followListType, screenName+"_follow_"+cast.ToString(i)+cast.ToString(time.Now().UnixMilli()), r, err)

		TList := jsonopt.Decode[TFollowList](r.String())

		i++

		if len(TList.Data.User.Result.Timeline.Timeline.Instructions) == 0 {
			logging.Info(screenName + "完成")
			break
		}

		activeCount := 0
		for _, value := range TList.Data.User.Result.Timeline.Timeline.Instructions {
			switch value.Type {
			case "TimelineAddEntries":
				for _, entry := range value.Entries {
					entryContent := entry.Content
					if entryContent.EntryType == "TimelineTimelineItem" {
						activeCount += 1
						// 用户
						// masterUserResult = entry.Content.ItemContent.TweetResults.Result.Core.UserResults.Result
						// 当前作者
						// userResult := entry.Content.ItemContent.TweetResults.Result.Core.UserResults.Result
						// 原文作者（如果没有可能为非转发）

						userResult := entry.Content.ItemContent.UserResults.Result

						if len(userResult.Legacy.ScreenName) > 0 {
							userInLegacy := userResult.Legacy
							twservice.SaveUserEntity(userResult.RestId, userInLegacy.ScreenName, userInLegacy.Description, userInLegacy.Name)
						}
						// 推文
						// entry.Content.ItemContent.TweetResults.Result.Legacy

					}
					// 选择下次节点
					if entryContent.CursorType == "Bottom" {
						cursor = entryContent.Value
					}
				}
				break
			default:
				fmt.Println(value.Type)
			}
		}

		if activeCount == 0 {
			logging.Info(screenName + "完成")
			break
		}
		logging.Info(screenName, "下一轮", i*pageCount, "-", (i+1)*pageCount)

		time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)

	}

	// Want to remove proxy setting
	//client.RemoveProxy()
}
