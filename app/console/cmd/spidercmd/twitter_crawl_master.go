package spidercmd

import (
	"fmt"
	"math/rand"
	"net/url"
	"path"
	"path/filepath"
	"regexp"
	"sync"
	"thh/app/models/FTwitter/FTwitterMedia"
	"thh/app/models/FTwitter/FTwitterTweet"
	"thh/app/service/ropt"
	"thh/app/service/twservice"
	"thh/bundles/logging"
	"thh/bundles/myfmt"
	"time"

	"github.com/leancodebox/goose/array"
	"github.com/leancodebox/goose/jsonopt"
	"github.com/leancodebox/goose/memqueue"

	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

const (
	userinfoType = iota
	tweetListType
	followListType
)

func init() {
	appendCommand(&cobra.Command{
		Use:   "spider:twitter:start",
		Short: "主程",
		Run:   spiderTwitterMainAction,
		//Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
	})
}

func ifErr(err error) bool {
	if err != nil {
		myfmt.PrintlnWithCaller(err)
		return true
	}
	return false
}

type spiderTwitterConfig struct {
	screenName string
	usePush    bool
	downMedia  bool
	spiderDeep int // 抓取深度，之后使用这个变量来控制
}

type QueueUnit struct {
	ScreenName string
	Deep       int
}

func spiderTwitterMainAction(_ *cobra.Command, _ []string) {
	SpiderTwitterMain()
}

func SpiderTwitterMain() {

	var screenNameMap map[string]bool
	screenNameMap = make(map[string]bool, 2048)
	var maxRoutineNum = 3
	if useProxy {
		ropt.SetProxy(proxy)
	}
	resp, _, err := ropt.Get("https://abs.twimg.com/responsive-web/client-web/main.b5030eda.js")
	if err != nil {
		myfmt.PrintlnWithCaller("获取queryId失败", err)
		return
	}
	regUnit := func(regStr string, matchStr string) (result [][]string) {
		defer func() {
			if p := recover(); p != nil {
				fmt.Printf("panic: %s\n", p)
			}
		}()
		reg := regexp.MustCompile(regStr)
		//根据规则提取关键信息
		result = reg.FindAllStringSubmatch(matchStr, -1)
		return
	}
	result := regUnit(`queryId:\"([a-zA-Z0-9\-]+)\",operationName:\"([a-zA-Z0-9]+)\"`, resp)
	queryIdMap := map[string]string{}
	for _, item := range result {
		if len(item) == 3 {
			queryIdMap[item[2]] = item[1]
		}
	}
	//myfmt.PrintlnWithCaller(jsonopt.Encode(queryIdMap))

	ch := make(chan int, maxRoutineNum)

	stdToolClient = newToolClient()

	if len(getScreenNameSlice()) == 0 {
		myfmt.PrintlnWithCaller("当前无配置")
		return
	}

	for _, jobScreenName := range getScreenNameSlice() {
		memqueue.QueueRPushObj(queueKey, QueueUnit{jobScreenName, 0})
	}

	time.Sleep(15 * time.Second)
	var wg sync.WaitGroup
	for {
		qu, err := memqueue.QueueLPopObj[QueueUnit](queueKey)
		screenName := qu.ScreenName
		if err != nil {
			break
		}
		if screenNameMap[screenName] {
			logging.Info(screenName + "当前已经查询过，跳过")
			continue
		}
		screenNameMap[screenName] = true
		ch <- 1
		wg.Add(1)
		go func(screenName string, usePush bool, ch chan int) {
			defer wg.Done()
			spiderTwitterList(spiderTwitterConfig{
				screenName: screenName,
				usePush:    usePush,
				downMedia:  needDownMedia(),
			})
			<-ch
		}(screenName, allUsePush, ch)
	}
	wg.Wait()
}

func spiderTwitterList(sConfig spiderTwitterConfig) {

	screenName := sConfig.screenName
	usePush := sConfig.usePush
	client := newTClient()
	twitterMediaDir := filepath.Join(outputPrefix, "/response_"+screenName+"media_source/")
	r, err := client.getUserInfo(screenName)
	twservice.SaveTSpiderHis(userinfoType, screenName+"_userinfo_"+cast.ToString(time.Now().UnixMilli()), r, err)
	if ifErr(err) {
		return
	}
	userInfo := jsonopt.Decode[TUserInfo](r.String())
	restId := userInfo.Data.User.Result.RestID
	desc := userInfo.Data.User.Result.Legacy.Description
	name := userInfo.Data.User.Result.Legacy.Name

	twservice.SaveUserEntity(restId, screenName, desc, name)

	var linkList []string
	for _, value := range userInfo.Data.User.Result.Legacy.Entities.URL.Urls {
		linkList = append(linkList, value.ExpandedURL)
	}

	cursor := ""

	// 谁 发了 什么
	// 谁 用户id 用户名 用户简介
	// 发了什么 内容 图片 视频地址
	i := 0
	for {
		r, err = client.getTList(restId, 40, cursor)
		twservice.SaveTSpiderHis(tweetListType, screenName+"_tweetList_"+cast.ToString(i), r, err)
		tweetResponse := jsonopt.Decode[UserTweetsResponse](r.String())
		i++
		if len(tweetResponse.Data.User.Result.TimelineV2.Timeline.Instructions) == 0 || i >= spiderTwitterMaxPage {
			logging.Info(screenName + "完成··································")
			break
		}

		activeCount := 0
		for _, value := range tweetResponse.Data.User.Result.TimelineV2.Timeline.Instructions {
			switch value.Type {
			case "TimelineAddEntries":
				for _, entry := range value.Entries {
					entryContent := entry.Content
					if entryContent.EntryType == "TimelineTimelineItem" {
						activeCount += 1
						//userId := entry.Content.ItemContent.TweetResults.Result.RestId
						// 用户
						// masterUserResult = entry.Content.ItemContent.TweetResults.Result.Core.UserResults.Result
						// 当前作者
						// userResult := entry.Content.ItemContent.TweetResults.Result.Core.UserResults.Result
						// 原文作者（如果没有可能为非转发）
						// userResult := entry.Content.ItemContent.TweetResults.Result.Legacy.RetweetedStatusResult.Result.Core.UserResults.Result

						orgUserResult := entry.Content.ItemContent.TweetResults.Result.Legacy.RetweetedStatusResult.Result.Core.UserResults.Result
						//masterUserResult := entry.Content.ItemContent.TweetResults.Result.Core.UserResults.Result

						conversationIdStr := entry.Content.ItemContent.TweetResults.Result.Legacy.ConversationIdStr
						itemFullText := entry.Content.ItemContent.TweetResults.Result.Legacy.FullText

						//原文时间不适合用来判断
						//createTime := str2time(entry.Content.ItemContent.TweetResults.Result.Legacy.CreatedAt)
						//if createTime.Unix() < time.Now().Add(-86400*time.Second).Unix() {
						//	myfmt.PrintlnWithCaller("创建时间", createTime.Format("2006-01-02 15:04:05"))
						//	break
						//}

						isForwarded := true
						// 原作者不存在当前非转发
						if orgUserResult.Legacy.ScreenName == "" {
							isForwarded = false
						}
						// 允许后续扩散查询 非原创 且深度
						if usePush && isForwarded && sConfig.spiderDeep < spiderDeep {
							memqueue.QueueRPushObj(queueKey, QueueUnit{orgUserResult.Legacy.ScreenName, sConfig.spiderDeep + 1})
							myfmt.PrintlnWithCaller(orgUserResult.Legacy.ScreenName, "进入后续查询队列")
						}

						userTweetEntity := FTwitterTweet.GetUserTweet(screenName, conversationIdStr)
						if userTweetEntity.Id != 0 && !array.InArray(screenName, getScreenNameSlice()) {
							continue
						}
						userTweetEntity.ScreenName = screenName
						userTweetEntity.Context = itemFullText
						userTweetEntity.ConversationId = conversationIdStr
						if isForwarded {
							userTweetEntity.OriginScreenName = orgUserResult.Legacy.ScreenName
						} else {
							userTweetEntity.OriginScreenName = screenName
						}
						userTweetEntity.CreateTime = time.Now()
						FTwitterTweet.Save(&userTweetEntity)

						// 当为转发，且属于目标转发，进行木匾转发统计，和原用户信息录入
						if isForwarded && array.InArray(screenName, getScreenNameSlice()) {
							userInLegacy := orgUserResult.Legacy
							twservice.SaveUserEntity(orgUserResult.RestId, userInLegacy.ScreenName, userInLegacy.Description, userInLegacy.Name)
						}
						// 推文
						// entry.Content.ItemContent.TweetResults.Result.Legacy

						medias := entry.Content.ItemContent.TweetResults.Result.Legacy.ExtendedEntities.Media
						for _, media := range medias {
							switch media.Type {
							case "animated_gif":
								// todo
							case "photo":
								//u, _ := url.Parse(media.MediaUrlHttps)
								basename := path.Base(media.MediaUrlHttps)
								stdToolClient.downMedia(media.MediaUrlHttps, twitterMediaDir+conversationIdStr+basename)
								FTwitterMedia.Save(&FTwitterMedia.FTwitterMedia{Type: "photo", TweetId: conversationIdStr, Path: twitterMediaDir + conversationIdStr + basename, Url: media.MediaUrlHttps})
							case "video":
								// 下载封面
								basename := path.Base(media.MediaUrlHttps)
								stdToolClient.downMedia(media.MediaUrlHttps, twitterMediaDir+conversationIdStr+basename)
								FTwitterMedia.Save(&FTwitterMedia.FTwitterMedia{Type: "video_photo", TweetId: conversationIdStr, Path: twitterMediaDir + conversationIdStr + basename, Url: media.MediaUrlHttps})
								// 下载视频
								variants := media.VideoInfo.Variants
								tmpBitrate := 0
								tmpUrl := ""
								for _, variant := range variants {
									if variant.Bitrate > tmpBitrate {
										u, pErr := url.Parse(variant.Url)
										if pErr != nil {
											myfmt.PrintlnWithCaller("url解析失败")
											continue
										}
										basename = path.Base(u.Path)
										tmpUrl = variant.Url
									}
								}
								if len(tmpUrl) == 0 {
									myfmt.PrintlnWithCaller("视频下载失败")
									break
								}
								stdToolClient.downMedia(tmpUrl, twitterMediaDir+conversationIdStr+basename)
								FTwitterMedia.Save(&FTwitterMedia.FTwitterMedia{Type: "video_photo", TweetId: conversationIdStr, Path: twitterMediaDir + conversationIdStr + basename, Url: tmpUrl})
							default:
								myfmt.PrintlnWithCaller(media.Type)
							}
						}
					}
					// 选择下次节点
					if entryContent.CursorType == "Bottom" {
						cursor = entryContent.Value
					}
				}
				break
			case "TimelinePinEntry":
				for _, entry := range value.Entries {
					userId := entry.Content.ItemContent.TweetResults.Result.RestId
					myfmt.PrintlnWithCaller(userId)
				}
				break
			case "TimelineClearCache":
				break
			case "animated_gif":
				break
			default:
				myfmt.PrintlnWithCaller(value.Type)
			}
		}

		if activeCount == 0 {
			myfmt.PrintlnWithCaller(screenName + "完成··································")
			break
		}
		//myfmt.PrintlnWithCaller(screenName, "下一轮", i*40, "-", (i+1)*40)

		time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)

	}

	// Want to remove proxy setting
	//client.RemoveProxy()
}

func str2time(s string) time.Time {
	// Sat Aug 13 03:37:20 +0000 2022
	// Mon Jan 03 15:04:05 -0700 2001
	var LOC, _ = time.LoadLocation("Asia/Shanghai")
	//timeTemplate := "2006-01-02 15:04:05"
	tim, _ := time.ParseInLocation(time.RubyDate, s, LOC)
	return tim
}
