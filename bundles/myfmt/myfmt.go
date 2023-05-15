package myfmt

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
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

	// 获取项目根目录路径
	rootDir := getRootDir()

	// 简化文件名，只保留项目内的路径部分
	file = strings.TrimPrefix(file, rootDir)
	file = strings.TrimPrefix(file, "/")

	// 打印消息
	fmt.Printf("%s:%d %s: ", file, line, packageName)
	fmt.Println(v...)
}

func getRootDir() string {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		file = "???"
		return ""
	}

	for i := 0; i < 3; i++ {
		file = filepath.Dir(file)
	}
	return strings.ReplaceAll(file, "\\", "/")
}
