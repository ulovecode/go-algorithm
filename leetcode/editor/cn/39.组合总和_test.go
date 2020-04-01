package cn

import (
	"reflect"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		//{
		//	name: "1",
		//	args: args{
		//		candidates: []int{2, 3, 6, 7},
		//		target:     7,
		//	},
		//	want: [][]int{
		//		{7},
		//		{2, 2, 3},
		//	},
		//},
		//{
		//	name: "2",
		//	args: args{
		//		candidates: []int{4,2,8},
		//		target:     8,
		//	},
		//	want: [][]int{
		//		{2,2,2,2},
		//		{2,2,4},
		//		{4,4},
		//		{8},
		//	},
		//},

		{
			name: "2",
			args: args{
				candidates: []int{5, 10, 8, 4, 3, 12, 9},
				target:     27,
			},
			want: [][]int{
				{3, 7},
				{5, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
