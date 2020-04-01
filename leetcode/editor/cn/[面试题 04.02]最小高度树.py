# 给定一个有序整数数组，元素各不相同且按升序排列，编写一个算法，创建一棵高度最小的二叉搜索树。示例: 给定有序数组: [-10,-3,0,5,9], 一个可能
# 的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：           0          / \        -3 
#   9        /   /      -10  5 Related Topics 树 深度优先搜索


# leetcode submit region begin(Prohibit modification and deletion)
# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
from typing import List


class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def sortedArrayToBST(self, nums: List[int]) -> TreeNode:
        """
        :type nums: List[int]
        :rtype: TreeNode
        """
        if len(nums) == 0:
            return None
        t = TreeNode(nums[int(len(nums) / 2)])
        root = t
        l = int(len(nums) / 2)
        k = 0
        while k < l:
            if root.val > nums[k]:
                root.left = TreeNode(nums[k])
                root = root.left
            else:
                root.right = TreeNode(nums[k])
                root = root.right
            k = k + 1

        root = t
        r = int(len(nums) / 2) + 1
        k = int(len(nums)) - 1
        while k > r:
            if root.val > nums[k]:
                root.left = TreeNode(nums[k])
                root = root.left
            else:
                root.right = TreeNode(nums[k])
                root = root.right
            k = k - 1
        return t

    # leetcode submit region end(Prohibit modification and deletion)


if __name__ == '__main__':
    solution = Solution()
    tree_node = solution.sortedArrayToBST([0, 1])
