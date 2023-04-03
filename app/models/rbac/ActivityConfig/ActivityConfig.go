package ActivityConfig

import (
	"time"
)

const tableName = "activity_config"
const pid = "id"
const fieldName = "name"
const fieldDetails = "details"
const fieldUrl = "url"
const fieldPicUrl = "pic_url"
const fieldDescUrl = "desc_url"
const fieldStartTime = "start_time"
const fieldEndTime = "end_time"
const fieldIsStop = "is_stop"
const fieldIsDeleted = "is_deleted"
const fieldCreatedAt = "created_at"
const fieldUpdatedAt = "updated_at"

type ActivityConfig struct {
	Id        uint64    `gorm:"primaryKey;column:id;autoIncrement;not null;" json:"id"`                                                    //
	Name      string    `gorm:"column:name;type:varchar(255);not null;default:活动名;" json:"name"`                                           //
	Details   string    `gorm:"column:details;type:varchar(1024);not null;default:活动内容;" json:"details"`                                   //
	Url       string    `gorm:"column:url;type:varchar(1024);not null;default:http://baidu.com;" json:"url"`                               //
	PicUrl    string    `gorm:"column:pic_url;type:varchar(1024);not null;default:http://baidu.com;" json:"picUrl"`                        //
	DescUrl   string    `gorm:"column:desc_url;type:varchar(1024);not null;default:http://baidu.com;" json:"descUrl"`                      //
	StartTime time.Time `gorm:"column:start_time;index;type:timestamp;not null;default:CURRENT_TIMESTAMP;" json:"startTime"`               //
	EndTime   time.Time `gorm:"column:end_time;index;type:timestamp;not null;default:CURRENT_TIMESTAMP;" json:"endTime"`                   //
	IsStop    int       `gorm:"column:is_stop;type:tinyint(1);not null;default:0;" json:"isStop"`                                          //
	IsDeleted int       `gorm:"column:is_deleted;type:tinyint(1);not null;default:0;" json:"isDeleted"`                                    //
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;" json:"createdAt"`                     //
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime:true;type:timestamp;not null;default:CURRENT_TIMESTAMP;" json:"updatedAt"` //

}

// func (itself *ActivityConfig) BeforeSave(tx *gorm.DB) (err error) {}
// func (itself *ActivityConfig) BeforeCreate(tx *gorm.DB) (err error) {}
// func (itself *ActivityConfig) AfterCreate(tx *gorm.DB) (err error) {}
// func (itself *ActivityConfig) BeforeUpdate(tx *gorm.DB) (err error) {}
// func (itself *ActivityConfig) AfterUpdate(tx *gorm.DB) (err error) {}
// func (itself *ActivityConfig) AfterSave(tx *gorm.DB) (err error) {}
// func (itself *ActivityConfig) BeforeDelete(tx *gorm.DB) (err error) {}
// func (itself *ActivityConfig) AfterDelete(tx *gorm.DB) (err error) {}
// func (itself *ActivityConfig) AfterFind(tx *gorm.DB) (err error) {}

func (ActivityConfig) TableName() string {
	return tableName
}
