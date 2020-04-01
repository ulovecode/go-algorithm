package cn

import "testing"

func Test_min_free(t *testing.T) {
	arr := []int{5, 1, 3, 2, 0, 6}
	println(min_free(arr, 0, len(arr)-1, len(arr)))
	println(min_free1(arr, len(arr)))
}
