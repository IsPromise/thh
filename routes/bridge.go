package routes

import (
	"github.com/gin-gonic/gin"
	"thh/app/http/controllers"
	"thh/app/http/middleware"
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
