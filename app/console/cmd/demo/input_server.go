package demo

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	cmd := &cobra.Command{
		Use:   "input_server",
		Short: "",
		Run:   runInputServer,
		// Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
	}
	// cmd.Flags().String("param", "value", "input params")
	appendCommand(cmd)
}

func runInputServer(_ *cobra.Command, _ []string) {
	// {"action":"get"}
	type Request struct {
		Action string `json:"action"`
	}
	decoder := json.NewDecoder(os.Stdin)
	encoder := json.NewEncoder(os.Stdout)

	for {
		var req Request
		err := decoder.Decode(&req)
		if err != nil {
			fmt.Println(err)
			return
		}

		switch req.Action {
		case "get":
			// 这里可以进行相关操作，根据需要获取所需数据
			response := "这是要获取的内容"
			err = encoder.Encode(response)
			if err != nil {
				fmt.Println(err)
				return
			}
		case "quit":
			return
		default:
			fmt.Println("未知操作")
		}
	}
}
