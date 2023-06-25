package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	cmd := &cobra.Command{
		Use:   "scan_root",
		Short: "",
		Run:   runScanRoot,
		// Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
	}
	// cmd.Flags().String("param", "value", "input params")
	appendCommand(cmd)
}

func runScanRoot(cmd *cobra.Command, args []string) {
	// param, _ := cmd.Flags().GetString("param")
	fmt.Println("success")
	rootDir := "/" // 根目录
	showTree(rootDir, 0)
}

func showTree(dirPath string, level int) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Printf("无法读取目录：%s\n", dirPath)
		return
	}

	for _, file := range files {
		filename := file.Name()
		indent := strings.Repeat("  ", level) // 缩进

		if file.IsDir() {
			fmt.Printf("%s|-- %s/\n", indent, filename)
			subDirPath := filepath.Join(dirPath, filename)
			showTree(subDirPath, level+1) // 递归进入子目录
		} else {
			fmt.Printf("%s|-- %s\n", indent, filename)
		}
	}
}
