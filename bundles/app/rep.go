package app

import "embed"

var actorFS embed.FS
var oConfig string

func ActorSave(dataWebFS embed.FS) {
	actorFS = dataWebFS
}

func GetActorFS() embed.FS {
	return actorFS
}

func SetOConfig(data string) {
	oConfig = data
}

func GetOConfig() string {
	return oConfig
}
