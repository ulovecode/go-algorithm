//给定一个没有重复数字的序列，返回其所有可能的全排列。
//
// 示例:
//
// 输入: [1,2,3]
//输出:
//[
//  [1,2,3],
//  [1,3,2],
//  [2,1,3],
//  [2,3,1],
//  [3,1,2],
//  [3,2,1]
//]
//
package cn

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	permuteNums(nums, 0, &result)
	return result
}

func permuteNums(nums []int, start int, result *[][]int) {
	if start == len(nums) {
		des := make([]int, len(nums))
		copy(des, nums)
		*result = append(*result, des)
		return
	}

	for i := start; i < len(nums); i++ {
		nums[start], nums[i] = nums[i], nums[start]
		permuteNums(nums, start+1, result)
		nums[start], nums[i] = nums[i], nums[start]
	}
}
