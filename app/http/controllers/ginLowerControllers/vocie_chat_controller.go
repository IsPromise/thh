package ginLowerControllers

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true // 允许跨域请求
		},
	}

	voiceChatClients = make(map[*VoiceChatClient]bool)
	broadcast        = make(chan []byte)
)

func GinVoiceChat(ws *websocket.Conn) {
	defer func() {
		_ = ws.Close()
	}()

	client := &VoiceChatClient{conn: ws, send: make(chan []byte)}
	voiceChatClients[client] = true

	go client.read()
	go client.write()

	log.Println("VoiceChatClient connected")
}

type VoiceChatClient struct {
	conn *websocket.Conn
	send chan []byte
}

func (c *VoiceChatClient) read() {
	defer func() {
		c.conn.Close()
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
		log.Printf("Received message: %s", message)
		// 将收到的消息广播给所有其他客户端
		broadcast <- message
	}
}

func (c *VoiceChatClient) write() {
	defer func() {
		c.conn.Close()
	}()

	for message := range c.send {
		err := c.conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			break
		}
	}
}

func broadcastLoop() {
	for {
		message := <-broadcast
		log.Printf("Broadcast message: %s", message)

		for client := range voiceChatClients {
			select {
			case client.send <- message:
			default:
				close(client.send)
				delete(voiceChatClients, client)
			}
		}
	}
}

func init() {
	go broadcastLoop()
}
