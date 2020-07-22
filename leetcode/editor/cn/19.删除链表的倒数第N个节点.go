package leetcode

//给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
//
// 示例：
//
// 给定一个链表: 1->2->3->4->5, 和 n = 2.
//
//当删除了倒数第二个节点后，链表变为 1->2->3->5.
//
//
// 说明：
//
// 给定的 n 保证是有效的。
//
// 进阶：
//
// 你能尝试使用一趟扫描实现吗？
// Related Topics 链表 双指针

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	node := head
	for node != nil {
		node = node.Next
		n--
	}
	if n == 0 {
		return head.Next
	}
	node = head
	for n != -1 {
		node = node.Next
		n++
	}
	if node != nil && node.Next != nil {
		node.Next = node.Next.Next
	}
	return head
}

//leetcode submit region end(Prohibit modification and deletion)
