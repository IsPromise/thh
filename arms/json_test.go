package arms

import (
	"fmt"
	"testing"
)

func TestJsonEncode(t *testing.T) {
	type tmp struct {
		Name string
	}
	fmt.Println(JsonEncode(tmp{Name: "name"}))
}

func TestJsonDecode(t *testing.T) {
	type tmp struct {
		Name string
	}
	fmt.Println(JsonDecode[tmp](`{"name":"name"}`))
	type Cat struct {
		Id int `json:"id"`
	}
	type DogKing[T any] struct {
		Body T `json:"body"`
	}
	catList := JsonDecode[[]Cat](`[{"id":1}]`)
	fmt.Println(catList)
	catMap := JsonDecode[map[string]Cat](`{"ok":{"id":1}}`)
	fmt.Println(catMap)
	dk := JsonDecode[DogKing[Cat]](`{"body":{"id":1231}}`)
	fmt.Println(dk)
	fmt.Println(dk.Body)
	type HighDogKing map[string]DogKing[Cat]
	dk2 := JsonDecode[HighDogKing](`{"key":{"body":{"id":1231}}}`)
	fmt.Println(dk2)
	fmt.Println(dk2["key"])
	fmt.Println(dk2["key"].Body)
}