package FTwitterMedia

import (
	db "thh/bundles/connect/dbconnect"

	"gorm.io/gorm"
)

// Prohibit manual changes
// 禁止手动更改本文件

func builder() *gorm.DB {
	return db.Std().Table(tableName)
}

func first(db *gorm.DB) (el FTwitterMedia) {
	db.First(&el)
	return
}

func List(db *gorm.DB) (el []*FTwitterMedia) {
	db.Find(&el)
	return
}
