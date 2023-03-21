package test

import (
	"fmt"
	"reflect"
	"testing"
)

type Cat struct {
	Name string `name:"catName" json:"name11"`
}

func TestTag(_ *testing.T) {
	t := reflect.TypeOf(&Cat{}).Elem()
	fmt.Println(t)
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Tag)
		fmt.Println(t.Field(i).Tag.Get("json"))
		fmt.Println("sadasd")
	}
	tt, _ := t.FieldByName("name")
	fmt.Println(tt.Tag)
	fmt.Println(tt.Tag.Get("name"))
	fmt.Println(tt.Tag.Lookup("name"))
}
