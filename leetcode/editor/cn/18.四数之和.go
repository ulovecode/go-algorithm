package leetcode

import (
	"math"
	"sort"
)

//给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
//
// 注意：
//
// 答案中不可以包含重复的四元组。
//
// 示例：
//
// 给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
//
//满足要求的四元组集合为：
//[
//  [-1,  0, 0, 1],
//  [-2, -1, 1, 2],
//  [-2,  0, 0, 2]
//]
//
// Related Topics 数组 哈希表 双指针

//leetcode submit region begin(Prohibit modification and deletion)
func fourSum(nums []int, target int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)

	var prePNum = math.MaxInt32
	var preINum = math.MaxInt32

	for p := 0; p < len(nums); p++ {
		if prePNum == nums[p] {
			continue
		}

		for i := p + 1; i < len(nums); i++ {
			if preINum == nums[i] {
				continue
			}
			for j, k := i+1, len(nums)-1; j < k; {
				if nums[p]+nums[i]+nums[j]+nums[k] == target {
					var preNums []int
					if len(result) != 0 {
						preNums = result[len(result)-1]
					}
					if preNums != nil && preNums[0] == nums[p] && preNums[1] == nums[i] && preNums[2] == nums[j] && preNums[3] == nums[k] {
					} else {
						result = append(result, []int{nums[p], nums[i], nums[j], nums[k]})
					}
					k--
					j++
				}
				if nums[p]+nums[i]+nums[j]+nums[k] > target {
					k--
				}
				if nums[p]+nums[i]+nums[j]+nums[k] < target {
					j++
				}
			}
			preINum = nums[i]
		}
		prePNum = nums[p]
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
