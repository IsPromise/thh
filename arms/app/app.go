// Package app 应用信息
package app

import (
	"thh/bundles/config"
)

const on = "on"
const off = "off"

func UseMigration() bool {
	openMigration := config.GetString("database.mysql.openMigration")
	return openMigration == on
}

func IsProduction() bool {
	return config.Get("APP_ENV", "local") == "production"
}

// URL 传参 path 拼接站点的 URL
func URL(path string) string {
	return config.Get("APP_URL") + path
}
