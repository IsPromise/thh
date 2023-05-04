package tspider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/leancodebox/goose/fileopt"
	"github.com/spf13/cobra"
	"gorm.io/gorm/utils"
	"net/url"
	"strings"
)

var commands = make([]*cobra.Command, 0)

func GetCommands() []*cobra.Command {
	return commands
}
func appendCommand(handle *cobra.Command) {
	commands = append(commands, handle)
}

func GetDocumentByR(r resty.Response, filename string) {
	document := `
# curl
#{curl}
# response:  
#{response}
`
	_, err := url.Parse(r.Request.URL)
	if err != nil {
		fmt.Println("url无法解析")
		return
	}
	jData, _ := json.MarshalIndent(json2map(r.String()), "", "  ")
	document = strings.Replace(document, "#{response}", mdCode(string(jData), "json"), -1)
	byR := GetCurlByR(r)
	document = strings.Replace(document, "#{curl}", mdCode(byR.String(), "shell"), -1)
	err = fileopt.FilePutContents(filename+"/data.md", []byte(document), true)
	err = fileopt.FilePutContents(filename+"/data.json", []byte(jData), true)
	err = fileopt.FilePutContents(filename+"/data.sh", []byte(byR.String()), true)
	ifErr(err)
}

func mdTag(content string) string {
	return fmt.Sprintf("`%v`", content)
}
func mdCode(content string, codeType string) string {
	return fmt.Sprintf("```%v\n%v\n```", codeType, content)
}

func json2map(jsonStr any) (mapResult any) {
	_ = json.Unmarshal([]byte(utils.ToString(jsonStr)), &mapResult)
	return mapResult
}

func GetCurlByR(r resty.Response) bytes.Buffer {
	b2 := bytes.Buffer{}
	b2.WriteString(fmt.Sprintf("curl '%v' -X '%v'", r.Request.URL, r.Request.Method))
	for header, headerValue := range r.Request.Header {
		b2.WriteString(fmt.Sprintf(" -H '%v:%v'", header, headerValue[len(headerValue)-1]))
	}
	if r.Request.Body != nil {
		body, _ := json.Marshal(r.Request.Body)
		b2.WriteString(fmt.Sprintf(" --data-raw '%v'", string(body)))
	}
	b2.WriteString(" --compressed --insecure ")
	return b2
}
