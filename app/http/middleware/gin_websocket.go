package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSocketMid(handler func(*websocket.Conn)) func(c *gin.Context) {
	return func(c *gin.Context) {
		ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			_, err = c.Writer.Write([]byte(err.Error()))
			if err != nil {
				return
			}
			return
		}
		handler(ws)
	}
}
