package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"thh/app/http/controllers/component"
)

func gincq(ginApp *gin.Engine) {
	cqGroup := ginApp.Group("/api/cq")
	cqGroup.POST("msgv2", ginUpP(CqMsg))

}

type MsgData struct {
	Time        int    `json:"time"`
	SelfId      int    `json:"self_id"`
	PostType    string `json:"post_type"`
	MessageType string `json:"message_type"`
	SubType     string `json:"sub_type"`
	MessageId   int    `json:"message_id"`
	UserId      int    `json:"user_id"`
	Message     string `json:"message"`
	RawMessage  string `json:"raw_message"`
	Font        int    `json:"font"`
	Sender      struct {
		Nickname string `json:"nickname"`
		Sex      string `json:"sex"`
		Age      int    `json:"age"`
	} `json:"sender"`
}

type MsgReply struct {
	Reply string `json:"reply"`
}

func CqMsg(data MsgData) component.Response {
	fmt.Println(data)
	return component.Data(MsgReply{
		//Reply: "ok",
	})
}
