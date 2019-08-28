package cn

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: struct{ s string }{s: "abcabcbb"},
			want: 3,
		},
		{
			name: "2",
			args: struct{ s string }{s: "bbbbb"},
			want: 1,
		},
		{
			name: "3",
			args: struct{ s string }{s: "pwwkew"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
