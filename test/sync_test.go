package test

import (
	"fmt"
	"sync"
	"testing"
)

func TestSyncMap(t *testing.T) {
	m := sync.Map{}
	m.Store(1, 1)
	v, ok := m.Load(1)
	if !ok {
		fmt.Println("加载失败")
	}
	if value, ok := v.(int); ok {
		fmt.Println(value)
	}
}

func TestSyncPool(t *testing.T) {
	type cat struct {
		Name string
	}
	var catPool = sync.Pool{
		New: func() any {
			return &cat{}
		},
	}
	cat1 := catPool.Get()
	var catTmp *cat
	catTmp, isCat := cat1.(*cat)

	if !isCat {
		return
	}
	catTmp.Name = "小红"

	catPool.Put(catTmp)

	cat1 = catPool.Get()
	catTmp, isCat = cat1.(*cat)
	if !isCat {
		return
	}
	fmt.Println(catTmp.Name)
	// 使用后不放回pool

	cat1 = catPool.Get()
	catTmp, isCat = cat1.(*cat)
	if !isCat {
		return
	}
	fmt.Println(catTmp.Name)

	fmt.Println("syncPool")
}

func TestAutoPool(_ *testing.T) {
	sp.gp(sp.dogRead)

	fmt.Println("autoPool")
}

type dog struct {
}
type superPool struct {
	pool sync.Pool
}

var sp = superPool{
	pool: sync.Pool{
		New: func() any {
			return &dog{}
		},
	},
}

func (itself *superPool) gp(dogUp func(dogEntity *dog)) {
	dogI := itself.pool.Get()
	dogEntity, ok := dogI.(*dog)
	if !ok {
		return
	}
	dogUp(dogEntity)

	itself.pool.Put(dogEntity)
}

func (itself *superPool) dogRead(dogEntity *dog) {
	fmt.Println("dog")
}
