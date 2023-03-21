package remoteservice

import (
	"fmt"
	"testing"

	"github.com/spf13/cast"
)

func TestBetterClient(t *testing.T) {
	client := NewBaiduClient()
	data := client.GetOneTokotoV2()
	fmt.Println(cast.ToString(data))
}
