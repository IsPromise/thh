package Users

import (
	"gorm.io/gorm"
	db "thh/app/bundles/connect/dbconnect"
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

func list(db *gorm.DB) (el []*Users) {
	db.Find(&el)
	return
}

func builderWithOutTable() *gorm.DB {
	return db.Std()
}
