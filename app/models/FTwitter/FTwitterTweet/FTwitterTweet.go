package FTwitterTweet

import (
	"time"
)

const tableName = "f_twitter_tweet"
const pid = "id"
const fieldScreenName = "screen_name"
const fieldConversationId = "conversation_id"
const fieldOriginScreenName = "origin_screen_name"
const fieldContext = "context"
const fieldCreateTime = "create_time"

type FTwitterTweet struct {
	Id               uint64    `gorm:"primaryKey;column:id;autoIncrement;not null;" json:"id"`                                                // 主键
	ScreenName       string    `gorm:"column:screen_name;index;type:varchar(255);not null;default:''" json:"screenName"`                      //
	ConversationId   string    `gorm:"column:conversation_id;index;type:varchar(255);not null;default:''" json:"conversationId"`              //
	OriginScreenName string    `gorm:"column:origin_screen_name;index;type:varchar(255);not null;default:''" json:"originScreenName"`         //
	Context          string    `gorm:"column:context;type:text;" json:"context"`                                                              //
	CreateTime       time.Time `gorm:"column:create_time;autoCreateTime;type:datetime;not null;default:CURRENT_TIMESTAMP;" json:"createTime"` //

}

// func (itself *FTwitterTweet) BeforeSave(tx *gorm.DB) (err error) {}
// func (itself *FTwitterTweet) BeforeCreate(tx *gorm.DB) (err error) {}
// func (itself *FTwitterTweet) AfterCreate(tx *gorm.DB) (err error) {}
// func (itself *FTwitterTweet) BeforeUpdate(tx *gorm.DB) (err error) {}
// func (itself *FTwitterTweet) AfterUpdate(tx *gorm.DB) (err error) {}
// func (itself *FTwitterTweet) AfterSave(tx *gorm.DB) (err error) {}
// func (itself *FTwitterTweet) BeforeDelete(tx *gorm.DB) (err error) {}
// func (itself *FTwitterTweet) AfterDelete(tx *gorm.DB) (err error) {}
// func (itself *FTwitterTweet) AfterFind(tx *gorm.DB) (err error) {}

func (FTwitterTweet) TableName() string {
	return tableName
}
