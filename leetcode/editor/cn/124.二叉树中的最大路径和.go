package leetcode

import (
	"math"
)

//给定一个非空二叉树，返回其最大路径和。
//
// 本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。
//
// 示例 1:
//
// 输入: [1,2,3]
//
//       1
//      / \
//     2   3
//
//输出: 6
//
//
// 示例 2:
//
// 输入: [-10,9,20,null,null,15,7]
//
//   -10
//   / \
//  9  20
//    /  \
//   15   7
//
//输出: 42
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
var maxValue = math.MinInt32

func maxPathSum(root *TreeNode) int {
	getPathSum(root)
	return maxValue
}

func getPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftSum, rightSum := 0, 0
	leftSum = getMaxValue(getPathSum(root.Left), leftSum)
	rightSum = getMaxValue(getPathSum(root.Right), rightSum)
	maxV := leftSum + rightSum + root.Val
	if maxV > maxValue {
		maxValue = maxV
	}
	return root.Val + getMaxValue(leftSum, rightSum)
}

func getMaxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
