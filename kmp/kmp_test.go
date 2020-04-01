package kmp

import "testing"

func Test_kmp(t *testing.T) {
	type args struct {
		next string
		str  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kmp(tt.args.next, tt.args.str); got != tt.want {
				t.Errorf("kmp() = %v, want %v", got, tt.want)
			}
		})
	}
}
