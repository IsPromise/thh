package arms

import (
	"fmt"
	"net/http"
	"time"
)

func Date() string {
	return time.Now().UTC().Format(http.TimeFormat)
}

func GetMicroTime() int64 {
	return time.Now().UnixMicro()
}

// MicrosecondsStr 将 time.Duration 类型（nano seconds 为单位）
// 输出为小数点后 3 位的 ms （microsecond 毫秒，千分之一秒）
func MicrosecondsStr(elapsed time.Duration) string {
	return fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)
}

func FirstElement(args []string) string {
	if len(args) > 0 {
		return args[0]
	}
	return ""
}
