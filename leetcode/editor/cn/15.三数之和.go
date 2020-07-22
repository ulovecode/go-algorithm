package leetcode

import (
	"math"
	"sort"
)

//给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
// 例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//满足要求的三元组集合为：
//[
//  [-1, 0, 1],
//  [-1, -1, 2]
//]
//
// Related Topics 数组 双指针

//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	var preINum = math.MaxInt32
	for i := range nums {
		if preINum == nums[i] {
			continue
		}
		if nums[i] > 0 {
			return result
		}
		for j, k := i+1, len(nums)-1; j < k; {
			if nums[i]+nums[j]+nums[k] == 0 {
				var preNums []int
				if len(result) != 0 {
					preNums = result[len(result)-1]
				}
				if preNums != nil && preNums[0] == nums[i] && preNums[1] == nums[j] && preNums[2] == nums[k] {
				} else {
					result = append(result, []int{nums[i], nums[j], nums[k]})
				}
				k--
				j++
			}
			if nums[i]+nums[j]+nums[k] > 0 {
				k--
			}
			if nums[i]+nums[j]+nums[k] < 0 {
				j++
			}
		}
		preINum = nums[i]
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
