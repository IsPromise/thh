package arms

import (
	"fmt"
	"testing"
)

func TestIp(t *testing.T) {
	fmt.Println(GetLocalIp())
	fmt.Println(ExternalIP())
}
