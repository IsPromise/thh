package ropt

import (
	"github.com/go-resty/resty/v2"
	"time"
)

var std = resty.New()

func init() {
	std.SetTimeout(time.Second * 10)
}

func SetProxy(proxyURL string) {
	std.SetProxy(proxyURL)
}

func RemoveProxy() {
	std.RemoveProxy()
}

func Get(uri string, data ...map[string]string) (string, *resty.Response, error) {
	queryData := map[string]string{}
	if len(data) == 1 {
		queryData = data[0]
	}
	return updateResponse(std.R().SetQueryParams(queryData).Get(uri))
}

func Post(uri string, data ...any) (string, *resty.Response, error) {
	var bodyData any
	if len(data) == 1 {
		bodyData = data[0]
	}
	return updateResponse(std.R().SetBody(bodyData).Post(uri))
}
func PostFormData(uri string, data ...map[string]string) (string, *resty.Response, error) {
	formData := map[string]string{}
	if len(data) == 1 {
		formData = data[0]
	}
	return updateResponse(std.R().SetFormData(formData).Post(uri))
}

func DownFile(url string, filename string) (*resty.Response, error) {
	return std.R().SetOutput(filename).Get(url)
}
