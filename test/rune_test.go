package test

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"
	"unicode/utf8"
)

func TestRune(t *testing.T) {

	fmt.Println("rune")
	fmt.Println([]byte("你好"))
	fmt.Println(string([]byte{228, 189, 160, 229, 165, 189}))
	fmt.Println([]rune("你好"))
	fmt.Println([]int32("你好"))
	fmt.Println(string([]rune("你好")))
	fmt.Println(string([]rune{20320, 22909}))
	fmt.Println(utf8.RuneCountInString("你好n"))
	fmt.Println(len([]rune("你好n")))

}

func TestUnicode(t *testing.T) {

	str := "Hello, \\u4e2d\\u6587! 42"
	// 输出转换后的字符串
	fmt.Println(unicodeDecode(str))

}
func unicodeDecode(str string) string {
	// 正则表达式，用于匹配Unicode编码部分
	re := regexp.MustCompile(`\\u[0-9a-fA-F]{4}`)

	// 将Unicode编码转换为中文字符串
	return re.ReplaceAllStringFunc(str, func(m string) string {
		code, _ := strconv.ParseInt(m[2:], 16, 32)
		return string(rune(code))
	})
}
