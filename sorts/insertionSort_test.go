package sorts

import "testing"

func TestInsertionSort(t *testing.T) {
	arr := []int{123121, -21, 21, 3, 12, 31, 3, 1, 521}
	InsertionSort(arr)
	for _, v := range arr {
		println(v)
	}
	arr = []int{}
	InsertionSort(arr)
	for _, v := range arr {
		println(v)
	}

}
