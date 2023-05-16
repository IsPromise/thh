package FTwitterFilterUser

import (
	"gorm.io/gorm"
	"time"
)

const tableName = "f_twitter_filter_user"
const pid = "id"
const fieldScreenName = "screen_name"
const fieldCreateTime = "create_time"
const fieldDeletedAt = "deleted_at"

type FTwitterFilterUser struct {
	Id         uint64    `gorm:"primaryKey;column:id;autoIncrement;not null;" json:"id"`                                 // 主键
	ScreenName string    `gorm:"column:screen_name;type:varchar(255);not null;" json:"screenName"`                       //
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP;" json:"createTime"` //
	DeletedAt  gorm.DeletedAt
}

// func (itself *FTwitterFilterUser) BeforeSave(tx *gorm.DB) (err error) {}
// func (itself *FTwitterFilterUser) BeforeCreate(tx *gorm.DB) (err error) {}
// func (itself *FTwitterFilterUser) AfterCreate(tx *gorm.DB) (err error) {}
// func (itself *FTwitterFilterUser) BeforeUpdate(tx *gorm.DB) (err error) {}
// func (itself *FTwitterFilterUser) AfterUpdate(tx *gorm.DB) (err error) {}
// func (itself *FTwitterFilterUser) AfterSave(tx *gorm.DB) (err error) {}
// func (itself *FTwitterFilterUser) BeforeDelete(tx *gorm.DB) (err error) {}
// func (itself *FTwitterFilterUser) AfterDelete(tx *gorm.DB) (err error) {}
// func (itself *FTwitterFilterUser) AfterFind(tx *gorm.DB) (err error) {}

func (FTwitterFilterUser) TableName() string {
	return tableName
}
