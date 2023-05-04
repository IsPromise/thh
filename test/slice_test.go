package test

import (
	"fmt"
	"testing"

	"github.com/leancodebox/goose/jsonopt"
)

type SliceCat struct {
	Age int
}

func TestNewSlice(t *testing.T) {
	l := []SliceCat{{1}, {1}, {1}, {1}, {1}}
	fmt.Println(l)
	updateCS2(l)
	fmt.Println(l)
	updateCS(l)
	fmt.Println(l)
	updateCS3(&l)
	fmt.Println(l)
}

func updateCS(l []SliceCat) {
	for i, cat := range l {
		cat.Age += 1
		l[i] = cat
	}
}

func updateCS2(l []SliceCat) {
	for _, cat := range l {
		cat.Age += 1
	}
}

func updateCS3(l *[]SliceCat) {
	for _, cat := range *l {
		cat.Age += 1
	}
}

func TestNewSlice2(t *testing.T) {
	l := []*SliceCat{{1}, {1}, {1}, {1}, {1}}
	fmt.Println(jsonopt.Encode(l))
	updateCS22(l)
	fmt.Println(jsonopt.Encode(l))
	updateCS21(l)
	fmt.Println(jsonopt.Encode(l))
}

func updateCS21(l []*SliceCat) {
	for i, cat := range l {
		cat.Age += 1
		l[i] = cat
	}
}

func updateCS22(l []*SliceCat) {
	for _, cat := range l {
		cat.Age += 1
	}
}
