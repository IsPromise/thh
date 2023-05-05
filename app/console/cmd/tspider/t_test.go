package tspider

import (
	"fmt"
	"testing"
)

func TestStr2Header(t *testing.T) {
	oData := `Accept: */*
Accept-Encoding: gzip, deflate
Accept-Language: zh-CN,zh;q=0.9
Connection: keep-alive
Host: j1.58cdn.com.cn
Referer: http://oa.58v5.cn/
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36`

	fmt.Println(len(parseHeaders(oData)))

	for key, value := range parseHeaders(oData) {
		fmt.Println(fmt.Sprintf("key %v value %v", key, value))
	}
}
