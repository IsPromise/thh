package one

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/spf13/cobra"
)

func init() {
	appendCommand(&cobra.Command{Use: "one", Short: "one", Run: one})
}

func one(_ *cobra.Command, _ []string) {
	src := `
package main
func main() {
    println("Hello, World!")
}
`
	// Create the AST by parsing src.
	fSet := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fSet, "", src, 0)
	if err != nil {
		panic(err)
	}

	// Print the AST.
	_ = ast.Print(fSet, f)
}
