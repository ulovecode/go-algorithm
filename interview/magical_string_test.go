package interview

import "testing"

func Test_magicalString(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			struct{ n int }{n: 3},
			1,
		},
		{
			"",
			struct{ n int }{n: 4},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := magicalString(tt.args.n); got != tt.want {
				t.Errorf("magicalString() = %v, want %v", got, tt.want)
			}
		})
	}
}
