//给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
//
// 示例 1：
//
// 输入: "babad"
//输出: "bab"
//注意: "aba" 也是一个有效答案。
//
//
// 示例 2：
//
// 输入: "cbbd"
//输出: "bb"
//
//
package cn

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	length := len(s)
	var dp = make([][]bool, length)
	for i := range dp {
		dp[i] = make([]bool, length)
	}
	begin, end := 0, 0
	for i := length - 1; i >= 0; i-- {
		for j := i; j < length; j++ {
			dp[i][j] = (j-i < 2 || dp[i+1][j-1]) && s[i] == s[j]
			if dp[i][j] && j-i > end-begin {
				begin = i
				end = j
			}
		}
	}
	return s[begin : end+1]
}
