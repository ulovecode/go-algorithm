package leetcode

import (
	"reflect"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	//root4 := &ListNode{Val: 5}
	//root3 := &ListNode{Val: 4, Next: root4}
	//root2 := &ListNode{Val: 3, Next: root3}
	root1 := &ListNode{Val: 2}
	root := &ListNode{Val: 1, Next: root1}
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "", args: struct {
			head *ListNode
			k    int
		}{head: root, k: 1}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroup(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
