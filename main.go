package main

import (
	"embed"
	_ "net/http/pprof"
	"thh/app/console"
	app2 "thh/bundles/app"
)

//go:embed  all:actor/dist/**
var actorFS embed.FS

//go:embed config.example.toml
var oConfig string

func main() {
	// 注册静态资源
	app2.InitStart()
	app2.ActorSave(actorFS)
	app2.SetOConfig(oConfig)
	console.Execute()
}
