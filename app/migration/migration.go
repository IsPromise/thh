package migration

import (
	"fmt"
	"thh/app/models/DataReps"
	"thh/bundles/app"
	"thh/bundles/dbconnect"
	"thh/bundles/logger"

	"gorm.io/gorm"
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
		&DataReps.DataReps{},
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
