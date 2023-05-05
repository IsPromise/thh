package remoteservice

import (
	"github.com/leancodebox/goose/jsonopt"

	"github.com/go-resty/resty/v2"
)

type ClientResponse[T any] struct {
	Response      T
	RestyResponse *resty.Response
}

func buildResponseEntity[R any](response *resty.Response, err error) (ClientResponse[R], error) {
	var result ClientResponse[R]
	result.RestyResponse = response
	if err != nil {
		return result, err
	}
	result.Response = jsonopt.Decode[R](response.Body())
	return result, err
}

//func t[R any](client resty.Client,url string,data any) (result ClientResponse[R], err2 error) {
//	return nil,err2
//}
