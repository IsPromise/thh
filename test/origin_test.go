package test

import (
	"fmt"
	"testing"

	"github.com/spf13/cast"
)

func TestHanming(t *testing.T) {
	var data [4]int
	data[0] = 1
	data[1] = 0
	data[2] = 1
	data[3] = 0

	var check [3]int
	check[0] = 0
	check[1] = 0
	check[2] = 0

	var haimingCode [7]int
	haimingCode[0] = 0
	haimingCode[1] = 0
	haimingCode[2] = 0
	haimingCode[3] = 0
	haimingCode[4] = 0
	haimingCode[5] = 1
	haimingCode[6] = 1
	start := 1

	for {
		fmt.Println(start)
		t := 0
		for i := 1; i <= 7; i++ {
			fmt.Println(i & start)
			if i&start >= 1 {
				t = t ^ haimingCode[i-1]
			}
		}
		fmt.Println(t)
		haimingCode[start-1] = t
		start = start << 1
		if start > len(haimingCode) {
			break
		}
	}

	fmt.Println(haimingCode)

	w1 := func(a int, b int) int {
		return cast.ToInt(!(a == b))
	}
	w2 := func(a int, b int) int {
		return cast.ToInt(a != b)
	}
	fmt.Print(w1(1, 1))
	fmt.Print(w1(1, 0))
	fmt.Print(w1(0, 0))
	fmt.Println(w1(0, 1))

	fmt.Print(w2(1, 1))
	fmt.Print(w2(1, 0))
	fmt.Print(w2(0, 0))
	fmt.Println(w2(0, 1))

	fmt.Print(1 ^ 1)
	fmt.Print(1 ^ 0)
	fmt.Print(0 ^ 0)
	fmt.Println(0 ^ 1)
}

func TestDirectMapping(t *testing.T) {

	type cacheUnit struct {
		tag       int    // memory的高位地址
		byteBlock string // memory数据
	}

	type memoryAddress struct {
		memoryTag        int
		cacheAddress     int
		byteBlockAddress int
	}

	// 假设mem数量为1024 则需要10位表示所有位置

	// **********
	// tag|cacheAddress|byteBlockAddress
	// 10-5=5|5|x x和byteBlock的长度相关

	// 因为需要10位才可以确定一个mem地址，所以cacheList的映射表的单条数据需要10位
	// 假设cache数量为32 则需要5位表示所有位置
	// 也就是10位中的前6位可以确定cache位置

	memory := make(map[int]int, 1024)
	for i := 0; i < 1024; i++ {
		memory[i] = i
	}

	cache := make(map[int]cacheUnit, 32)
	// 将100～131的数据存下来
	for i := 0; i < 32; i++ {
		t := 100 + i
		tagIndex := t % 32
		tag := t / 32
		byteBlock := memory[t]
		cache[tagIndex] = cacheUnit{tag, "byte" + cast.ToString(byteBlock)}
	}

	var address memoryAddress

	address = memoryAddress{
		memoryTag:        3,
		cacheAddress:     18,
		byteBlockAddress: 1, // 无意义，目前并没有对byteBlock进行进一步细分
	}

	address2 := memoryAddress{
		memoryTag:        5,
		cacheAddress:     18,
		byteBlockAddress: 1, // 无意义，目前并没有对byteBlock进行进一步细分
	}

	fmt.Println(memory, cache, address)
	cUnit, ok := cache[address2.cacheAddress]
	if !ok {
		fmt.Println("缓存不存在")
	} else {
		if address2.memoryTag == cUnit.tag {
			fmt.Println("缓存存在且命中,内容为:", cUnit.byteBlock)
		} else {
			fmt.Println("缓存存在，但未命中，因为当前存储的缓存非对应内存的缓存")
		}
	}
	cUnit, ok = cache[address.cacheAddress]
	if !ok {
		fmt.Println("缓存不存在")
	} else {
		if address.memoryTag == cUnit.tag {
			fmt.Println("缓存存在且命中,内容为:", cUnit.byteBlock)
		} else {
			fmt.Println("缓存存在，但未命中，因为当前存储的缓存非对应内存的缓存")
		}
	}

}

func TestFullJoin(t *testing.T) {
	type cacheUnit struct {
		tag       int
		byteBlock string
	}

	type memoryAddress struct {
		memTag           int
		byteBlockAddress int
	}

	memory := make(map[int]int, 1024)
	for i := 0; i < 1024; i++ {
		memory[i] = i
	}

	// 假设mem数量为1024 则需要10位表示所有位置
	// **********
	// memTag|byteBlockAddress
	// 10|x x和byteBlock的长度相关

	// 因为需要10位才可以确定一个mem地址
	// 所以前10位均用于存储标记
	// 因为 memoryAddress 和 cacheList 不存在映射关系，
	// 所以需要用memoryAddress的数据去遍历cacheList,
	// 与每一个cacheUnit中的tag对比，
	// 如果相等则命中，如果全都不相等，则未命中

	// 因为全相联映射无法直接定位到cacheList的地址，所以要和所有的地址进行比较，所以这里采用数组模拟
	var cache [32]cacheUnit
	cacheListIndex := 0
	for i := 0; i < 32; i++ {
		t := 100 + i
		tag := t
		byteBlock := memory[t]
		cache[cacheListIndex] = cacheUnit{tag, "byte" + cast.ToString(byteBlock)}
		cacheListIndex += 1
	}
	mAddress := memoryAddress{103, 0}
	tagFind := false
	var cacheData string
	for _, value := range cache {

		if value.tag == mAddress.memTag {
			tagFind = true
		}

		if tagFind {
			cacheData = value.byteBlock
			break
		}
	}

	fmt.Println("使用", mAddress, "查找")
	if tagFind {
		fmt.Println("命中")
		fmt.Println(cacheData)
	} else {
		fmt.Println("未命中")
	}

}

func TestGroupJoin(t *testing.T) {
	type cacheUnit struct {
		valid     bool
		tag       int
		byteBlock string
	}

	type memoryAddress struct {
		memTag           int
		groupAddress     int
		byteBlockAddress int
	}

	memory := make(map[int]int, 1024)
	for i := 0; i < 1024; i++ {
		memory[i] = i
	}

	// 假设mem数量为1024 则需要10位表示所有位置
	// **********
	// memTag|byteBlockAddress
	// 10|x x和byteBlock的长度相关

	// 因为需要10位才可以确定一个mem地址
	// 因为cache分为了32块
	// 两块为1组
	// 所以有16组，16 为 2^4
	// 所以前10位中 后4位存为组地址,剩余的8位存为标记
	groupList := make(map[int][2]cacheUnit, 16)
	for i := 0; i < 16; i++ {
		groupList[i%16] = [2]cacheUnit{{false, 0, "byte"}, {false, 0, "byte"}}
	}

	// 将100～131的数据存下来
	for i := 0; i < 32; i++ {
		t := 100 + i
		groupAddress := t % 16
		tag := t / 16
		byteBlock := memory[t]
		groupUnitEntity, _ := groupList[groupAddress]
		for key, cacheUnitEntity := range groupUnitEntity {
			if !cacheUnitEntity.valid {
				groupUnitEntity[key] = cacheUnit{true, tag, "byte" + cast.ToString(byteBlock)}
				groupList[groupAddress] = groupUnitEntity
				break
			}
		}
	}

	// 搜索128

	mAddress := memoryAddress{124 / 16, 124 % 16, 1}

	cacheGroupEntity := groupList[mAddress.groupAddress]
	var data string
	for _, cacheUnitItem := range cacheGroupEntity {
		if cacheUnitItem.tag == mAddress.memTag && cacheUnitItem.valid {
			data = cacheUnitItem.byteBlock
			break
		}
	}

	fmt.Println(groupList)

	if len(data) > 0 {
		fmt.Println("成功命中", data)
	} else {
		fmt.Println("未命中", mAddress)
	}

}
