package restyopt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/go-resty/resty/v2"
)

var std = resty.New()

func SetProxy(proxyURL string) {
	std.SetProxy(proxyURL)
}

func RemoveProxy() {
	std.RemoveProxy()
}

func updateResponse(resp *resty.Response, err error) (string, *resty.Response, error) {
	if resp == nil {
		return "", nil, err
	}
	responseBody := string(resp.Body())
	return responseBody, resp, err
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

func GetCurlByR(r resty.Response) string {
	curlCmd := bytes.Buffer{}
	curlCmd.WriteString(fmt.Sprintf("curl '%v' -X '%v'", r.Request.URL, r.Request.Method))
	for header, headerValue := range r.Request.Header {
		curlCmd.WriteString(fmt.Sprintf(" -H '%v:%v'", header, headerValue[len(headerValue)-1]))
	}
	if r.Request.Body != nil {
		body, _ := json.Marshal(r.Request.Body)
		curlCmd.WriteString(fmt.Sprintf(" --data-raw '%v' ", string(body)))
	}
	if r.Request.FormData != nil {
		for key, value := range r.Request.FormData {
			curlCmd.WriteString(fmt.Sprintf(` --form '%v="%v" `, key, value))
		}
	}
	curlCmd.WriteString(" --compressed --insecure ")
	return curlCmd.String()
}

func GenerateCurlCommand(resp *resty.Response) string {
	var cmd strings.Builder

	cmd.WriteString(fmt.Sprintf("curl '%v' -X '%v'", resp.Request.URL, resp.Request.Method))

	// 处理请求头信息
	for k, v := range resp.Request.Header {
		if k == "Content-Type" {
			cmd.WriteString(fmt.Sprintf(" -H '%v:%v'", k, v[0]))
		} else {
			cmd.WriteString(fmt.Sprintf(" -H '%v:%v'", k, strings.Join(v, ",")))
		}
	}

	// 处理请求体信息
	if resp.Request.Body != nil {
		bodyBytes, _ := json.Marshal(resp.Request.Body)
		if len(bodyBytes) > 0 {
			cmd.WriteString(fmt.Sprintf(" --data-raw '%v'", string(bodyBytes)))
		}
	}

	// 处理表单数据信息
	if len(resp.Request.FormData) > 0 {
		formValues := url.Values{}
		for key, value := range resp.Request.FormData {
			formValues.Set(key, value[0])
		}
		cmd.WriteString(fmt.Sprintf(" '%v'", formValues.Encode()))
	}

	// 添加其他参数信息
	cmd.WriteString(" --compressed --insecure ")

	return cmd.String()
}
