package FTwitterMedia

import (
	"time"
)

const tableName = "f_twitter_media"
const pid = "id"
const fieldTweetId = "tweet_id"
const fieldType = "type"
const fieldPath = "path"
const fieldUrl = "url"
const fieldCreateTime = "create_time"

type FTwitterMedia struct {
	Id         uint64    `gorm:"primaryKey;column:id;autoIncrement;not null;" json:"id"`                                 // 主键
	TweetId    string    `gorm:"column:tweet_id;index;type:varchar(255);not null;default:0;" json:"tweetId"`             //
	Type       string    `gorm:"column:type;index;type:varchar(255);not null;default:pic;" json:"type"`                  //
	Path       string    `gorm:"column:path;index(255);type:varchar(1024);default:'';" json:"path"`                      //
	Url        string    `gorm:"column:url;index(255);type:varchar(1024);default:'';" json:"url"`                        //
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP;" json:"createTime"` //

}

// func (itself *FTwitterMedia) BeforeSave(tx *gorm.DB) (err error) {}
// func (itself *FTwitterMedia) BeforeCreate(tx *gorm.DB) (err error) {}
// func (itself *FTwitterMedia) AfterCreate(tx *gorm.DB) (err error) {}
// func (itself *FTwitterMedia) BeforeUpdate(tx *gorm.DB) (err error) {}
// func (itself *FTwitterMedia) AfterUpdate(tx *gorm.DB) (err error) {}
// func (itself *FTwitterMedia) AfterSave(tx *gorm.DB) (err error) {}
// func (itself *FTwitterMedia) BeforeDelete(tx *gorm.DB) (err error) {}
// func (itself *FTwitterMedia) AfterDelete(tx *gorm.DB) (err error) {}
// func (itself *FTwitterMedia) AfterFind(tx *gorm.DB) (err error) {}

func (FTwitterMedia) TableName() string {
	return tableName
}
