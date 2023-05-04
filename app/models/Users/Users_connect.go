package Users

import (
	"gorm.io/gorm"
	db "thh/bundles/connect/dbconnect"
)

// Prohibit manual changes
// 禁止手动更改本文件

func builder() *gorm.DB {
	return db.Std().Table(tableName)
}

func first(db *gorm.DB) (el Users) {
	db.First(&el)
	return
}

func List(db *gorm.DB) (el []*Users) {
	db.Find(&el)
	return
}
