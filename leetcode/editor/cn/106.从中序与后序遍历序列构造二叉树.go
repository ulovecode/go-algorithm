package leetcode

//根据一棵树的中序遍历与后序遍历构造二叉树。
//
// 注意:
//你可以假设树中没有重复的元素。
//
// 例如，给出
//
// 中序遍历 inorder = [9,3,15,20,7]
//后序遍历 postorder = [9,15,7,20,3]
//
// 返回如下的二叉树：
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
//
// Related Topics 树 深度优先搜索 数组

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postBuildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	var i, mid = len(postorder) - 1, -1
	var node *TreeNode
TAG:
	for ; i >= 0; i-- {
		val := postorder[i]
		node = &TreeNode{
			Val: val,
		}
		for i, v := range inorder {
			if v == val {
				mid = i
				break TAG
			}
		}
	}
	if mid <= len(inorder) {
		node.Left = buildTree(inorder[0:mid], postorder[:i])
	}
	if mid+1 >= 0 {
		node.Right = buildTree(inorder[mid+1:], postorder[:i])
	}
	return node
}

//leetcode submit region end(Prohibit modification and deletion)
