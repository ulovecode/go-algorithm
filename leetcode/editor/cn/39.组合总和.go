package cn

//给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
// candidates 中的数字可以无限制重复被选取。
//
// 说明：
//
//
// 所有数字（包括 target）都是正整数。
// 解集不能包含重复的组合。
//
//
// 示例 1:
//
// 输入: candidates = [2,3,6,7], target = 7,
//所求解集为:
//[
//  [7],
//  [2,2,3]
//]
//
//
// 示例 2:
//
// 输入: candidates = [2,3,5], target = 8,
//所求解集为:
//[
//  [2,2,2,2],
//  [2,3,3],
//  [3,5]
//]
// Related Topics 数组 回溯算法

//leetcode submit region begin(Prohibit modification and deletion)
func combinationSum(candidates []int, target int) [][]int {

	resultSet := make(map[int][]int, 16)
	arr := make([]int, 0)
	combination(0, len(candidates), candidates, target, resultSet, &arr)
	result := make([][]int, len(resultSet))
	index := 0
	for _, v := range resultSet {
		result[index] = v
		index++
	}
	return result
}

func combination(start int, end int, candidates []int, target int, result map[int][]int, arr *[]int) {
	sum := 0
	multiply := 1
	min := 1
	for i, v := range *arr {
		sum += v
		multiply *= v * (i + 1)
		if v > min {
			min = v
		}
	}
	if sum > target {
		return
	}
	if sum == target {
		des := make([]int, len(*arr))
		copy(des, *arr)
		key := multiply + len(*arr)
		if _, ok := result[key]; !ok {
			result[key] = des
		}
	}
	for i := start; i < end; i++ {
		*arr = append(*arr, candidates[i])
		combination(start, end, candidates, target, result, arr)
		*arr = (*arr)[:len(*arr)-1]
	}
}

//leetcode submit region end(Prohibit modification and deletion)
