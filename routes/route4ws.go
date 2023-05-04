package routes

import (
	"thh/app/http/controllers/im"
	"thh/app/http/controllers/im/messagechat"
	"thh/app/http/controllers/im/voicechat"
	"thh/app/http/middleware"

	"github.com/gin-gonic/gin"
	"github.com/leancodebox/goose/power"
)

func ginWs(ginApp *gin.Engine) {
	ginApp.GET("api/ws-info", ginUpNP(im.ImInfo))

	ginApp.GET("ws", middleware.WebSocketMid(messagechat.GinIm))
	power.GuardGoRoutine(messagechat.Broadcaster)

	ginApp.GET("ws-vc", middleware.WebSocketMid(voicechat.GinVoiceChat))
	power.GuardGoRoutine(voicechat.BroadcastWebSocket)
}
