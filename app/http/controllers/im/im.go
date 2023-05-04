package im

import (
	"github.com/leancodebox/goose/preferences"
	"thh/app/http/controllers/component"
)

func ImInfo() component.Response {
	var port = preferences.GetString("")
	return component.SuccessResponse(map[string]any{
		"ws": port,
	})
}
