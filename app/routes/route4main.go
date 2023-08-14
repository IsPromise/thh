package routes

import (
	"github.com/gin-contrib/gzip"
	kernel2 "thh/app/bundles/kernel"
	"thh/app/http/controllers/ginLowerControllers"
	"thh/app/http/middleware"

	"github.com/gin-gonic/gin"
)

func ginWeb(ginApp *gin.Engine) {
	actGroup := ginApp.Group("/actor")
	if kernel2.IsProduction() {
		actGroup.
			Use(middleware.BrowserCache).
			Use(middleware.CacheMiddleware).
			Use(gzip.Gzip(gzip.DefaultCompression)).
			StaticFS("", PFilSystem("./actor/dist", kernel2.GetActorFS()))
	} else {
		actGroup.
			Use(gzip.Gzip(gzip.DefaultCompression)).
			Static("", "./actor/dist")
	}
	ginApp.GET("get-clash-config", ginLowerControllers.GinGetClashConfig)
	ginApp.GET("get-clash-config-plus", ginLowerControllers.GinGetClashConfigPlus)
}
