package app

import "embed"

var actorFS embed.FS
var envExample string

func ActorSave(dataWebFS embed.FS) {
	actorFS = dataWebFS
}

func GetActorFS() embed.FS {
	return actorFS
}

func EnvExample(data string) {
	envExample = data
}

func GetEnvExample() string {
	return envExample
}
