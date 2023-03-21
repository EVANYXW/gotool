package sugar_plus

import (
	"golang.org/x/exp/constraints"
	"math/rand"
	"time"
)

// CreateSlice 多维切片
func CreateSlice(row, col int) [][]int {
	flag := make([][]int, row)
	for k, _ := range flag {
		flag[k] = make([]int, col)
	}
	return flag

}

// Find slice 找值
func Find[T element](slice []T, val T) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func SliceFiltrate[V any](collection []V, filtrate func(V, int) bool) []V {

	result := []V{}

	for i, v := range collection {
		if filtrate(v, i) {
			result = append(result, v)
		}
	}

	return result
}

func SliceUpdateElement[T any, R any](collection []T, iteratee func(T, int) R) []R {
	result := make([]R, len(collection))

	for i, t := range collection {
		result[i] = iteratee(t, i)
	}

	return result
}

//func SliceUniq[T any, U comparable](collection []T, iteratee func(T) U) []T {
//	result := make([]T, len(collection))
//
//	seen := make(map[U]struct{}, len(collection))
//	for _, item := range collection {
//		key := iteratee(item)
//		if _, ok := seen[key]; ok {
//			continue
//		}
//		seen[key] = struct{}{}
//	}
//
//	return result
//}

func SliceGroupBy[T any, U comparable](collection []T, iteratee func(T) U) map[U][]T {
	result := map[U][]T{}

	for _, item := range collection {
		key := iteratee(item)

		result[key] = append(result[key], item)
	}
	return result
}

// CheckInSlice  check value in slice
func CheckInSlice[T constraints.Ordered](a T, s []T) bool {
	for _, val := range s {
		if a == val {
			return true
		}
	}
	return false
}

// DelOneInSlice  delete one element of slice  left->right
func DelOneInSlice[T constraints.Ordered](a T, old []T) (new []T) {
	for i, val := range old {
		if a == val {
			new = append(old[:i], old[i+1:]...)
			return
		}
	}
	return old
}

// Shuffle 打乱数组
func Shuffle[T any](slice []T) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1] //这里很巧妙的利用了外部slice的长度不可改变的特性来实现
	}
}
