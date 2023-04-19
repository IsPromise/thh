package test

import (
	"fmt"
	"github.com/spf13/cast"
	"testing"
)

func TestSetV(t *testing.T) {
	type Cat struct {
		Name string
	}

	catMap := make(map[int]Cat, 4)

	for key, data := range catMap {
		fmt.Println(key, data)
	}

	for i := 0; i < 4; i++ {
		catMap[i] = Cat{Name: cast.ToString(i)}
	}
	fmt.Println("init end")
	for key, data := range catMap {
		fmt.Println(key, data)
	}

	for _, data := range catMap {
		data.Name = "1" + data.Name
	}
	fmt.Println("update1 end")

	for key, data := range catMap {
		fmt.Println(key, data)
	}

	for key, data := range catMap {
		data.Name = "1" + data.Name
		catMap[key] = data
	}
	fmt.Println("update2 end")

	for key, data := range catMap {
		fmt.Println(key, data)
	}

}

func TestSetV2(t *testing.T) {

	stringMap := make(map[int]*string, 4)
	s := "hello"

	// 将 key 为 1 的元素的 value 设置为指向字符串 s 的指针

	stringMap[1] = &s

	*stringMap[1] += "asdasda"
	fmt.Println(stringMap)
}

func TestSetV3(t *testing.T) {
	type Cat struct {
		Name string
	}

	catMap := make(map[int]*Cat, 4)

	for key, data := range catMap {
		fmt.Println(key, data)
	}

	for i := 0; i < 4; i++ {
		catMap[i] = &Cat{Name: cast.ToString(i)}
	}
	fmt.Println("init end")
	for key, data := range catMap {
		fmt.Println(key, data)
	}

	for _, data := range catMap {
		data.Name = "1" + data.Name
	}
	for key, data := range catMap {
		fmt.Println(key, data)
	}
}
