package cn

import "fmt"

//给定一个二叉树，返回所有从根节点到叶子节点的路径。
//
// 说明: 叶子节点是指没有子节点的节点。
//
// 示例:
//
// 输入:
//
//   1
// /   \
//2     3
// \
//  5
//
//输出: ["1->2->5", "1->3"]
//
//解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
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
func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	getBinaryTreePaths(root, "", &res)
	return res
}

func getBinaryTreePaths(root *TreeNode, prefix string, res *[]string) {
	if root == nil {
		return
	}
	value := fmt.Sprint(prefix, root.Val)
	if root.Left == nil && root.Right == nil {
		*res = append(*res, value)
		return
	}
	getBinaryTreePaths(root.Left, value+"->", res)
	getBinaryTreePaths(root.Right, value+"->", res)
}

//leetcode submit region end(Prohibit modification and deletion)
