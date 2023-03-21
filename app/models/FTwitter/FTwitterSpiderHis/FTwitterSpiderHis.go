package FTwitterSpiderHis

import (
	"time"
)

const tableName = "f_twitter_spider_his"
const pid = "id"
const fieldTarget = "target"
const fieldCurl = "curl"
const fieldSuccess = "success"
const fieldContent = "content"
const fieldType = "type"
const fieldCreateTime = "create_time"

type FTwitterSpiderHis struct {
	Id         uint64    `gorm:"primaryKey;column:id;autoIncrement;not null;" json:"id"`                                    // 主键
	Type       int       `gorm:"column:type;index;type:tinyint;not null;default:0;" json:"type"`                            // 类型
	Target     string    `gorm:"column:target;index;type:varchar(255);not null;default:screenName_userinfo;" json:"target"` // 请求参数
	Curl       string    `gorm:"column:curl;type:text;" json:"curl"`                                                        // 请求参数
	Success    string    `gorm:"column:success;index;type:varchar(255);not null;default:1;" json:"success"`                 // 请求参数
	Content    string    `gorm:"column:content;type:text;" json:"content"`                                                  // 返回内容
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP;" json:"createTime"`    //
}

// func (itself *FTwitterSpiderHis) BeforeSave(tx *gorm.DB) (err error) {}
// func (itself *FTwitterSpiderHis) BeforeCreate(tx *gorm.DB) (err error) {}
// func (itself *FTwitterSpiderHis) AfterCreate(tx *gorm.DB) (err error) {}
// func (itself *FTwitterSpiderHis) BeforeUpdate(tx *gorm.DB) (err error) {}
// func (itself *FTwitterSpiderHis) AfterUpdate(tx *gorm.DB) (err error) {}
// func (itself *FTwitterSpiderHis) AfterSave(tx *gorm.DB) (err error) {}
// func (itself *FTwitterSpiderHis) BeforeDelete(tx *gorm.DB) (err error) {}
// func (itself *FTwitterSpiderHis) AfterDelete(tx *gorm.DB) (err error) {}
// func (itself *FTwitterSpiderHis) AfterFind(tx *gorm.DB) (err error) {}

func (FTwitterSpiderHis) TableName() string {
	return tableName
}
