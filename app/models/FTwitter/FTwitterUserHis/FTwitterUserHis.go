package FTwitterUserHis

import (
	"time"
)

const tableName = "f_twitter_user_his"
const pid = "id"
const fieldRestId = "rest_id"
const fieldScreenName = "screen_name"
const fieldName = "name"
const fieldDesc = "desc"
const fieldCreateTime = "create_time"

type FTwitterUserHis struct {
	Id         uint64    `gorm:"primaryKey;column:id;autoIncrement;not null;" json:"id"`                                 // 主键
	RestId     string    `gorm:"column:rest_id;index;type:varchar(255);not null;default:0;" json:"rest_id"`              // 用户id
	ScreenName string    `gorm:"column:screen_name;index;type:varchar(255);not null;default:0;" json:"userId"`           // 用户id
	Name       string    `gorm:"column:name;index;type:varchar(255);not null;default:'';" json:"name"`                   // 用户id
	Desc       string    `gorm:"column:desc;type:text;" json:"desc"`                                                     //
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP;" json:"createTime"` //

}

// func (itself *FTwitterUserHis) BeforeSave(tx *gorm.DB) (err error) {}
// func (itself *FTwitterUserHis) BeforeCreate(tx *gorm.DB) (err error) {}
// func (itself *FTwitterUserHis) AfterCreate(tx *gorm.DB) (err error) {}
// func (itself *FTwitterUserHis) BeforeUpdate(tx *gorm.DB) (err error) {}
// func (itself *FTwitterUserHis) AfterUpdate(tx *gorm.DB) (err error) {}
// func (itself *FTwitterUserHis) AfterSave(tx *gorm.DB) (err error) {}
// func (itself *FTwitterUserHis) BeforeDelete(tx *gorm.DB) (err error) {}
// func (itself *FTwitterUserHis) AfterDelete(tx *gorm.DB) (err error) {}
// func (itself *FTwitterUserHis) AfterFind(tx *gorm.DB) (err error) {}

func (FTwitterUserHis) TableName() string {
	return tableName
}
