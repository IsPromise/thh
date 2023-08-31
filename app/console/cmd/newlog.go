package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"log/slog"
	"os"
	"path/filepath"
	"thh/app/bundles/logging"
)

func init() {
	cmd := &cobra.Command{
		Use:   "newlog",
		Short: "",
		Run:   runNewlog,
		// Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
	}
	// cmd.Flags().String("param", "value", "input params")
	appendCommand(cmd)
}

func runNewlog(cmd *cobra.Command, args []string) {
	// param, _ := cmd.Flags().GetString("param")

	replace := func(groups []string, a slog.Attr) slog.Attr {
		// Remove time.
		if a.Key == slog.TimeKey && len(groups) == 0 {
			return slog.Attr{}
		}
		// Remove the directory from the source's filename.
		if a.Key == slog.SourceKey {
			source := a.Value.Any().(*slog.Source)
			source.File = filepath.Base(source.File)
		}
		return a
	}
	fmt.Println("success")
	slog.Info("hello", "count", 3)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, ReplaceAttr: replace}))
	logger.Info("hello", "count", 3)
	data := map[string]any{"data": []string{"21", "121", "21"}, "sa": 12}
	logger.InfoContext(context.Background(), "dasds", "ok", data)
	slog.Default().Info("sadasdas")

	logging.Info()

	//logPath := preferences.Get("log.path", "./storage/logs/thh222.log")
	//if err := fileopt.FilePutContents(logPath, []byte(""), true); err != nil {
	//}
	//file, _ := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//
	//slog.New(slog.NewTextHandler(file, nil))

}
