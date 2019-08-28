package cn

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_intervalsSort(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"1",
			struct{ intervals [][]int }{intervals: [][]int{{23, 41}, {1, 2}, {1, 3}, {2, 4}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			intervalsSort(tt.args.intervals)
			println(fmt.Sprint(tt.args.intervals))
		})
	}
}

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"1",
			struct{ intervals [][]int }{intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}},
			[][]int{{1, 6}, {8, 10}, {15, 18}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
