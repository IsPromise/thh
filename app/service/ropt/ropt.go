package ropt

import (
	"github.com/leancodebox/goose/jsonopt"

	"github.com/go-resty/resty/v2"
)

func decodeResponse[R any](resp *resty.Response, err error) (R, *resty.Response, error) {
	var entity R
	if resp == nil || err != nil {
		return entity, nil, err
	}
	entity, err = jsonopt.DecodeE[R](resp.Body())
	return entity, resp, err
}

func updateResponse(resp *resty.Response, err error) (string, *resty.Response, error) {
	if resp == nil || err != nil {
		return "", nil, err
	}
	responseBody := string(resp.Body())
	return responseBody, resp, err
}
