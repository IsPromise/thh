package test

import (
	"fmt"
	"testing"
)

func TestBitCompute(t *testing.T) {
	// 10进制数2进制打印
	var ii int64 = -5
	fmt.Printf("%b\n", ii)

	var i int64 = 5
	fmt.Printf("%b\n", i)

	// 8进制
	var a int64 = 011
	fmt.Println("a=", a)

	// 16进制
	var j = 0x11
	fmt.Println("j=", j)
}
