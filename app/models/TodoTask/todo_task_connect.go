package TodoTask

import (
	"gorm.io/gorm"

	db "thh/bundles/connect/dbconnect"
)

// Prohibit manual changes
// 禁止手动更改本文件

func builder() *gorm.DB {
	return db.Std().Table(tableName)
}

func first(db *gorm.DB) (el Entity) {
	db.First(&el)
	return
}

func getList(db *gorm.DB) (el []*Entity) {
	db.Find(&el)
	return
}