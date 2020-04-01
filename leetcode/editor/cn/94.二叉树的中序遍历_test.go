package cn

import (
	"fmt"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	//t6 := &TreeNode{
	//	Val:   6,
	//	Left:  nil,
	//	Right: nil,
	//}
	//t5:= &TreeNode{
	//	Val:   5,
	//	Left:  nil,
	//	Right: nil,
	//}
	//t4 := &TreeNode{
	//	Val:   4,
	//	Left:  t5,
	//	Right: t6,
	//}
	t3 := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	t2 := &TreeNode{
		Val:   2,
		Left:  t3,
		Right: nil,
	}
	t1 := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: t2,
	}
	ints := inorderTraversal(t1)
	fmt.Println(ints)
}
