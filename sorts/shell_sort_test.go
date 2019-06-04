package sorts

import "testing"

func TestShellSort(t *testing.T) {
	arr := []int{123121, -21, 21, 3, 12, 31, 3, 1, 521}
	ShellSort(arr)
	for _, v := range arr {
		println(v)
	}

	arr = []int{123121}
	ShellSort(arr)
	for _, v := range arr {
		println(v)
	}

	arr = []int{}
	ShellSort(arr)
	for _, v := range arr {
		println(v)
	}
}
