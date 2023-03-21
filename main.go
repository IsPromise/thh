package main

import (
	"embed"
	_ "net/http/pprof"
	"thh/app/console"
	"thh/arms/app"
)

//go:embed  all:actor/dist/**
var actorFS embed.FS

//go:embed .env.example
var envExample string

func main() {
	// 注册静态资源
	app.InitStart()
	app.ActorSave(actorFS)
	app.EnvExample(envExample)
	console.Execute()
}
