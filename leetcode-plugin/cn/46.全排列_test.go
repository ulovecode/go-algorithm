package cn

import (
	"reflect"
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
			if got := permute(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permute() = %v,\n want %v", got, tt.want)
			}
		})
	}
}
