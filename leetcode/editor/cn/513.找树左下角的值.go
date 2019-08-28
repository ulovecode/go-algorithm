//给定一个二叉树，在树的最后一行找到最左边的值。
//
// 示例 1:
//
//
//输入:
//
//    2
//   / \
//  1   3
//
//输出:
//1
//
//
//
//
// 示例 2:
//
//
//输入:
//
//        1
//       / \
//      2   3
//     /   / \
//    4   5   6
//       /
//      7
//
//输出:
//7
//
//
//
//
// 注意: 您可以假设树（即给定的根节点）不为 NULL。
//

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package cn

func findBottomLeftValue(root *TreeNode) int {
	_, result := leftBottom(root, 0)
	return result
}

func leftBottom(node *TreeNode, depth int) (int, int) {
	lDepth, rDepth, left, right := -1, -1, -1, -1
	if node.Left != nil {
		lDepth, left = leftBottom(node.Left, depth+1)
	}
	if node.Right != nil {
		rDepth, right = leftBottom(node.Right, depth+1)
	}
	if lDepth != -1 || rDepth != -1 {
		if lDepth >= rDepth {
			return lDepth, left
		} else {
			return rDepth, right
		}
	}
	return depth, node.Val
}
