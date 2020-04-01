package main

import "testing"

func Test_getMinKByBF(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{name: "", args: struct {
		//	arr []int
		//	k   int
		//}{arr: []int{6, 9, 1, 3, 1, 2, 2, 5, 6, 1, 3, 5, 9, 7, 2, 5, 6, 1, 9 }, k:5}, want: 2},
		//{name: "", args: struct {
		//	arr []int
		//	k   int
		//}{arr: []int{6, 9, 1, 3, 1, 2, 2, 5, 6, 1, 3, 5, 9, 7, 2, 5, 6, 1, 9 }, k:6}, want: 2},
		//{name: "", args: struct {
		//	arr []int
		//	k   int
		//}{arr: []int{6, 9, 1, 3, 1, 2, 2, 5, 6, 1, 3, 5, 9, 7, 2, 5, 6, 1, 9 }, k:8}, want: 3},
		//{name: "", args: struct {
		//	arr []int
		//	k   int
		//}{arr: []int{6, 9, 1, 3, 1, 2, 2, 5, 6, 1, 3, 5, 9, 7, 2, 5, 6, 1, 9 }, k:10}, want: 5},
		{name: "", args: struct {
			arr []int
			k   int
		}{arr: []int{6, 9, 1, 3, 1, 2, 2, 5, 6, 1, 3, 5, 9, 7, 2, 5, 6, 1, 9}, k: 16}, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMinKByBFPRT(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("getMinKByBFPRT() = %v, want %v", got, tt.want)
			}
		})
	}
}
