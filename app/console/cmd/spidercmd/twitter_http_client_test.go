package spidercmd

import (
	"fmt"
	"testing"
)

func TestTest(t *testing.T) {
	c := newTClient()
	r, err := c.getUserInfo("laravelphp")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r.String())
	}
}
