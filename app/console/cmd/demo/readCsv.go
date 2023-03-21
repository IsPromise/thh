package demo

import (
	"fmt"
	"thh/arms"

	"github.com/spf13/cobra"
)

func init() {
	appendCommand(&cobra.Command{Use: "demo:readCsv", Short: "readCsv", Run: readCsv})
}
func readCsv(_ *cobra.Command, _ []string) {
	g := func() chan int {
		r := make(chan int)
		go func() {
			for i := 0; i < 10; i++ {
				r <- i
			}
			close(r)
		}()
		return r
	}
	for r := range g() {
		fmt.Println(r)
	}
	fmt.Println("readCsv")
	for r := range arms.ReadCsv("./tmp/t.csv") {
		fmt.Println(r)
	}
}
