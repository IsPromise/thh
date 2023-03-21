package lowerControllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"thh/arms"
	"thh/arms/logger"

	"github.com/gofiber/websocket/v2"
	"github.com/spf13/cast"
)

var idm = arms.IdMakerInOnP{}

func Im(ws *websocket.Conn) {
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
		logger.Info("user:", client.name, "websocket connect success!")
	}
	for {
		if msgType, msg, err = ws.ReadMessage(); err != nil {
			log.Println("read:", err, msgType)
			leave <- client
			break
		}
		log.Printf("recv: %s", msg)
		message <- Message{0, client.name, string(msg)}

		//if err = ws.WriteMessage(mt, msg); err != nil {
		//	log.Println("write:", err)
		//	break
		//}
	}
}

type Client struct {
	conn     *websocket.Conn // 用户websocket连接
	name     string          // 用户名称
	clientId string          // 客户端唯一id
}

func (client *Client) SendMessage(messageType int, data []byte) (err error) {
	defer func() {
		if info := recover(); info != nil {
			fmt.Println(info)
			err = errors.New(cast.ToString(info))
			return
		}
	}()
	err = client.conn.WriteMessage(websocket.TextMessage, data)
	return
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

// 此处要设置有缓冲的通道。因为这是goroutine自己从通道中发送并接受数据。
// 若是无缓冲的通道，该goroutine发送数据到通道后就被锁定，需要数据被接受后才能解锁，而恰恰接受数据的又只能是它自己
var join = make(chan Client, 10)     // 用户加入通道
var leave = make(chan Client, 10)    // 用户退出通道
var message = make(chan Message, 10) // 消息通道
func Broadcaster() {
	for {
		// 哪个case可以执行，则转入到该case。若都不可执行，则堵塞。
		select {
		// 消息通道中有消息则执行，否则堵塞
		case msg := <-message:
			str := fmt.Sprintf("broadcaster-----------%s send message: %s\n", msg.Name, msg.Message)
			logger.Info(str)
			// 将某个用户发出的消息发送给所有用户
			for client := range clients {
				// 将数据编码成json形式，data是[]byte类型
				// json.Marshal()只会编码结构体中公开的属性(即大写字母开头的属性)
				data, err := json.Marshal(msg)
				if err != nil {
					logger.Info("Fail to marshal message:", err)
					return
				}
				//
				if err = client.SendMessage(websocket.TextMessage, data); err != nil {
					delete(clients, client)
					logger.Info("Fail to write message")
				}
			}

		// 有用户加入
		case client := <-join:
			logger.Info(fmt.Sprintf("broadcaster-----------%s join in the chat room\n", client.name))

			clients[client] = true // 将用户加入映射

			message <- Message{1, client.name, fmt.Sprintf("%s join in, there are %d preson in room", client.name, len(clients))}

		// 有用户退出
		case client := <-leave:
			logger.Info(fmt.Sprintf("broadcaster-----------%s leave the chat room\n", client.name))

			// 如果该用户已经被删除
			if !clients[client] {
				logger.Info("the client had leaved, client's name:" + client.name)
				break
			}

			delete(clients, client) // 将用户从映射中删除

			// 将用户退出消息放入消息通道
			message <- Message{2, client.name, fmt.Sprintf("%s leave, there are %d preson in room", client.name, len(clients))}
		}
	}
}
