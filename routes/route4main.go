package routes

import (
	"thh/app/http/controllers/ginLowerControllers"
	"thh/app/http/middleware"
	"thh/bundles/kernel"

	"github.com/gin-contrib/gzip"

	"github.com/gin-gonic/gin"
)

func ginWeb(ginApp *gin.Engine) {
	actGroup := ginApp.Group("/actor")
	if kernel.IsProduction() {
		actGroup.
			Use(middleware.BrowserCache).
			Use(middleware.CacheMiddleware).
			Use(gzip.Gzip(gzip.DefaultCompression)).
			StaticFS("", PFilSystem("./actor/dist", kernel.GetActorFS()))
	} else {
		actGroup.
			Use(gzip.Gzip(gzip.DefaultCompression)).
			Static("", "./actor/dist")
	}
	ginApp.GET("get-clash-config", ginLowerControllers.GinGetClashConfig)
	ginApp.GET("get-clash-config-plus", ginLowerControllers.GinGetClashConfigPlus)
}
