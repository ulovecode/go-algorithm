//给定一个可包含重复数字的序列，返回所有不重复的全排列。
//
// 示例:
//
// 输入: [1,1,2]
//输出:
//[
//  [1,1,2],
//  [1,2,1],
//  [2,1,1]
//]
//

package cn

//func permuteUnique(nums []int) [][]int {
//	result := make([][]int, 0)
//	permuteUniqueNums(nums, 0, &result)
//	return result
//}
//
////[[1,1,2,2],[1,2,1,2],[1,2,2,1],[2,1,1,2],[2,1,2,1],[2,2,1,1]]
//
//func permuteUniqueNums(nums []int, start int, result *[][]int) {
//	if start == len(nums) {
//		des := make([]int, len(nums))
//		copy(des, nums)
//		*result = append(*result, des)
//		return
//	}
//
//	for i := start; i < len(nums); i++ {
//		if start != i && nums[start] == nums[i] {
//			continue
//		}
//		nums[start], nums[i] = nums[i], nums[start]
//		permuteUniqueNums(nums, start+1, result)
//		nums[start], nums[i] = nums[i], nums[start]
//	}
//}
