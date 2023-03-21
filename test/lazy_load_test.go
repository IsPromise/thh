package test

import (
	"fmt"
	"github.com/spf13/cast"
	"sync"
	"testing"
)

type lazyCay struct {
	Name string
}

var l *lazyCay

func TestPTest(t *testing.T) {
	// 指针的赋值操作是原子性的
	l = &lazyCay{Name: "init"}
	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		defer wg.Done()
		for i := 1; i <= 100000; i++ {
			l = &lazyCay{Name: "1" + cast.ToString(i)}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 1; i <= 100000; i++ {
			l = &lazyCay{Name: "2" + cast.ToString(i)}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 1; i <= 100000; i++ {
			l = &lazyCay{Name: "3" + cast.ToString(i)}
		}
	}()
	go func() {
		for {
			fmt.Println(l.Name)
		}
	}()

	wg.Wait()
}
