package migration

import (
	"fmt"
	"gorm.io/gorm"
	"thh/app/models/PhoneLocation"
	"thh/app/models/bbs/Articles"
	"thh/app/models/bbs/Comment"
	"thh/arms/app"
	"thh/arms/logger"
	"thh/bundles/dbconnect"
)

func M() {
	fmt.Println("init migration")
	// 数据库迁移
	migration(app.UseMigration(), dbconnect.Std())
}

func migration(migration bool, db *gorm.DB) {
	if migration == false {
		return
	}
	// 自动迁移
	var err error

	if err = db.AutoMigrate(
		&PhoneLocation.PhoneLocation{},
		&Comment.Comment{},
		&Articles.Articles{},
		//&DataReps.DataReps{},
		//&Users.Users{},
		//&Role.Role{},
		//&RolePermission.RolePermission{},
		//&Permission.Permission{},
		//&ActivityConfig.ActivityConfig{},
		//&ActivityLimitConfig.ActivityLimitConfig{},
		//&FTwitterMedia.FTwitterMedia{},
		//&FTwitterSpiderHis.FTwitterSpiderHis{},
		//&FTwitterTweet.FTwitterTweet{},
		//&FTwitterUser.FTwitterUser{},
		//&FTwitterUserHis.FTwitterUserHis{},
	); err != nil {
		logger.Error(err)
	} else {
		logger.Info("migration end")
	}
}
