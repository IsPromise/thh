package routes

import (
	"github.com/gin-gonic/gin"
	"thh/app/http/controllers"
	"thh/app/http/middleware"
)

func ginAuth(ginApp *gin.Engine) {
	authGroup := ginApp.Group("api").Use(middleware.JWTAuth4Gin)
	authGroup.GET("get-user-info", ginUpNPAuth(controllers.UserInfoV3))
}
