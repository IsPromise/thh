package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许跨域请求
	},
}

func WebSocketMid(handler func(*websocket.Conn)) func(c *gin.Context) {
	return func(c *gin.Context) {
		ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println(err)
			_, err = c.Writer.Write([]byte(err.Error()))
			if err != nil {
				fmt.Println(err)
				return
			}
			return
		}
		handler(ws)
	}
}
