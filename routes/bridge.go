package routes

import (
	"thh/app/http/controllers"
	"thh/app/http/middleware"

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
