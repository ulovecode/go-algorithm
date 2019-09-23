package cn

import (
	"math"
	"sort"
)

//给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
//
// 例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.
//
//与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
//
// Related Topics 数组 双指针

//leetcode submit region begin(Prohibit modification and deletion)
func threeSumClosest(nums []int, target int) int {
	result := math.MaxInt32
	sort.Ints(nums)
	for i := range nums {
		for l, r := i+1, len(nums)-1; l < r; {
			if int(math.Abs(float64((nums[i]+nums[l]+nums[r])-target))) < int(math.Abs(float64(result-target))) {
				result = nums[i] + nums[l] + nums[r]
			}
			if nums[i]+nums[l]+nums[r] == target {
				return target
			}
			if nums[i]+nums[l]+nums[r] > target {
				r--
			} else if nums[i]+nums[l]+nums[r] < target {
				l++
			}
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
