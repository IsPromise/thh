package routes

import (
	"github.com/gin-gonic/gin"
	"thh/app/http/controllers"
	"thh/app/http/controllers/ginLowerControllers"
	"thh/app/http/middleware"
)

func ginApi(ginApp *gin.Engine) {
	ginApp.GET("/api", controllers.Api)

	apiGroup := ginApp.Group("api")
	// lowerControllers
	apiGroup.GET("gin-upload", ginLowerControllers.GinUpload)
	apiGroup.GET("/file/:filename", ginLowerControllers.GinUpload)
	apiGroup.GET("show-pic", ginLowerControllers.GinShowPic)

	apiGroup.GET("traefik-provider", ginLowerControllers.TraefikProvider)
	apiGroup.GET("memUse", ginUpNP(controllers.GetUseMem))
	apiGroup.GET("about", ginUpNP(controllers.About))
	apiGroup.GET("sys-info", ginUpNP(controllers.SysInfo))
	apiGroup.Use(middleware.IpLimit).GET("git-status-list", ginUpNP(controllers.GitStatusList))

	twitterApi := apiGroup.Group("twitter")
	twitterApi.GET("get-filter-user", ginUpNP(controllers.GetFilterUserList))
	twitterApi.POST("set-filter-user", ginUpP(controllers.SetFilterUser))
	twitterApi.POST("delete-filter-user", ginUpP(controllers.DeleteFilterUser))
	twitterApi.POST("get-twitter-tweet-list", ginUpP(controllers.GetTwitterTweetList))
	twitterApi.POST("get-mix-list", ginUpP(controllers.GetMixList))
	twitterApi.POST("get-twitter-user-list", ginUpP(controllers.GetTwitterUserList))
	twitterApi.POST("get-spider-twitter-his", ginUpP(controllers.GetSpiderTwitterHis))
	twitterApi.GET("run-spider-twitter-master", ginUpNP(controllers.RunSpiderTwitterMaster))
	twitterApi.GET("get-queue-len", ginUpNP(controllers.GetQueueLen))
}
