package leetcode

import "math"

//给定一个二叉树，判断它是否是高度平衡的二叉树。
//
// 本题中，一棵高度平衡二叉树定义为：
//
//
// 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。
//
//
// 示例 1:
//
// 给定二叉树 [3,9,20,null,null,15,7]
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
//
// 返回 true 。
//
//示例 2:
//
// 给定二叉树 [1,2,2,3,3,null,null,4,4]
//
//        1
//      / \
//     2   2
//    / \
//   3   3
//  / \
// 4   4
//
//
// 返回 false 。
//
//
// Related Topics 树 深度优先搜索

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	return isBalancedTree(root).isBalanced
}
func isBalancedTree(root *TreeNode) TreeDp {
	if root == nil {
		return TreeDp{
			isBalanced: true,
			high:       0,
		}
	}
	leftTreeDp := isBalancedTree(root.Left)
	rightTreeDp := isBalancedTree(root.Right)
	return TreeDp{
		isBalanced: leftTreeDp.isBalanced && rightTreeDp.isBalanced &&
			(int(math.Abs(float64(leftTreeDp.high-rightTreeDp.high))) <= 1),
		high: int(math.Max(float64(leftTreeDp.high), float64(rightTreeDp.high))) + 1,
	}
}

type TreeDp struct {
	isBalanced bool
	high       int
}

//leetcode submit region end(Prohibit modification and deletion)
