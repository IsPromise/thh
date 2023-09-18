package test

import (
	"fmt"
	"github.com/leancodebox/goose/jsonopt"
	"testing"
)

type ResultEnum int

const (
	SUCCESS ResultEnum = iota
	ERROR
)

type EnumJson struct {
	Result ResultEnum `json:"result"`
}

var resultEnumNames = [...]string{
	"发起回款",
	"已核销回款",
	"保证金转合同",
	"申请开票",
	"开票确认",
	"发票回执",
	"发票归档",
	"撤销",
}

func (d ResultEnum) String() string {
	if d < SUCCESS || d > ERROR {
		return "Unknown"
	}
	return resultEnumNames[d]
}

func TestEnum(t *testing.T) {

	data := jsonopt.Decode[EnumJson](`{"result":3}`)
	fmt.Println(data)
	var eData ResultEnum
	eData = 4
	fmt.Println(eData)

}
