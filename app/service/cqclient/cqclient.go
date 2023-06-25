package cqclient

import (
	"github.com/go-resty/resty/v2"
	"github.com/leancodebox/goose/preferences"
	"github.com/spf13/cast"
	"time"
)

var (
	cqBaseUri = preferences.Get("cq.cqHttp")
	std       = resty.New().SetBaseURL(cqBaseUri)
)

func init() {
	std.SetTimeout(time.Second * 10)
}

type sendUnit struct {
	Type string         `json:"type"`
	Data map[string]any `json:"data"`
}

func Send4friend(userId, message any) (*resty.Response, error) {
	return std.R().SetQueryParams(map[string]string{
		"user_id":     cast.ToString(userId),
		"message":     cast.ToString(message),
		"auto_escape": cast.ToString(true),
	}).Get("send_private_msg")
}

func Send4group(groupId, message any) (*resty.Response, error) {
	return std.R().SetBody(map[string]any{
		"group_id":    groupId,
		"message":     message,
		"auto_escape": true,
	}).Post("send_group_msg")
}
