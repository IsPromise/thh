package routes

import (
	"thh/app/http/controllers"
	"thh/app/http/middleware"

	"github.com/gin-gonic/gin"
)

func ginAuth(ginApp *gin.Engine) {
	ginApp.Group("api").
		POST("reg", ginUpJP(controllers.Register)).
		POST("login", ginUpJP(controllers.Login))

	ginApp.Group("api").Use(middleware.JWTAuth4Gin).
		GET("get-user-info-v4", UpButterReq(controllers.UserInfoV4)).
		POST("set-user-info", UpButterReq(controllers.EditUserInfo))
}
