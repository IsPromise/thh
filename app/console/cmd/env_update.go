package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"thh/bundles/config"
	"time"
)

func init() {
	appendCommand(&cobra.Command{
		Use:   "env_update",
		Short: "",
		Run:   runEnvUpdate,
		// Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
	})
}

func runEnvUpdate(_ *cobra.Command, _ []string) {

	for {
		fmt.Println(config.Env("env_test", "init"))
		time.Sleep(time.Second * 1)
	}
}
