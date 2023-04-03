package routes

import (
	"thh/app/http/controllers/im"
	"thh/app/http/controllers/im/messagechat"
	"thh/app/http/controllers/im/voicechat"
	"thh/app/http/middleware"
	"thh/arms"

	"github.com/gin-gonic/gin"
)

func ginWs(ginApp *gin.Engine) {
	ginApp.GET("api/ws-info", ginUpNP(im.ImInfo))

	ginApp.GET("ws", middleware.WebSocketMid(messagechat.GinIm))
	arms.GuardGoRoutine(messagechat.Broadcaster)

	ginApp.GET("ws-vc", middleware.WebSocketMid(voicechat.GinVoiceChat))
	arms.GuardGoRoutine(voicechat.BroadcastWebSocket)
}
