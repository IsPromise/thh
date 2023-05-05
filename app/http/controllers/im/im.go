package im

import (
	"thh/app/http/controllers/component"

	"github.com/leancodebox/goose/preferences"
)

func ImInfo() component.Response {
	var port = preferences.GetString("app.outsitePort")
	return component.SuccessResponse(map[string]any{
		"ws": port,
	})
}
