package remoteservice

import (
	"github.com/leancodebox/goose/jsonopt"

	"github.com/go-resty/resty/v2"
)

type ClientResponse[T any] struct {
	Response      T
	RestyResponse *resty.Response
}

func buildResponseEntity[R any](response *resty.Response, err error) (result ClientResponse[R], err2 error) {
	result.RestyResponse = response
	err2 = err
	if err2 != nil {
		return
	}
	jsonopt.Decode[R](response.Body())
	return result, err2
}

//func t[R any](client resty.Client,url string,data any) (result ClientResponse[R], err2 error) {
//	return nil,err2
//}
