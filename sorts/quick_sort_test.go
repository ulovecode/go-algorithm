package sorts

import "testing"

func TestQuickSort(t *testing.T) {
	arr := []int{123121, -21, 1231, 21, 3, -123, 12, 31, 3, 1, 521}
	QuickSort(arr)
	for _, v := range arr {
		println(v)
	}

	//arr = []int{123121}
	// QuickSort(arr)
	//for _, v := range arr {
	//	println(v)
	//}
	//
	//arr = []int{}
	// QuickSort(arr)
	//for _, v := range arr {
	//	println(v)
	//}
}
