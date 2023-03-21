package controllers

import (
	"fmt"
	"reflect"
	"sync"
	"thh/app/console/cmd/tspider"
	"thh/app/http/controllers/component"
	"thh/app/models/FTwitter/FTwitterSpiderHis"
	"thh/app/models/FTwitter/FTwitterTweet"
	"thh/app/models/FTwitter/FTwitterUser"
	"thh/arms"

	"github.com/spf13/cast"
)

type TLink struct {
	ScreenName       string
	OriginScreenName string `json:"originScreenName,omitempty"`
	Name             string
	Desc             string
	Url              string
	CreateTime       string
}
type TListRequest struct {
	SearchList []string
}

func TListV2(request TListRequest) component.Response {
	list := []TLink{}
	for _, desc := range request.SearchList {
		if desc == "" {
			continue
		}
		list1 := arms.ArrayMap(func(item FTwitterUser.FTwitterUser) TLink {
			return TLink{
				ScreenName: item.ScreenName,
				Name:       item.Name,
				Desc:       item.Desc,
				Url:        fmt.Sprintf("https://twitter.com/%v/with_replies", item.ScreenName),
				CreateTime: item.CreateTime.Format("2006-01-02 15:05:05"),
			}
		}, FTwitterUser.GetByDesc(desc))
		list = append(list, list1...)
		list2 := arms.ArrayMap(func(item FTwitterTweet.FTwitterTweet) TLink {
			return TLink{
				ScreenName: item.ScreenName,
				Name:       "tweet",
				Desc:       item.Context,
				Url:        fmt.Sprintf("https://twitter.com/%v/status/%v", item.ScreenName, item.ConversationId),
				CreateTime: item.CreateTime.Format("2006-01-02 15:05:05"),
			}
		}, FTwitterTweet.GetByContent(desc))

		list = append(list, list2...)
	}

	return component.SuccessResponse(list)
}

type GetTwitterTweetListParam struct {
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
	Search   string `form:"search"`
}

func GetTwitterTweetList(param GetTwitterTweetListParam) component.Response {
	pageData := FTwitterTweet.Page(FTwitterTweet.PageQuery{
		Page: param.Page, PageSize: param.PageSize, Search: param.Search,
	})
	return component.SuccessResponse(component.DataMap{
		"itemList": arms.ArrayMap(func(item FTwitterTweet.FTwitterTweet) TLink {
			return TLink{
				ScreenName:       item.ScreenName,
				OriginScreenName: item.OriginScreenName,
				Name:             "tweet",
				Desc:             item.Context,
				Url:              fmt.Sprintf("https://twitter.com/%v/status/%v", item.ScreenName, item.ConversationId),
				CreateTime:       item.CreateTime.Format("2006-01-02 15:05:05"),
			}
		}, pageData.Data),
		"size":    pageData.PageSize,
		"total":   pageData.Total,
		"current": param.Page,
	})
}

type GetTwitterUserListParam struct {
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
	Search   string `form:"search"`
}

func GetTwitterUserList(param GetTwitterUserListParam) component.Response {
	pageData := FTwitterUser.Page(FTwitterUser.PageQuery{
		Page: param.Page, PageSize: param.PageSize, Search: param.Search,
	})

	return component.SuccessResponse(component.DataMap{
		"itemList": arms.ArrayMap(func(item FTwitterUser.FTwitterUser) TLink {
			return TLink{
				ScreenName: item.ScreenName,
				Name:       item.Name,
				Desc:       item.Desc,
				Url:        fmt.Sprintf("https://twitter.com/%v/with_replies", item.ScreenName),
				CreateTime: item.CreateTime.Format("2006-01-02 15:05:05"),
			}
		}, pageData.Data),
		"size":    pageData.PageSize,
		"total":   pageData.Total,
		"current": param.Page,
	})
}

var tMasterRunningLock sync.Mutex

func RunTSpiderMaster() component.Response {
	if tMasterRunningLock.TryLock() {
		go func() {
			defer tMasterRunningLock.Unlock()
			tspider.TMasterRun()
		}()
		return component.SuccessResponse(component.DataMap{
			"message": "success start",
		})
	} else {
		return component.SuccessResponse(component.DataMap{
			"message": "上次任务未完成",
		})
	}
}

func GetQueueLen() component.Response {
	return component.SuccessResponse(component.DataMap{
		"message": "当前队列长度:" + cast.ToString(arms.QueueLen("twitter:screenName:list")),
	})
}

type GetTSpiderReq struct {
	Page     int `form:"page"`
	PageSize int `form:"pageSize"`
}

type THisItem struct {
	Id         uint64 `json:"id"`
	Curl       string `json:"curl"`
	Content    string `json:"content"`
	CreateTime string `json:"createTime"`
}

func GetKeyItemList(data any) (result []string) {
	t := reflect.TypeOf(data).Elem()
	for i := 0; i < t.NumField(); i++ {
		key := t.Field(i).Tag.Get("json")
		if key != "" {
			result = append(result, key)
		}
	}
	return result
}

var THisItemKey = GetKeyItemList(&THisItem{})

func GetTSpiderHis(req GetTSpiderReq) component.Response {
	pageData := FTwitterSpiderHis.Page(FTwitterSpiderHis.PageQuery{
		Page:     req.Page,
		PageSize: req.PageSize,
	})

	return component.SuccessResponse(component.DataMap{
		"keyList": THisItemKey,
		"itemList": arms.ArrayMap(func(item FTwitterSpiderHis.FTwitterSpiderHis) THisItem {
			sContent := "有数据"
			if item.Content == "" {
				sContent = "无数据无数据无数据无数据无数据无数据"
			}
			return THisItem{
				Id:         item.Id,
				Curl:       item.Curl,
				Content:    sContent,
				CreateTime: item.CreateTime.Format("2006-01-02 15:04:05"),
			}
		}, pageData.Data),
		"size":    pageData.PageSize,
		"total":   pageData.Total,
		"current": req.Page,
	})
}