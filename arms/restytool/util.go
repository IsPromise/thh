package restytool

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/go-resty/resty/v2"

	"sync"
)

// LonelyClient 孤独的客户端，不设置baseurl
type LonelyClient struct {
	httpClient *resty.Client
}

func (itself *LonelyClient) SetProxy(proxyURL string) *LonelyClient {
	itself.httpClient.SetProxy(proxyURL)
	return itself
}

var LonelyOnce sync.Once

var LonelyStd = LonelyClient{}

func StdLonelyClient() *LonelyClient {
	LonelyOnce.Do(func() {
		client := resty.New()
		LonelyStd.httpClient = client
	})
	return &LonelyStd
}

func (itself *LonelyClient) Get(uri string, data ...map[string]string) (resp *resty.Response, err error) {
	queryData := map[string]string{}
	if len(data) == 1 {
		queryData = data[0]
	}
	return itself.httpClient.R().SetQueryParams(queryData).Get(uri)
}

func (itself *LonelyClient) Post(uri string, data ...any) (resp *resty.Response, err error) {
	var bodyData any
	if len(data) == 1 {
		bodyData = data[0]
	}
	return itself.httpClient.R().SetBody(bodyData).Post(uri)
}
func (itself *LonelyClient) PostFormData(uri string, data ...map[string]string) (resp *resty.Response, err error) {
	formData := map[string]string{}
	if len(data) == 1 {
		formData = data[0]
	}
	return itself.httpClient.R().SetFormData(formData).Post(uri)
}

func (itself *LonelyClient) downHttpFile(url string, filename string) (*resty.Response, error) {
	return itself.httpClient.R().SetOutput(filename).Get(url)
}

func DownFile(url string, filename string) (*resty.Response, error) {
	return StdLonelyClient().downHttpFile(url, filename)
}

func GetCurlByR(r resty.Response) bytes.Buffer {
	b2 := bytes.Buffer{}
	b2.WriteString(fmt.Sprintf("curl '%v' -X '%v'", r.Request.URL, r.Request.Method))
	for header, headerValue := range r.Request.Header {
		b2.WriteString(fmt.Sprintf(" -H '%v:%v'", header, headerValue[len(headerValue)-1]))
	}
	if r.Request.Body != nil {
		body, _ := json.Marshal(r.Request.Body)
		b2.WriteString(fmt.Sprintf(" --data-raw '%v' ", string(body)))
	}
	if r.Request.FormData != nil {
		for key, value := range r.Request.FormData {
			b2.WriteString(fmt.Sprintf(` --form '%v="%v" `, key, value))
		}
	}
	b2.WriteString(" --compressed --insecure ")
	return b2
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
