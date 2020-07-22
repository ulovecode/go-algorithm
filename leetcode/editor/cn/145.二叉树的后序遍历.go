package leetcode

import "container/list"

//给定一个二叉树，返回它的 后序 遍历。
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
//输出: [3,2,1]
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := list.New()
	stack.PushBack(root)
	stack1 := list.New()
	for stack.Len() != 0 {
		element := stack.Back()
		stack.Remove(element)
		node := element.Value.(*TreeNode)
		stack1.PushBack(node)
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
	}
	for stack1.Len() != 0 {
		element := stack1.Back()
		stack1.Remove(element)
		node := element.Value.(*TreeNode)
		res = append(res, node.Val)
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
