package PhoneLocation

import (
	db "thh/bundles/dbconnect"

	"gorm.io/gorm"
)

// Prohibit manual changes
// 禁止手动更改本文件

func builder() *gorm.DB {
	return db.Std().Table(tableName)
}

func First(db *gorm.DB) (el PhoneLocation) {
	db.First(&el)
	return
}

func List(db *gorm.DB) (el []PhoneLocation) {
	db.Find(&el)
	return
}
