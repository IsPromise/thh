package arms

import "github.com/spf13/cast"

func InArray[vT any](need vT, haystack []vT) bool {
	for _, v := range haystack {
		if cast.ToString(need) == cast.ToString(v) {
			return true
		}
	}
	return false
}

func ArrayMap[T, R any](f func(T) R, list []T) (result []R) {
	result = make([]R, len(list))
	for index, item := range list {
		result[index] = f(item)
	}
	return
}

func ArrayFilter[T any](f func(T) bool, list []T) (result []T) {
	for _, item := range list {
		if f(item) {
			result = append(result, item)
		}
	}
	return
}

func ArrayFill[T any](startIndex int, num uint, value T) map[int]T {
	result := make(map[int]T, num)
	var i uint
	for i = 0; i < num; i++ {
		result[startIndex] = value
		startIndex++
	}
	return result
}

func ArrayKeyExists(key any, arr map[any]any) bool {
	if len(arr) == 0 {
		return false
	}
	for k := range arr {
		if key == k {
			return true
		}
	}
	return false
}

func Map2Slice[K comparable, V any](m map[K]V) []V {
	s := make([]V, len(m))
	i := 0
	for _, v := range m {
		s[i] = v
		i += 1
	}
	return s
}

func Slice2Map[V any, K comparable](s []V, f func(V) K) map[K]V {
	m := make(map[K]V, len(s))
	for _, v := range s {
		m[f(v)] = v
	}
	return m
}

func Map2Map[mK comparable, mV any, newK comparable, newV any](oldMap map[mK]mV, f func(mK, mV) (newK, newV)) map[newK]newV {
	newMap := make(map[newK]newV, len(oldMap))
	for k, v := range oldMap {
		nk, nv := f(k, v)
		newMap[nk] = nv
	}
	return newMap
}

//type RangeType[K comparable, V any] interface {
//	map[K]V | []V
//}
//
//func ToSplice[K comparable, V, V2 any, CollectionType map[K]V | []V](f func(K, V) V2, collection CollectionType) []V2 {
//	r := make([]V2, len(collection))
//	i := 0
//	for k, v := range collection {
//		r[i] = f(k, v)
//		i += 1
//	}
//	return r
//}

type OrderedType interface {
	Integer | Float | ~string
}

type Integer interface {
	Signed | Unsigned
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Float interface {
	~float32 | ~float64
}
