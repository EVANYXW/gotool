package sugar_plus

import (
	"fmt"
	"testing"
)

func TestSliceGroupBy(t *testing.T) {
	type st struct {
		key   string
		label string
	}

	or := []st{
		st{key: "group1", label: "lable1"},
		st{key: "group2", label: "lable2"},
		st{key: "group1", label: "lable3"},
		st{key: "group2", label: "lable4"},
	}

	s := SliceGroupBy(or, func(t st) string {
		return t.key
	})
	fmt.Println(s)
}

func TestShuffle(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	Shuffle(arr)
}
