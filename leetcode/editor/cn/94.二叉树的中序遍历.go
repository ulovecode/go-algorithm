package leetcode

import (
	"container/list"
)

//给定一个二叉树，返回它的中序 遍历。
//
// 示例:
//
// 输入: [1,null,2,3]
//   1
//    \
//     2
//    /
//   3
//
//输出: [1,3,2]
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树 哈希表

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := list.New()
	for stack.Len() != 0 || root != nil {
		if root != nil {
			stack.PushBack(root)
			root = root.Left
		} else {
			element := stack.Back()
			stack.Remove(element)
			node := element.Value.(*TreeNode)
			res = append(res, node.Val)
			root = node.Right
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
