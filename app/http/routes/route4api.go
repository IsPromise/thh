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
	apiGroup.POST("file-upload", ginLowerControllers.GinUpload)
	// localhost:90/api/file?filename=storage/upload/2023/05/23/1684836490.png
	ginApp.Static("file/storage", "./storage")
	apiGroup.GET("show-pic", ginLowerControllers.GinShowPic)

	apiGroup.GET("traefik-provider", ginLowerControllers.TraefikProvider)
	apiGroup.GET("memUse", ginUpNP(controllers.GetUseMem))
	apiGroup.GET("about", ginUpNP(controllers.About))
	apiGroup.GET("sys-info", ginUpNP(controllers.SysInfo))
	apiGroup.GET("git-status-list", middleware.IpLimit, ginUpNP(controllers.GitStatusList))

	twitterApi := apiGroup.Group("twitter")
	twitterApi.Use(middleware.IpLimit)
	twitterApi.GET("get-filter-user", ginUpNP(controllers.GetFilterUserList))
	twitterApi.POST("set-filter-user", ginUpP(controllers.SetFilterUser))
	twitterApi.POST("delete-filter-user", ginUpP(controllers.DeleteFilterUser))
	twitterApi.POST("get-twitter-tweet-list", ginUpP(controllers.GetTwitterTweetList))
	twitterApi.POST("get-mix-list", ginUpP(controllers.GetMixList))
	twitterApi.POST("get-twitter-user-list", ginUpP(controllers.GetTwitterUserList))
	twitterApi.POST("get-spider-twitter-his", ginUpP(controllers.GetSpiderTwitterHis))
	twitterApi.GET("run-spider-twitter-master", ginUpNP(controllers.RunSpiderTwitterMaster))
	twitterApi.GET("get-queue-len", ginUpNP(controllers.GetQueueLen))

	apiGroup.Group("todo-task").
		GET("status-list", ginUpNP(controllers.TodoStatusList)).
		POST("create", ginUpP(controllers.CreateTask)).
		POST("update", ginUpP(controllers.UpdateTask)).
		GET("list", ginUpP(controllers.FindTodoList))
}
