package demo

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
	"os/exec"
)

func init() {
	cmd := &cobra.Command{
		Use:   "input_client",
		Short: "",
		Run:   runInputClient,
		// Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
	}
	// cmd.Flags().String("param", "value", "input params")
	appendCommand(cmd)
}

func runInputClient(_ *cobra.Command, args []string) {

	type Request struct {
		Action string `json:"action"`
	}

	cmd := exec.Command("thh", "input_server") // 替换为你的交互式程序可执行文件名

	stdinPipe, stdinPipeErr := cmd.StdinPipe()
	if stdinPipeErr != nil {
		fmt.Println(stdinPipeErr)
		return
	}

	stdoutPipe, stdoutPipeErr := cmd.StdoutPipe()
	if stdoutPipeErr != nil {
		fmt.Println(stdoutPipeErr)
		return
	}

	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		defer stdinPipe.Close()

		encoder := json.NewEncoder(stdinPipe)
		decoder := json.NewDecoder(os.Stdin)
		for {
			var req Request
			err := decoder.Decode(&req)
			if err != nil {
				fmt.Println(err)
				break
			}

			err = encoder.Encode(req)
			if err != nil {
				fmt.Println(err)
				break
			}
		}
	}()

	decoder := json.NewDecoder(stdoutPipe)
	for {
		var response string
		err := decoder.Decode(&response)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}

		fmt.Println("Response:", response)
	}

	cmd.Wait()
}
