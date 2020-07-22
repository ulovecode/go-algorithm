package leetcode

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func Test_permute(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"1",
			struct{ nums []int }{nums: []int{1, 2, 3}},
			[][]int{{1, 2, 3}, {2, 1, 3}, {2, 3, 1}, {1, 3, 2}, {3, 1, 2}, {3, 2, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := permute(tt.args.nums)
			strGot := make([]string, 0)
			for _, v := range got {
				strGot = append(strGot, fmt.Sprint(v))
			}
			strWant := make([]string, 0)
			for _, v := range tt.want {
				strWant = append(strWant, fmt.Sprint(v))
			}
			sort.Strings(strGot)
			sort.Strings(strWant)
			if !reflect.DeepEqual(strGot, strWant) {
				t.Errorf("permute() = %v,\n want %v", strGot, strWant)
			}
		})
	}
}
