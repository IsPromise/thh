package test

import (
	"fmt"
	"testing"
)

func TestForSwitch(t *testing.T) {
SELECT:
	for i := 1; i <= 10; i++ {
		switch {
		case i > 7:
			goto SELECT2
		case i > 5:
			break SELECT
		default:
			fmt.Println(i)
		}
	}
SELECT2:
	fmt.Println("select2")
}

func TestSwitch1609(t *testing.T) {
	a := 10
	switch a {
	case 10:
		fmt.Println(10)
		fallthrough
	case 9:
		fmt.Println(9)
	default:
		fmt.Println("无")
	}
}

func TestSwitch(t *testing.T) {
	a := 10
	a += 1
	switch {
	case a > 15:
		fmt.Println("大于1")
	case a > 5:
		fmt.Println("大于5")
		fallthrough
	case a > 0:
		fmt.Println("大于0")
	default:
		fmt.Println("无")
	}
}
