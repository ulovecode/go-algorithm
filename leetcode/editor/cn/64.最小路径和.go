//给定一个包含非负整数的 friends x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//
// 说明：每次只能向下或者向右移动一步。
//
// 示例:
//
// 输入:
//[
//  [1,3,1],
//  [1,5,1],
//  [4,2,1]
//]
//输出: 7
//解释: 因为路径 1→3→1→1→1 的总和最小。
//
//

package cn

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[0]))
	}

	for i := range grid {
		for j, v := range grid[i] {
			if i == 0 || j == 0 {
				if i == 0 && j == 0 {
					dp[i][j] = grid[i][j]
					continue
				}
				if i == 0 {
					dp[i][j] = grid[i][j] + dp[i][j-1]
				}
				if j == 0 {
					dp[i][j] = grid[i][j] + dp[i-1][j]
				}
				continue
			}
			num := dp[i-1][j]
			if num > dp[i][j-1] {
				num = dp[i][j-1]
			}
			dp[i][j] = num + v
		}
	}

	return dp[len(grid)-1][len(grid[0])-1]
}
