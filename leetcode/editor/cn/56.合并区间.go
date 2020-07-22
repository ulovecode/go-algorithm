//给出一个区间的集合，请合并所有重叠的区间。
//
// 示例 1:
//
// 输入: [[1,3],[2,6],[8,10],[15,18]]
//输出: [[1,6],[8,10],[15,18]]
//解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
//
//
// 示例 2:
//
// 输入: [[1,4],[4,5]]
//输出: [[1,5]]
//解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
//

package leetcode

import (
	"math"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	intervalsSort(intervals)
	result := make([][]int, 0)
	var pre []int

	for _, v := range intervals {
		if pre == nil {
			pre = v
		}
		if v[0] <= pre[1] {
			pre[1] = int(math.Max(float64(pre[1]), float64(v[1])))
		} else {
			result = append(result, pre)
			pre = v
		}
	}
	if pre != nil {
		result = append(result, pre)
	}
	return result
}

func intervalsSort(intervals [][]int) {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		}
		return false
	})
}
