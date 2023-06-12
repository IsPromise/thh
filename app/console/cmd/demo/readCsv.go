package demo

import (
	"github.com/spf13/cobra"
)

func init() {
	appendCommand(&cobra.Command{Use: "demo:readCsv", Short: "readCsv", Run: readCsv})
}
func readCsv(_ *cobra.Command, _ []string) {

}
