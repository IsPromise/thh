package test

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestRuntime(t *testing.T) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d B\n", m.Alloc/8)
	var i int64

	for i = 0; i < 100000000000; i++ {
		useMem()
		//start := time.Now()
		//runtime.GC()
		//elapsed := time.Now().Sub(start)
		//fmt.Println("该函数执行完成耗时：", elapsed)

		if i%100000 == 0 {
			runtime.ReadMemStats(&m)
			fmt.Printf("%d B\n", m.Alloc/8)
		}
	}
}
func useMem() {
	data := dataStruct{}
	data.name = "name"
}

type dataStruct struct {
	name string
}

func TestCaller(t *testing.T) {
	go func() {
		var data []any
		var buf [64]byte
		data = append(data, "d1", runtime.Stack(buf[:], false))
		fmt.Println(data)
		fmt.Println(fmt.Println(GoID()))
		go func() {
			var data []any
			var buf [64]byte
			data = append(data, "d2", runtime.Stack(buf[:], false))
			fmt.Println(fmt.Println(GoID()))
			fmt.Println(data)
		}()
	}()

	var data []any
	var buf [64]byte
	data = append(data, "d0", runtime.Stack(buf[:], false))

	time.Sleep(time.Second * 3)
	fmt.Println(runtime.Caller(0))
}
func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	// 得到id字符串
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
