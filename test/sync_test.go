package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
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

type useLockEntity struct {
	lock sync.Mutex
}

func (itself *useLockEntity) sleep() {
	itself.lock.Lock()
	defer itself.lock.Unlock()
	time.Sleep(time.Second * 3)
	fmt.Println(time.Now())
}

func TestUseLock(t *testing.T) {
	u := useLockEntity{}
	wg := sync.WaitGroup{}
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			u.sleep()
		}()
	}
	wg.Wait()
}

type nodeDeep2 struct {
	age int
}

type nodeDeep1 struct {
	age int
	nd2 nodeDeep2
}

type nodeDeep0 struct {
	age int
	nd1 nodeDeep1
}

func updateND0(data nodeDeep0, newAge int) {
	data.age = newAge
}
func updateND1(data nodeDeep1, newAge int) {
	data.age = newAge
}
func updateND2(data nodeDeep2, newAge int) {
	data.age = newAge
}

func updateND0p(data *nodeDeep0, newAge int) {
	data.age = newAge
}
func updateND1p(data *nodeDeep1, newAge int) {
	data.age = newAge
}
func updateND2p(data *nodeDeep2, newAge int) {
	data.age = newAge
}

func updateND0pAll(data *nodeDeep0, newAge int) {
	data.age = newAge
	data.nd1.age = newAge
	data.nd1.nd2.age = newAge
}

func TestUpP(t *testing.T) {
	nd := nodeDeep0{}
	updateND1p(&nd.nd1, 10)
	updateND2(nd.nd1.nd2, 10)
	fmt.Println(nd)
	nd2 := nodeDeep0{}
	updateND0pAll(&nd2, 100)
	fmt.Println(nd2)

}
