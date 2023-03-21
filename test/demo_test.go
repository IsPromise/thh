package test

import (
	"bytes"
	"fmt"
	"math/big"
	"regexp"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestPracticeBigNumber(_ *testing.T) {
	xf, _, _ := big.ParseFloat("1000000000.0000001", 10, 512, big.ToZero)
	yf, _, _ := big.ParseFloat("1000000000.0000001", 10, 512, big.ToZero)
	fmt.Println("xf:", xf.Text('f', 10))
	fmt.Println("yf:", yf.Text('f', 10))
	fmt.Println("xf+yf", xf.Add(xf, yf).Text('f', 10))
	fmt.Println("xf+xf", xf.Add(xf, xf).Text('f', 10))
	fmt.Println("xf*yf", xf.Mul(xf, yf).Text('f', 10))
	fmt.Println("xf/yf", xf.Quo(xf, yf).Text('f', 10))
	fmt.Println("xf-yf", xf.Sub(xf, yf).Text('f', 10))
	fmt.Println("cmp(xf,yf)", xf.Cmp(yf))
}

const (
	A int = iota
	B
	C
	D
	E
	F
	G
)

func EnumTString(enumValue int) {
	switch enumValue {
	case A:
		break
	case B:
		break
	case C:
		break
	case D:
		break
	case E:
		break
	case F:
		break
	case G:
		break
	}
}

func TestSyncOnce(_ *testing.T) {
	type runOne struct {
		once sync.Once
	}
	run := new(runOne)
	run.once.Do(func() {
		fmt.Println("第一次")
	})
	run.once.Do(func() {
		fmt.Println("第一次")
	})
}

func practiceReg(_ *testing.T) {
	regUnit := func(regStr string, matchStr string, unMatchStr string) {
		defer func() {
			fmt.Println("Enter defer function.")
			if p := recover(); p != nil {
				fmt.Printf("panic: %s\n", p)
			}
			fmt.Println("Exit defer function.")
		}()
		reg := regexp.MustCompile(regStr)
		//根据规则提取关键信息
		result1 := reg.FindAllStringSubmatch(matchStr, -1)
		fmt.Println(regStr, matchStr, "match", len(result1))
		//根据规则提取关键信息
		result2 := reg.FindAllStringSubmatch(unMatchStr, -1)
		fmt.Println(regStr, unMatchStr, "match", len(result2))

	}
	buf := "abc azc a7c aac 888 a9c  tac"
	//解析正则表达式，如果成功返回解释器
	reg := regexp.MustCompile(`a[0-9]c`)
	if reg == nil { //解释失败，返回nil
		fmt.Println("regexp err")
		return
	}
	//根据规则提取关键信息
	result1 := reg.FindAllStringSubmatch(buf, -1)
	fmt.Println("result1 = ", result1)

	fmt.Println("reg start")

	// 结尾匹配
	regUnit(`^[abcdef]*$`, `accede`, `beam`)

	// 结尾匹配
	regUnit(`[a-z]k$`, `Mick`, `Nickneven`)

	// 单词结尾
	regUnit(`fu\b`, `tofu`, `futz`)

	// ()重复出现 就是 匹配 （$字符串1）匹配 $字符串1
	// allochirally 可以匹配到 (all) ochir all 前面的all在后面再次出现了
	// go 目前不支持\1
	//regUnit(`(...).*\1`,`allochirally`,`anticker`)

	// go 目前不支持?!
	regUnit(`^(?!.*(.)(.)\2\1.*)`, `acritan`, `anallagmatic`)
	//
	//regUnit(``,``,``)
	//
	//regUnit(``,``,``)

}

func TestSuggestStrOperation(_ *testing.T) {
	fmt.Println("this is a suggestStrOperation")
	k := 5
	d := [5]time.Duration{}
	for i := 0; i < k; i++ {
		d[i] = benchmarkStringFunction(30000, i)
	}

	for i := 0; i < k-1; i++ {
		fmt.Printf("way %d is %6.1f times of way %d\n", i, float32(d[i])/float32(d[k-1]), k-1)
	}
}

func benchmarkStringFunction(n int, index int) (d time.Duration) {
	v := "ni shuo wo shi bu shi tai wu liao le a?"
	var s string
	var buf bytes.Buffer

	t0 := time.Now()
	for i := 0; i < n; i++ {
		switch index {
		case 0: // fmt.Sprintf
			s = fmt.Sprintf("%s[%s]", s, v)
		case 1: // string +
			s = s + "[" + v + "]"
		case 2: // strings.Join
			s = strings.Join([]string{s, "[", v, "]"}, "")
		case 3: // temporary bytes.Buffer
			// 每次声明一个bytes.buffer 仍然比正常的字符串拼接快很多
			b := bytes.Buffer{}
			b.WriteString("[")
			b.WriteString(v)
			b.WriteString("]")
			s = b.String()
		case 4: // stable bytes.Buffer
			buf.WriteString("[")
			buf.WriteString(v)
			buf.WriteString("]")
		}

	}

	if index == 4 { // for stable bytes.Buffer
		s = buf.String()
	}
	fmt.Println(len(s)) // consume s to avoid compiler optimization
	t1 := time.Now()
	d = t1.Sub(t0)
	fmt.Printf("time of way(%d)=%v\n", index, d)
	return d
}
