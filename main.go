package main

import (
	"embed"
	_ "net/http/pprof"
	kernel2 "thh/app/bundles/kernel"
	"thh/app/console"
)

//go:embed  all:actor/dist/**
var actorFS embed.FS

//go:embed config.example.toml
var oConfig string

func main() {
	// 注册静态资源
	kernel2.InitStart()
	kernel2.ActorSave(actorFS)
	kernel2.SetOConfig(oConfig)
	console.Execute()
}
