package codemake

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CmdMakeRemoteService = &cobra.Command{
	Use:   "remoteService",
	Short: "Create request file, example gen request user",
	Run:   runMakeRemoteService,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakeRemoteService(_ *cobra.Command, args []string) {

	// 格式化模型名称，返回一个 Model 对象
	model := makeModelFromString(args[0])

	buildWithOutput(
		map[string]any{
			"clientName": model.ClientName,
		},
		fmt.Sprintf("app/service/ropt/%s_client.go", model.PackageName),
		"tmpl/client.tmpl",
	)
}
