package Users

import (
	"gorm.io/gorm"
	"thh/app/bundles/algorithm"
	"time"
)

const tableName = "users"
const pid = "id"
const fieldCreatedAt = "created_at"
const fieldUpdatedAt = "updated_at"
const fieldDeletedAt = "deleted_at"
const fieldUsername = "username"
const fieldEmail = "email"
const fieldPassword = "password"

type Users struct {
	Id        uint64    `gorm:"primaryKey;column:id;autoIncrement;not null;" json:"id"`                             //
	Verify    int       `gorm:"column:verify;type:tinyint(4);not null;default:0" json:"verify"`                     // 验证
	Username  string    `gorm:"column:username;type:varchar(255);uniqueIndex;not null;default:'';" json:"username"` //
	Email     string    `gorm:"column:email;type:varchar(255);uniqueIndex;not null;default:'';" json:"email"`       //
	Password  string    `gorm:"column:password;type:varchar(255);not null;default:'';" json:"-"`                    //
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;type:datetime;" json:"createdAt"`                   //
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime;type:datetime;" json:"updatedAt"`                   //
	DeletedAt gorm.DeletedAt
	// *time.Time `gorm:"column:deleted_at;type:datetime;" json:"-"`                                          //
}

// func (itself *Users) BeforeSave(tx *gorm.DB) (err error) {}
// func (itself *Users) BeforeCreate(tx *gorm.DB) (err error) {}
// func (itself *Users) AfterCreate(tx *gorm.DB) (err error) {}
// func (itself *Users) BeforeUpdate(tx *gorm.DB) (err error) {}
// func (itself *Users) AfterUpdate(tx *gorm.DB) (err error) {}
// func (itself *Users) AfterSave(tx *gorm.DB) (err error) {}
// func (itself *Users) BeforeDelete(tx *gorm.DB) (err error) {}
// func (itself *Users) AfterDelete(tx *gorm.DB) (err error) {}
// func (itself *Users) AfterFind(tx *gorm.DB) (err error) {}

func (itself *Users) TableName() string {
	return tableName
}

func (itself *Users) SetPassword(password string) *Users {
	itself.Password, _ = algorithm.MakePassword(password)
	return itself
}
