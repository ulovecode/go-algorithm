//给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
//
// 说明: 叶子节点是指没有子节点的节点。
//
// 示例:
//给定如下二叉树，以及目标和 sum = 22，
//
//               5
//             / \
//            4   8
//           /   / \
//          11  13  4
//         /  \    / \
//        7    2  5   1
//
//
// 返回:
//
// [
//   [5,4,11,2],
//   [5,8,4,5]
//]
//
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

func pathSum(root *TreeNode, sum int) [][]int {
	result := make([][]int, 0)
	pathGenerate(root, sum, nil, &result)
	return result
}

func pathGenerate(root *TreeNode, sum int, path []int, result *[][]int) {
	if path == nil {
		path = make([]int, 0)
	}
	if root == nil {
		return
	}
	sum -= root.Val
	path = append(path, root.Val)
	if sum == 0 && root.Left == nil && root.Right == nil {
		*result = append(*result, path)
	}
	pathGenerate(root.Left, sum, deepCopyArray(path), result)
	pathGenerate(root.Right, sum, path, result)
}

func deepCopyArray(arr []int) []int {
	copys := make([]int, len(arr))
	copy(copys, arr)
	return copys
}
