package im

import (
	"thh/app/http/controllers/component"

	"github.com/leancodebox/goose/preferences"
)

func ImInfo() component.Response {
	var port = preferences.GetString("")
	return component.SuccessResponse(map[string]any{
		"ws": port,
	})
}
