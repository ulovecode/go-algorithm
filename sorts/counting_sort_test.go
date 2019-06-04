package sorts

import "testing"

func TestCountingSort(t *testing.T) {
	arr := []int{-123121, -21, -21, -3, -12, -31, -3, -1, -521}
	arr = CountingSort(arr)
	for _, v := range arr {
		println(v)
	}
	arr = []int{}
	arr = CountingSort(arr)
	for _, v := range arr {
		println(v)
	}
}
