package demo

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func init() {
	appendCommand(&cobra.Command{
		Use:   "demo:stdin",
		Short: "标准输入",
		Run:   runStdin,
		Args:  cobra.ExactArgs(0), // 只允许且必须传 0 个参数
	})
}

func runStdin(_ *cobra.Command, _ []string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		result, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		fmt.Println(cast.ToString(result))
	}
}
