package cn

import (
	"container/list"
)

//合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。
//
// 示例:
//
// 输入:
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//输出: 1->1->2->3->4->4->5->6
// Related Topics 堆 链表 分治算法

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	var root *TreeNode
	for _, l := range lists {
		for ; l != nil; l = l.Next {
			root = insertNode(root, l.Val)
		}
	}
	l := list.New()
	orderTree(root, l)

	lh := l.Front()
	var head = &ListNode{}
	node := head
	for ; lh != nil; lh, node = lh.Next(), node.Next {
		node.Next = &ListNode{Val: lh.Value.(int)}
	}

	return head.Next
}
func orderTree(node *TreeNode, n *list.List) {
	if node == nil {
		return
	}
	orderTree(node.Left, n)
	n.PushFront(node.Val)
	orderTree(node.Right, n)
}

func insertNode(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return &TreeNode{Val: val}
	}
	v := node.Val
	if v < val {
		if node.Left != nil {
			insertNode(node.Left, val)
		} else {
			node.Left = &TreeNode{Val: val}
		}
	} else {
		if node.Right != nil {
			insertNode(node.Right, val)
		} else {
			node.Right = &TreeNode{Val: val}
		}
	}
	return node
}

//leetcode submit region end(Prohibit modification and deletion)
