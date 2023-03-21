package test

import (
	"fmt"
	"testing"
)

type gStruct[T any] struct {
	Flag  string
	Value T
}

func TestGenerics(t *testing.T) {
	a := gStruct[int]{
		Flag:  "Flag name",
		Value: 1,
	}
	fmt.Println(a)
}
