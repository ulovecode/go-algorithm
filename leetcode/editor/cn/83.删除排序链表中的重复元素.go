package leetcode

//给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
//
// 示例 1:
//
// 输入: 1->1->2
//输出: 1->2
//
//
// 示例 2:
//
// 输入: 1->1->2->3->3
//输出: 1->2->3
// Related Topics 链表

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	m := map[int]struct{}{}
	node := head
	preNode := head
	for ; node != nil; node = node.Next {
		for _, ok := m[node.Val]; ok; _, ok = m[node.Val] {
			if node.Next != nil && node.Next.Val == node.Val {
				node = node.Next
			} else {
				preNode.Next = node.Next
				break
			}
		}
		preNode = node
		m[node.Val] = struct{}{}
	}
	return head
}

//leetcode submit region end(Prohibit modification and deletion)
