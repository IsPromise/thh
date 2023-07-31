package cmd

import (
	"fmt"
	"github.com/leancodebox/goose/jsonopt"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "json_compact",
		Short: "",
		Run:   runJsonCompact,
		// Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
	}
	// cmd.Flags().String("param", "value", "input params")
	appendCommand(cmd)
}

func runJsonCompact(cmd *cobra.Command, args []string) {
	oJson := `
{
    "interceptRule":{
      "tagList":[
        {"headerList":["path","auth"],"tagDetail":"md5值"},
        {"headerList":["path","au"],"tagDetail":"md5值"},
        {"headerList":["ip","ua"],"tagDetail":"md5值"}
      ],
      "headerKeyForbid":[
        {"key":"ua","value":"(android)","type":"reg"},
        {"key":"ua","value":"android xxx xx xxx xxx xx","type":"all"}
      ],
      "ipBlackList":{
          "headerKey":"",
          "ipList":["1.2.3.4","1.2.3.4"]
      },
      "timeout":5
    },
    "interceptAction":{
      "type":"",
      "body":""
    }
}`
	jsonData := jsonopt.Decode[any](oJson)
	fmt.Println(jsonData)
	fmt.Println(jsonopt.Encode(jsonData))
	// param, _ := cmd.Flags().GetString("param")
	fmt.Println("success")
}
