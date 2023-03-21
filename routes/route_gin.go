package routes

import (
	"thh/app/http/controllers"
	"thh/app/http/controllers/genLowerControllers"
	"thh/app/http/controllers/lowerControllers"
	"thh/app/http/middleware"
	"thh/arms"
	"thh/arms/app"

	"github.com/gin-gonic/gin"
)

func RegisterByGin(ginApp *gin.Engine) {

	ginApp.Use(middleware.TraceInit)
	ginApp.Use(middleware.GinCors)
	ginApp.Use(middleware.GinLogger)

	ginWeb(ginApp)
	ginApi(ginApp)
	ginAuth(ginApp)
	ginWs(ginApp)

	ginApp.NoRoute(controllers.NotFound)
}

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
	ginApp.GET("ws", middleware.WebSocketMid(genLowerControllers.GinIm))
	arms.GuardGoRoutine(genLowerControllers.Broadcaster)
}
func ginApi(ginApp *gin.Engine) {
	ginApp.GET("/api", controllers.Api)

	apiGroup := ginApp.Group("api")
	apiGroup.POST("reg", ginUpJP(controllers.Register))
	apiGroup.GET("login", ginUpJP(controllers.Login))
	// lowerControllers
	apiGroup.GET("gin-upload", genLowerControllers.GinUpload)
	apiGroup.GET("show-pic", genLowerControllers.GinShowPic)
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
	apiGroup.GET("traefik-provider", genLowerControllers.TraefikProvider)
	apiGroup.Any("test-bind", ginUpP(controllers.Params))

	apiGroup.Any("get-articles", ginUpP(controllers.GetArticles))
	apiGroup.Any("get-articles-detail", ginUpP(controllers.GetArticlesDetail))
}

func ginAuth(ginApp *gin.Engine) {
	authGroup := ginApp.Group("api").Use(middleware.JWTAuth4Gin)
	authGroup.GET("get-user-info", ginUpNPAuth(controllers.UserInfoV3))
}
