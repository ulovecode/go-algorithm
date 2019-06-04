package sorts

import "testing"

func TestMerge(t *testing.T) {
	arr := []int{123121, -21, 21, 3, 12, 31, 3, 1, 521}
	arr = MergeSort(arr)
	for _, v := range arr {
		println(v)
	}

	arr = []int{123121}
	arr = MergeSort(arr)
	for _, v := range arr {
		println(v)
	}

	arr = []int{}
	arr = MergeSort(arr)
	for _, v := range arr {
		println(v)
	}
}
