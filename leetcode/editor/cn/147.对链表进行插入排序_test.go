package leetcode

import (
	"reflect"
	"testing"
)

func Test_insertionSortList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	node3 := &ListNode{
		Val:  3,
		Next: nil,
	}
	node1 := &ListNode{
		Val:  1,
		Next: node3,
	}
	node2 := &ListNode{
		Val:  2,
		Next: node1,
	}
	node4 := &ListNode{
		Val:  4,
		Next: node2,
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "1",
			args: args{
				head: node4,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertionSortList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertionSortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
