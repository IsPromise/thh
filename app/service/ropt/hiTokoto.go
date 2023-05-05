package ropt

import (
	"github.com/go-resty/resty/v2"
)

var hiTokotoStd = resty.New().SetBaseURL("https://v1.hitokoto.cn/")

type HiTokotoResponse struct {
	ID         int    `json:"id"`
	UUID       string `json:"uuid"`
	Hitokoto   string `json:"hitokoto"`
	Type       string `json:"type"`
	From       string `json:"from"`
	FromWho    string `json:"from_who"`
	Creator    string `json:"creator"`
	CreatorUID int    `json:"creator_uid"`
	Reviewer   int    `json:"reviewer"`
	CommitFrom string `json:"commit_from"`
	CreatedAt  string `json:"created_at"`
	Length     int    `json:"length"`
}

func GetOneTokoto() (HiTokotoResponse, *resty.Response, error) {
	resp, err := hiTokotoStd.R().Get("/")
	return decodeResponse[HiTokotoResponse](resp, err)
}
