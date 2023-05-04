package remoteservice

import (
	"sync"

	"github.com/go-resty/resty/v2"
	"github.com/leancodebox/goose/jsonopt"
)

type HiTokotoClient struct {
	httpClient *resty.Client
}

var HiTokotoClientIns HiTokotoClient

var hiTokotoClientInsOnce sync.Once

func HiTokotoClientConnection(host ...string) HiTokotoClient {
	hiTokotoClientInsOnce.Do(func() {
		baseUri := "https://v1.hitokoto.cn/"
		if len(host) == 1 && len(host[0]) != 0 {
			baseUri = host[0]
		}
		HiTokotoClientIns = HiTokotoClient{
			httpClient: resty.New().SetBaseURL(baseUri),
		}
	})
	return HiTokotoClientIns
}

func (itself *HiTokotoClient) GetOneTokotoV2() (result ClientResponse[HiTokotoResponse], err error) {
	a, b := itself.httpClient.R().Get("/")
	return buildResponseEntity[HiTokotoResponse](a, b)
}

func (itself *HiTokotoClient) GetOneTokotoV3() (result HiTokotoResponse) {
	a, _ := itself.httpClient.R().Get("/")
	result = jsonopt.Decode[HiTokotoResponse]([]byte(a.String()))
	return
}

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
