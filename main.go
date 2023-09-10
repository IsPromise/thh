package main

import (
	"embed"
	_ "net/http/pprof"
	"thh/app/bundles/kernel"
	"thh/app/console"
)

//go:embed  all:actor/dist/**
var actorFS embed.FS

//go:embed config.example.toml
var oConfig string

//go:generate kuai tool:build_catalogue --output=catalogue.md
func main() {
	// 注册静态资源
	kernel.InitStart()
	kernel.ActorSave(actorFS)
	kernel.SetOConfig(oConfig)
	console.Execute()
}
