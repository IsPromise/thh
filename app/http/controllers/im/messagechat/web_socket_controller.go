package messagechat

import (
	"encoding/json"
	"fmt"
	"github.com/leancodebox/goose/luckrand"
	"log"
	"sync"
	"thh/bundles/logging"

	"github.com/gorilla/websocket"
	"github.com/spf13/cast"
)

var idm = &luckrand.IdMakerInOnP{}

func GinIm(ws *websocket.Conn) {
	defer func() {
		_ = ws.Close()
	}()
	var (
		msgType int
		msg     []byte
		err     error
	)

	var client Client
	client.name = "t"
	client.conn = ws
	client.clientId = cast.ToString(idm.Get())

	if !clients[client] {
		join <- client
		logging.Info("user:", client.name, "websocket connect success!")
	}
	for {
		if msgType, msg, err = ws.ReadMessage(); err != nil {
			log.Println("read:", err, msgType)
			leave <- client
			break
		}
		log.Printf("recv: %s", msg)
		message <- Message{0, client.name, string(msg)}
	}
}

// onOpen
// onMessage
// onSend
// onClose

type Client struct {
	conn     *websocket.Conn // 用户websocket连接
	name     string          // 用户名称
	clientId string          // 客户端唯一id
}

// Message
// 1.设置为公开属性(即首字母大写)，是因为属性值私有时，外包的函数无法使用或访问该属性值(如：json.Marshal())
// 2.`json:"name"` 是为了在对该结构类型进行json编码时，自定义该属性的名称
type Message struct {
	EventType byte   `json:"type"`    // 0表示用户发布消息；1表示用户进入；2表示用户退出
	Name      string `json:"name"`    // 用户名称
	Message   string `json:"message"` // 消息内容
}

var clients = make(map[Client]bool) // 用户组映射
var clientsManager = &sync.Mutex{}

// 此处要设置有缓冲的通道。因为这是goroutine自己从通道中发送并接受数据。
// 若是无缓冲的通道，该goroutine发送数据到通道后就被锁定，需要数据被接受后才能解锁，而恰恰接受数据的又只能是它自己
var join = make(chan Client, 10)     // 用户加入通道
var leave = make(chan Client, 10)    // 用户退出通道
var message = make(chan Message, 10) // 消息通道

func joinList(client Client) {
	clientsManager.Lock()
	defer clientsManager.Unlock()
	clients[client] = true
}

func leaveList(client Client) {
	clientsManager.Lock()
	defer clientsManager.Unlock()
	delete(clients, client)
}

// Broadcaster 有新的客户端加入时群发消息，广播时发送失败则断开该客户端
func Broadcaster() {
	for {
		select {
		case msg := <-message:
			logging.Printf("broadcaster-----------%s send message: %s\n", msg.Name, msg.Message)
			for client := range clients {
				data, err := json.Marshal(msg)
				if logging.ErrIf(err) {
					continue
				}
				if logging.ErrIf(client.conn.WriteMessage(websocket.TextMessage, data)) {
					leaveList(client)
				}
			}
			break
		// 有用户加入
		case client := <-join:
			logging.Printf("broadcaster-----------%s join in the chat room\n", client.name)
			joinList(client)
			message <- Message{1, client.name, fmt.Sprintf("%s join in, there are %d preson in room", client.name, len(clients))}
			break
		// 有用户退出
		case client := <-leave:
			logging.Printf("broadcaster-----------%s leave the chat room\n", client.name)
			leaveList(client)
			// 将用户退出消息放入消息通道
			message <- Message{2, client.name, fmt.Sprintf("%s leave, there are %d preson in room", client.name, len(clients))}
			break
		}
	}
}
