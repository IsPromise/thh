package remoteservice

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type BetterClient struct {
	httpClient *resty.Client
	errHandle  func(err error)
}

func (itself *BetterClient) Init(h func(client *resty.Client)) *BetterClient {
	itself.httpClient = resty.New()
	h(itself.httpClient)
	return itself
}

func (itself *BetterClient) ErrHandle(h func(err error)) *BetterClient {
	itself.errHandle = h
	return itself
}

func (itself *BetterClient) PostBodyE(path string, data any) (*resty.Response, error) {
	return itself.httpClient.R().SetBody(data).Post(path)
}

func (itself *BetterClient) PostFormE(path string, data map[string]string) (*resty.Response, error) {
	return itself.httpClient.R().SetFormData(data).Post(path)
}

func (itself *BetterClient) GetQueryE(path string, data ...map[string]string) (*resty.Response, error) {
	requestData := map[string]string{}
	if len(data) == 1 {
		requestData = data[0]
	}
	return itself.httpClient.R().SetQueryParams(requestData).Get(path)
}

func (itself *BetterClient) PostBody(path string, data any) []byte {
	rep, _ := itself.PostBodyE(path, data)
	return rep.Body()
}
func (itself *BetterClient) PostForm(path string, data map[string]string) []byte {
	rep, _ := itself.PostFormE(path, data)
	return rep.Body()
}

func (itself *BetterClient) GetQuery(path string, data ...map[string]string) []byte {
	rep, _ := itself.GetQueryE(path, data...)
	return rep.Body()
}

type HiTokotoService struct {
	BetterClient
}

func NewBaiduClient() HiTokotoService {
	client := HiTokotoService{}
	client.
		Init(func(client *resty.Client) {
			client.SetBaseURL("https://v1.hitokoto.cn/")
		}).
		ErrHandle(func(err error) {
			fmt.Println(err)
		})
	return client
}

func (itself *HiTokotoService) GetOneTokotoV2() []byte {
	return itself.GetQuery("/")
}
