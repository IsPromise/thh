package routes

import (
	"github.com/gin-gonic/gin"
	"thh/app/http/controllers/im"
	"thh/app/http/controllers/im/messagechat"
	"thh/app/http/controllers/im/voicechat"
	"thh/app/http/middleware"
	"thh/arms"
)

func ginWs(ginApp *gin.Engine) {
	ginApp.GET("api/ws-info", ginUpNP(im.ImInfo))

	ginApp.GET("ws", middleware.WebSocketMid(messagechat.GinIm))
	arms.GuardGoRoutine(messagechat.Broadcaster)

	ginApp.GET("ws-vc", middleware.WebSocketMid(voicechat.GinVoiceChat))
	arms.GuardGoRoutine(voicechat.BroadcastWebSocket)
}
