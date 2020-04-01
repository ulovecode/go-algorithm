package sorts

import "testing"

func TestHeapSort(t *testing.T) {
	ints := []int{5, 4, 2, 3, 1}
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"",
			struct{ arr []int }{arr: ints},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HeapSort(tt.args.arr)
		})
	}
}
