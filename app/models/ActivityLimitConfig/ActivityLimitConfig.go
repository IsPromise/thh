package ActivityLimitConfig

import (
	"time"
)

const tableName = "activity_limit_config"
const pid = "id"
const fieldActivityId = "activity_id"
const fieldType = "type"
const fieldTypeId = "type_id"
const fieldIsDeleted = "is_deleted"
const fieldCreatedAt = "created_at"
const fieldUpdatedAt = "updated_at"

type ActivityLimitConfig struct {
	Id         uint64    `gorm:"primaryKey;column:id;autoIncrement;not null;" json:"id"`                                                     //
	ActivityId uint      `gorm:"column:activity_id;type:int unsigned;not null;default:1;" json:"activityId"`                                 //
	Type       string    `gorm:"column:type;type:varchar(50);not null;default:'';" json:"type"`                                              //
	TypeId     string    `gorm:"column:type_id;type:varchar(50);not null;default:'';" json:"typeId"`                                         //
	IsDeleted  int       `gorm:"column:is_deleted;type:tinyint(1);not null;default:0;" json:"isDeleted"`                                     //
	CreatedAt  time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;" json:"createdAt"`                      //
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime:true;type:datetime;not null;default:CURRENT_TIMESTAMP;" json:"updateTime"` //

}

// func (itself *ActivityLimitConfig) BeforeSave(tx *gorm.DB) (err error) {}
// func (itself *ActivityLimitConfig) BeforeCreate(tx *gorm.DB) (err error) {}
// func (itself *ActivityLimitConfig) AfterCreate(tx *gorm.DB) (err error) {}
// func (itself *ActivityLimitConfig) BeforeUpdate(tx *gorm.DB) (err error) {}
// func (itself *ActivityLimitConfig) AfterUpdate(tx *gorm.DB) (err error) {}
// func (itself *ActivityLimitConfig) AfterSave(tx *gorm.DB) (err error) {}
// func (itself *ActivityLimitConfig) BeforeDelete(tx *gorm.DB) (err error) {}
// func (itself *ActivityLimitConfig) AfterDelete(tx *gorm.DB) (err error) {}
// func (itself *ActivityLimitConfig) AfterFind(tx *gorm.DB) (err error) {}

func (ActivityLimitConfig) TableName() string {
	return tableName
}
