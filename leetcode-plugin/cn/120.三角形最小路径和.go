//给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
//
// 例如，给定三角形：
//
// [
//     [2],
//    [3,4],
//   [6,5,7],
//  [4,1,8,3]
//]
//
//
// 自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
//
// 说明：
//
// 如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。
//

package cn

import "math"

//func minimumTotal(triangle [][]int) int {
//	dp := make([][]int, len(triangle))
//	for i := range dp {
//		dp[i] = make([]int, len(triangle))
//	}
//	for i := 0; i < len(triangle); i++ {
//		for j := 0; j <= i; j++ {
//			if (j == 0 || i == 0) || i == j {
//				if i == 0 && j == 0 {
//					dp[i][j] = triangle[i][j]
//					continue
//				}
//				if j == 0 {
//					dp[i][j] = triangle[i][j] + dp[i-1][j]
//					continue
//				}
//				if j == i {
//					dp[i][j] = triangle[i][j] + dp[i-1][j-1]
//					continue
//				}
//			}
//			dp[i][j] = int(math.Min(float64(dp[i-1][j-1]), float64(dp[i-1][j]))) + triangle[i][j]
//		}
//	}
//
//	min := math.MaxInt32
//	for _,v := range dp[len(triangle)-1] {
//		if min > v {
//			min = v
//		}
//	}
//	return min
//}
func minimumTotal(triangle [][]int) int {
	dp := make([]int, len(triangle))
	for i := 0; i < len(triangle); i++ {
		for j := i; j >= 0; j-- {
			if (j == 0 || i == 0) || i == j {
				if i == 0 && j == 0 {
					dp[j] = triangle[i][j]
					continue
				}
				if j == 0 {
					dp[j] = triangle[i][j] + dp[j]
					continue
				}
				if j == i {
					dp[j] = triangle[i][j] + dp[j-1]
					continue
				}
			}
			dp[j] = int(math.Min(float64(dp[j-1]), float64(dp[j]))) + triangle[i][j]
		}
	}

	min := math.MaxInt32
	for _, v := range dp {
		if min > v {
			min = v
		}
	}
	return min
}
