package codemake

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CmdMakeCMD = &cobra.Command{
	Use:   "cmd",
	Short: "Create a command, should be snake_case, exmaple: gen cmd buckup_database",
	Run:   runMakeCMD,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakeCMD(_ *cobra.Command, args []string) {

	// 格式化模型名称，返回一个 Model 对象
	model := makeModelFromString(args[0])

	buildWithOutput(
		map[string]any{
			"PackageName": model.PackageName,
			"StructName":  model.StructName,
		},
		fmt.Sprintf("app/console/cmd/%s.go", model.PackageName),
		"tmpl/cmd.tmpl",
	)
}
