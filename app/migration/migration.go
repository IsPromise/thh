package migration

import (
	"gorm.io/gorm"
	"thh/app/bundles/connect/dbconnect"
	"thh/app/bundles/kernel"
	"thh/app/bundles/logging"
	"thh/app/models/DataReps"
	"thh/app/models/FTwitter/FTwitterFilterUser"
	"thh/app/models/FTwitter/FTwitterMedia"
	"thh/app/models/FTwitter/FTwitterSpiderHis"
	"thh/app/models/FTwitter/FTwitterTweet"
	"thh/app/models/FTwitter/FTwitterUser"
	"thh/app/models/FTwitter/FTwitterUserHis"
	"thh/app/models/TodoTask"
	"thh/app/models/Users"
	"thh/app/models/rbac/ActivityConfig"
	"thh/app/models/rbac/ActivityLimitConfig"
	"thh/app/models/rbac/Permission"
	"thh/app/models/rbac/Role"
	"thh/app/models/rbac/RolePermission"
)

func M() {
	// 数据库迁移
	migration(kernel.UseMigration(), dbconnect.Std())
}

func migration(migration bool, db *gorm.DB) {
	if migration == false {
		return
	}
	logging.Info("init migration")
	// 自动迁移
	var err error

	if err = db.AutoMigrate(
		&DataReps.DataReps{},
		&Users.Users{},
		&Role.Role{},
		&RolePermission.RolePermission{},
		&Permission.Permission{},
		&ActivityConfig.ActivityConfig{},
		&ActivityLimitConfig.ActivityLimitConfig{},
		&FTwitterMedia.FTwitterMedia{},
		&FTwitterSpiderHis.FTwitterSpiderHis{},
		&FTwitterTweet.FTwitterTweet{},
		&FTwitterUser.FTwitterUser{},
		&FTwitterUserHis.FTwitterUserHis{},
		&FTwitterFilterUser.FTwitterFilterUser{},
		&TodoTask.Entity{},
	); err != nil {
		logging.Error(err.Error())
	} else {
		logging.Info("migration end")
	}
}
