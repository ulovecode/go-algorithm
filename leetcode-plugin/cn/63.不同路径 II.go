//一个机器人位于一个 friends x n 网格的左上角 （起始点在下图中标记为“Start” ）。
//
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
//
// 现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
//
//
//
// 网格中的障碍物和空位置分别用 1 和 0 来表示。
//
// 说明：friends 和 n 的值均不超过 100。
//
// 示例 1:
//
// 输入:
//[
//  [0,0,0],
//  [0,1,0],
//  [0,0,0]
//]
//输出: 2
//解释:
//3x3 网格的正中间有一个障碍物。
//从左上角到右下角一共有 2 条不同的路径：
//1. 向右 -> 向右 -> 向下 -> 向下
//2. 向下 -> 向下 -> 向右 -> 向右
//
//
package cn

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n := len(obstacleGrid)
	m := len(obstacleGrid[0])
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	dp[0][0] = 1
	for i := 1; i < n; i++ {
		if obstacleGrid[i][0] != 1 {
			if dp[i-1][0] == 1 {
				dp[i][0] = 1
			}
		}
	}
	for j := 1; j < m; j++ {
		if obstacleGrid[0][j] != 1 {
			if dp[0][j-1] == 1 {
				dp[0][j] = 1
			}

		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			}
		}
	}
	return dp[n-1][m-1]
}
