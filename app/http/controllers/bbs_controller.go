package controllers

import (
	"thh/app/http/controllers/component"
	"thh/app/models/FTwitter/FTwitterTweet"
	Articles2 "thh/app/models/bbs/Articles"
	Comment2 "thh/app/models/bbs/Comment"
	"thh/arms"
	"time"
)

type GetArticlesRequest struct {
	MaxId    uint64 `json:"maxId"`
	PageSize int    `json:"pageSize"`
}

type ArticlesDto struct {
	Id             uint64 `json:"id"`
	Content        string `json:"content"`
	Title          string `json:"title"`
	LastUpdateTime string `json:"lastUpdateTime"`
}

func GetArticles(request GetArticlesRequest) component.Response {
	if request.PageSize == 0 {
		request.PageSize = 10
	}
	articles := Articles2.GetByMaxIdPage(request.MaxId, request.PageSize)
	var maxId uint64
	if len(articles) > 0 {
		maxId = articles[0].Id
	}
	list := arms.ArrayMap(func(t Articles2.Articles) ArticlesDto {
		return ArticlesDto{
			Id:             t.Id,
			Title:          t.Title,
			Content:        t.Content,
			LastUpdateTime: t.UpdateTime.Format("2006-01-02 15:04:05"),
		}
	}, articles)

	return component.SuccessResponse(map[string]any{
		"maxId": maxId,
		"list":  list,
	})
}

type GetArticlesPageRequest struct {
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
	Search   string `form:"search"`
}

func GetArticlesPage(param GetArticlesPageRequest) component.Response {
	articles := Articles2.Page(Articles2.PageQuery{Page: param.Page, PageSize: param.PageSize})
	pageData := FTwitterTweet.Page(FTwitterTweet.PageQuery{
		Page: param.Page, PageSize: param.PageSize, Search: param.Search,
	})
	return component.SuccessResponse(component.DataMap{
		"list": arms.ArrayMap(func(t Articles2.Articles) ArticlesDto {
			return ArticlesDto{Id: t.Id,
				Title:          t.Title,
				Content:        t.Content,
				LastUpdateTime: t.UpdateTime.Format("2006-01-02 15:04:05"),
			}
		}, articles.Data),
		"size":    pageData.PageSize,
		"total":   pageData.Total,
		"current": param.Page,
	})
}

type GetArticlesDetailRequest struct {
	Id           uint64 `json:"id"`
	MaxCommentId uint64 `json:"maxCommentId"`
	PageSize     int    `json:"pageSize"`
}

type CommentDto struct {
	ArticleId  uint64 `json:"articleId"`
	UserId     uint64 `json:"userId"`
	Content    string `json:"content"`
	CreateTime string `json:"createTime"`
}

func GetArticlesDetail(request GetArticlesDetailRequest) component.Response {
	if request.PageSize == 0 {
		request.PageSize = 10
	}
	article := Articles2.Get(request.Id)
	comments := Comment2.GetByMaxIdPage(request.Id, request.MaxCommentId, request.PageSize)

	commentList := arms.ArrayMap(func(item Comment2.Comment) CommentDto {
		return CommentDto{
			ArticleId:  item.ArticleId,
			UserId:     item.UserId,
			Content:    item.Content,
			CreateTime: item.CreateTime.Format(time.RFC3339),
		}
	}, comments)
	return component.SuccessResponse(map[string]any{
		"articleTitle":   &article.Title,
		"articleContent": &article.Content,
		"commentList":    commentList,
	})

}

type WriteArticleReq struct {
	Id      int64  `json:"id"`
	Content string `json:"content" validate:"required"`
}

func WriteArticles(req component.BetterRequest[WriteArticleReq]) component.Response {
	if Articles2.CantWriteNew(req.UserId, 66) {
		return component.FailResponse("您当天已发布较多，为保证质量，请明天再发布新帖")
	}
	var article Articles2.Articles
	if req.Params.Id != 0 {
		article = Articles2.Get(req.Params.Id)
		if article.UserId != req.UserId {
			return component.FailResponse("不要更改别人发出的帖子哦")
		}
	} else {
		article.UserId = req.UserId
	}
	article.Content = req.Params.Content
	Articles2.Save(&article)
	return component.SuccessResponse(map[string]any{})
}

type ArticleCommentReq struct {
	ArticleId uint64 `json:"articleId"`
	Comment   string `json:"comment"`
}

func ArticleComment(req ArticleCommentReq) component.Response {
	if Articles2.Get(req.ArticleId).Id == 0 {
		return component.FailResponse("文章不存在")
	}
	Comment2.Save(&Comment2.Comment{Content: req.Comment})
	return component.SuccessResponse(true)
}
