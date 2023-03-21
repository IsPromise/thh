package demo

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
)

func init() {
	appendCommand(&cobra.Command{Use: "demo:signal", Short: "demo:signal", Run: signalHandle})
}

func signalHandle(_ *cobra.Command, _ []string) {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	fmt.Println("开始等待")
	<-quit
	fmt.Println("收到信号，结束")
}
