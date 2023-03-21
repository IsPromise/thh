package routes

import (
	"github.com/gin-gonic/gin"
	"thh/app/http/controllers"
)

func ginBBS(ginApp *gin.Engine) {
	bbs := ginApp.Group("bbs")
	// 发布文章
	bbs.Any("write-articles", ginUpP(controllers.WriteArticles))
	// 文章列表
	bbs.Any("get-articles", ginUpP(controllers.GetArticles))
	// 文章详情
	bbs.Any("get-articles-detail", ginUpP(controllers.GetArticlesDetail))
	// 发布评论
	bbs.Any("articles-comment", ginUpP(controllers.ArticleComment))

	// 热门链接
	// 用户主页
	// tag/分类

}
