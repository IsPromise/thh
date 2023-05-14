package myfmt

import (
	"fmt"
	"runtime"
)

func PrintlnWithCaller(v ...interface{}) {
	// 获取调用方的信息
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}

	// 获取调用方的包名
	fn := runtime.FuncForPC(pc)
	packageName := "???"
	if fn != nil {
		packageName = fn.Name()
		dotIndex := len(packageName) - 1
		for i := dotIndex; i >= 0; i-- {
			if packageName[i] == '.' {
				packageName = packageName[:i]
				break
			}
		}
	}

	// 打印消息
	fmt.Printf("[%s:%d] %s: ", file, line, packageName)
	fmt.Println(v...)
}
