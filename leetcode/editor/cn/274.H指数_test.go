package cn

import "testing"

func Test_hIndex(t *testing.T) {
	i := hIndex([]int{1, 1})
	println(i)
	i = hIndex([]int{3, 0, 6, 1, 5})
	println(i)
	i = hIndex([]int{100})
	println(i)
	i = hIndex([]int{0})
	println(i)
}
