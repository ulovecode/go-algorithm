package leetcode

import "testing"

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		//{"1",
		//	struct {
		//		s       string
		//		numRows int
		//	}{s: "LEETCODEISHIRING", numRows: 4},
		//	"LDREOEIIECIHNTSG",
		//},
		{"2",
			struct {
				s       string
				numRows int
			}{s: "ABCD", numRows: 2},
			"ACBD",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
