package controllers

import (
	"thh/app/http/controllers/component"
	"thh/app/models/Articles"
	"thh/app/models/Comment"
	"thh/arms"
	"time"
)

type GetArticlesRequest struct {
	MaxId    uint64 `json:"maxId"`
	PageSize int    `json:"pageSize"`
}

type ArticlesDto struct {
	Id      uint64 `json:"id"`
	Content string `json:"content"`
}

func GetArticles(request GetArticlesRequest) component.Response {
	if request.PageSize == 0 {
		request.PageSize = 10
	}
	articles := Articles.GetByMaxIdPage(request.MaxId, request.PageSize)
	var maxId uint64
	list := arms.ArrayMap(func(t Articles.Articles) ArticlesDto {
		maxId = t.Id
		return ArticlesDto{
			Id:      t.Id,
			Content: t.Content,
		}
	}, articles)

	return component.SuccessResponse(map[string]any{
		"maxId": maxId,
		"list":  list,
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
	article := Articles.Get(request.Id)
	comments := Comment.GetByMaxIdPage(request.Id, request.MaxCommentId, request.PageSize)

	commentList := arms.ArrayMap(func(item Comment.Comment) CommentDto {
		return CommentDto{
			ArticleId:  item.ArticleId,
			UserId:     item.UserId,
			Content:    item.Content,
			CreateTime: item.CreateTime.Format(time.RFC3339),
		}
	}, comments)
	return component.SuccessResponse(map[string]any{
		"articleContent": &article.Content,
		"commentList":    commentList,
	})

}

type WriteArticleReq struct {
	Id      int64  `json:"id"`
	Content string `json:"content"`
}

func WriteArticles(req component.BetterRequest[WriteArticleReq]) component.Response {
	if Articles.CantWriteNew(req.UserId, 66) {
		return component.FailResponse("您当天已发布较多，为保证质量，请明天再发布新帖")
	}
	var article Articles.Articles
	if req.Params.Id == 0 {
		article = Articles.Get(req.Params.Id)
	}
	article.Content = req.Params.Content
	Articles.Save(&article)
	return component.SuccessResponse(map[string]any{})
}

type ArticleCommentReq struct {
	ArticleId uint64 `json:"articleId"`
	Comment   string `json:"comment"`
}

func ArticleComment(req ArticleCommentReq) component.Response {
	if Articles.Get(req.ArticleId).Id == 0 {
		return component.FailResponse("文章不存在")
	}
	Comment.Save(&Comment.Comment{Content: req.Comment})
	return component.SuccessResponse(true)
}
