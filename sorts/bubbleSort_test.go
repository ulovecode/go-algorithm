package sorts

import "testing"

func TestBubbleSort(t *testing.T) {
	arr := []int{123121, -21, 21, 3, 12, 31, 3, 1, 521}
	BubbleSort(arr)
	for _, v := range arr {
		println(v)
	}
}
