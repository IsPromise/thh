package test

import (
	"fmt"
	"math"
	"testing"
)

func TestValue(t *testing.T) {
	var b int
	b = math.MaxInt
	fmt.Println(b)
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MaxInt64)
}
