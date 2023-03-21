package test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestAstDemo(t *testing.T) {
	src := `
package main
func main() {
	println("Hello, World!")
}
`

	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}

	// Print the AST.
	ast.Print(fset, f)
}
