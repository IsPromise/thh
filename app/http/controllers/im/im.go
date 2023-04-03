package im

import (
	"thh/app/http/controllers/component"
	"thh/bundles/config"
)

func ImInfo() component.Response {
	var port = config.GetString("APP_OUTSIDE_PORT")
	return component.SuccessResponse(map[string]any{
		"ws": port,
	})
}
