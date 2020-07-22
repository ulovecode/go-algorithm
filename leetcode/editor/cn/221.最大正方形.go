package leetcode

import "math"

//在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。
//
// 示例:
//
// 输入:
//
//1 0 1 0 0
//1 0 1 1 1
//1 1 1 1 1
//1 0 0 1 0
//
//输出: 4
// Related Topics 动态规划

//leetcode submit region begin(Prohibit modification and deletion)
func maximalSquare(matrix [][]byte) int {

	maxLen := 0
	dp := make([]byte, len(matrix[0]))
	for i := range matrix {
		var pre byte
		for j, v := range matrix[i] {
			temp := dp[j]
			if v == 1 {
				if j == 0 {
					dp[j] = 1
					continue
				}
				dp[j] = byte(math.Min(math.Min(float64(dp[j-1]), float64(dp[j])), float64(pre))) + 1
			} else {
				dp[j] = 0
			}
			pre = temp
			if int(dp[i]) > maxLen {
				maxLen = int(dp[i])
			}
		}
	}
	return maxLen * maxLen
}

//leetcode submit region end(Prohibit modification and deletion)
