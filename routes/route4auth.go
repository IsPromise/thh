package routes

import (
	"github.com/gin-gonic/gin"
	"thh/app/http/controllers"
	"thh/app/http/middleware"
)

func ginAuth(ginApp *gin.Engine) {
	ginApp.Group("api").
		POST("reg", ginUpJP(controllers.Register)).
		POST("login", ginUpJP(controllers.Login))

	ginApp.Group("api").Use(middleware.JWTAuth4Gin).
		GET("get-user-info-v4", UpButterReq(controllers.UserInfoV4))
}
