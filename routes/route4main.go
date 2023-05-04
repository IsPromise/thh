package routes

import (
	"thh/app/http/controllers/ginLowerControllers"
	app2 "thh/bundles/app"

	"github.com/gin-contrib/gzip"

	"github.com/gin-gonic/gin"
)

func ginWeb(ginApp *gin.Engine) {
	if app2.IsProduction() {
		ginApp.Use(gzip.Gzip(gzip.DefaultCompression)).StaticFS("/actor", PFilSystem("./actor/dist", app2.GetActorFS()))
	} else {
		ginApp.Use(gzip.Gzip(gzip.DefaultCompression)).Static("/actor", "./actor/dist")
	}
	ginApp.GET("get-clash-config", ginLowerControllers.GinGetClashConfig)
	ginApp.GET("get-clash-config-plus", ginLowerControllers.GinGetClashConfigPlus)
}
