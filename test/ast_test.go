package test

import (
	"fmt"
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

type adder struct {
	sum int
}

func (a *adder) Add(x int) *adder {
	a.sum += x
	return a
}

func add(x int) *adder {
	return &adder{sum: x}
}

func TestAdd(t *testing.T) {
	fmt.Println(add(1).Add(2).sum)                      // 输出 3
	fmt.Println(add(1).Add(2).Add(3).sum)               // 输出 6
	fmt.Println(add(2).Add(3).Add(4).Add(5).sum)        // 输出 14
	fmt.Println(add(1).Add(2).Add(3).Add(4).sum)        // 输出 10
	fmt.Println(add(1).Add(2).Add(3).Add(4).Add(5).sum) // 输出 15
}
