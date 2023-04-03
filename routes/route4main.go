package routes

import (
	"github.com/gin-contrib/gzip"
	"thh/app/http/controllers/ginLowerControllers"
	"thh/arms/app"

	"github.com/gin-gonic/gin"
)

func ginWeb(ginApp *gin.Engine) {
	if app.IsProduction() {
		ginApp.Use(gzip.Gzip(gzip.DefaultCompression)).StaticFS("/actor", PFilSystem("./actor/dist", app.GetActorFS()))
	} else {
		ginApp.Use(gzip.Gzip(gzip.DefaultCompression)).Static("/actor", "./actor/dist")
	}
	ginApp.GET("get-clash-config", ginLowerControllers.GinGetClashConfig)
	ginApp.GET("get-clash-config-plus", ginLowerControllers.GinGetClashConfigPlus)
}
