package main

import (
	"embed"
	_ "net/http/pprof"
	"thh/app/console"
	"thh/bundles/kernel"
)

//go:embed  all:actor/dist/**
var actorFS embed.FS

//go:embed config.example.toml
var oConfig string

func main() {
	// 注册静态资源
	kernel.InitStart()
	kernel.ActorSave(actorFS)
	kernel.SetOConfig(oConfig)
	console.Execute()
}
