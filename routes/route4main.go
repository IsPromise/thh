package routes

import (
	"thh/app/http/controllers"
	"thh/app/http/controllers/ginLowerControllers"
	"thh/app/http/controllers/lowerControllers"
	"thh/app/http/middleware"
	"thh/arms"
	"thh/arms/app"

	"github.com/gin-gonic/gin"
)

func ginWeb(ginApp *gin.Engine) {
	if app.IsProduction() {
		ginApp.StaticFS("/actor", PFilSystem("./actor/dist", app.GetActorFS()))
	} else {
		ginApp.Static("/actor", "./actor/dist")
	}
	ginApp.GET("get-clash-config", lowerControllers.GinGetClashConfig)
	ginApp.GET("get-clash-config-plus", lowerControllers.GinGetClashConfigPlus)
}
func ginWs(ginApp *gin.Engine) {
	ginApp.GET("ws", middleware.WebSocketMid(ginLowerControllers.GinIm))
	arms.GuardGoRoutine(ginLowerControllers.Broadcaster)
}

func ginApi(ginApp *gin.Engine) {
	ginApp.GET("/api", controllers.Api)

	apiGroup := ginApp.Group("api")
	// lowerControllers
	apiGroup.GET("gin-upload", ginLowerControllers.GinUpload)
	apiGroup.GET("show-pic", ginLowerControllers.GinShowPic)
	apiGroup.POST("t-list", ginUpP(controllers.TListV2))
	apiGroup.GET("get-twitter-user-list", ginUpP(controllers.GetTwitterUserList))
	apiGroup.GET("get-twitter-tweet-list", ginUpP(controllers.GetTwitterTweetList))
	apiGroup.GET("get-tspider-his", ginUpP(controllers.GetTSpiderHis))
	//store := persistence.NewInMemoryStore(time.Second)
	//apiGroup.GET("/GetTwitterTweetList",
	//	cache.CachePage(store, time.Minute, ginUpP(controllers.GetTwitterTweetList)),
	//)
	apiGroup.GET("run-tspider-master", ginUpNP(controllers.RunTSpiderMaster))
	apiGroup.GET("get-queue-len", ginUpNP(controllers.GetQueueLen))
	apiGroup.GET("memUse", ginUpNP(controllers.GetUseMem))
	apiGroup.GET("about", ginUpNP(controllers.About))
	apiGroup.GET("sys-info", ginUpNP(controllers.SysInfo))
	apiGroup.GET("traefik-provider", ginLowerControllers.TraefikProvider)
	apiGroup.Any("test-bind", ginUpP(controllers.Params))

	apiGroup.Any("get-articles", ginUpP(controllers.GetArticles))
	apiGroup.Any("get-articles-detail", ginUpP(controllers.GetArticlesDetail))
}
