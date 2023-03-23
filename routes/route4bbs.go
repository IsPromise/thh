package routes

import (
	"github.com/gin-gonic/gin"
	"thh/app/http/controllers"
	"thh/app/http/middleware"
)

func ginBBS(ginApp *gin.Engine) {
	bbs := ginApp.Group("api/bbs")
	// 文章列表
	bbs.POST("get-articles", ginUpJP(controllers.GetArticles))
	// 文章详情
	bbs.POST("get-articles-detail", ginUpJP(controllers.GetArticlesDetail))
	// 热门链接
	// 用户主页
	// tag/分类

	bbsAuth := bbs.Use(middleware.JWTAuth4Gin)
	// 发布文章
	bbsAuth.POST("write-articles", UpButterReq(controllers.WriteArticles))
	// 发布评论
	bbsAuth.POST("articles-comment", ginUpJP(controllers.ArticleComment))

}
