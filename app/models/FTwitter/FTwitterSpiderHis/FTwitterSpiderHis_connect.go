package FTwitterSpiderHis

import (
	"gorm.io/gorm"
	db "thh/app/bundles/connect/dbconnect"
)

// Prohibit manual changes
// 禁止手动更改本文件

func builder() *gorm.DB {
	return db.Std().Table(tableName)
}

func first(db *gorm.DB) (el FTwitterSpiderHis) {
	db.First(&el)
	return
}

func List(db *gorm.DB) (el []*FTwitterSpiderHis) {
	db.Find(&el)
	return
}
