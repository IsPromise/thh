package voicechat

import (
	"github.com/gorilla/websocket"
	"log"
)

// clients holds the connected WebSocket clients
var clients = make(map[*websocket.Conn]bool)

// broadcast is a channel to send messages to all connected clients
var broadcast = make(chan []byte)

func GinVoiceChat(conn *websocket.Conn) {
	defer func() {
		_ = conn.Close()
	}()

	clients[conn] = true

	for {
		_, messageBytes, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Failed to read message: %v", err)
			delete(clients, conn)
			conn.Close()
			break
		}

		broadcast <- messageBytes
	}
}

// BroadcastWebSocket broadcasts messages to all connected clients
func BroadcastWebSocket() {
	for {
		messageBytes := <-broadcast

		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, messageBytes)
			if err != nil {
				log.Printf("Failed to send message to client: %v", err)
				delete(clients, client)
				client.Close()
			}
		}
	}
}

func init() {
	go BroadcastWebSocket()
}
