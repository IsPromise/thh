package FTwitterUser

import (
	"time"
)

const tableName = "f_twitter_user"
const pid = "id"
const fieldRestId = "rest_id"
const fieldScreenName = "screen_name"
const fieldName = "name"
const fieldDesc = "desc"
const fieldCreateTime = "create_time"

type FTwitterUser struct {
	Id         uint64    `gorm:"primaryKey;column:id;autoIncrement;not null;" json:"id"`                                 // 主键
	RestId     string    `gorm:"column:rest_id;index;type:varchar(255);not null;default:0;" json:"rest_id"`              // 用户id
	ScreenName string    `gorm:"column:screen_name;index;type:varchar(255);not null;default:'';" json:"screenName"`      // 用户id
	Name       string    `gorm:"column:name;index;type:varchar(255);not null;default:'';" json:"name"`                   // 用户id
	Desc       string    `gorm:"column:desc;type:text;" json:"desc"`                                                     //
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP;" json:"createTime"` //

}

// func (itself *FTwitterUser) BeforeSave(tx *gorm.DB) (err error) {}
// func (itself *FTwitterUser) BeforeCreate(tx *gorm.DB) (err error) {}
// func (itself *FTwitterUser) AfterCreate(tx *gorm.DB) (err error) {}
// func (itself *FTwitterUser) BeforeUpdate(tx *gorm.DB) (err error) {}
// func (itself *FTwitterUser) AfterUpdate(tx *gorm.DB) (err error) {}
// func (itself *FTwitterUser) AfterSave(tx *gorm.DB) (err error) {}
// func (itself *FTwitterUser) BeforeDelete(tx *gorm.DB) (err error) {}
// func (itself *FTwitterUser) AfterDelete(tx *gorm.DB) (err error) {}
// func (itself *FTwitterUser) AfterFind(tx *gorm.DB) (err error) {}

func (FTwitterUser) TableName() string {
	return tableName
}
