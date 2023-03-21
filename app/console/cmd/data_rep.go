package cmd

import (
	"fmt"
	"thh/app/models/DataReps"
	"time"

	"github.com/spf13/cobra"
	"gorm.io/gorm/utils"
)

func init() {
	appendCommand(&cobra.Command{Use: "demo:sqliteDBKV", Short: "demo:sqliteDBKV", Run: sqliteDBKV})
}

func sqliteDBKV(_ *cobra.Command, _ []string) {
	for i := 1; i <= 100; i++ {
		_ = DataReps.Set(utils.ToString(i), utils.ToString(time.Now().String()))
	}

	for i := 1; i <= 100; i++ {
		val := DataReps.Get(utils.ToString(i))
		fmt.Print(val)
	}
	fmt.Println()
}
