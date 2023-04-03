package test

import (
	"fmt"
	"testing"

	"github.com/spf13/cast"
	"gopkg.in/yaml.v3"
)

type buildYml struct {
	Name  string
	Age   int
	Other struct {
		Ok   string
		Name string
	}
}

func TestName(t *testing.T) {
	b := buildYml{}
	//v.WriteConfig()
	aa, _ := yaml.Marshal(b)
	fmt.Println(cast.ToString(aa))
}
