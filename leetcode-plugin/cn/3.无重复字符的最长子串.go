//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
// 示例 1:
//
// 输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
//
// 示例 2:
//
// 输入: "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
//
// 示例 3:
//
// 输入: "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
//
package cn

func lengthOfLongestSubstring(s string) int {
	var slidingWindow = make(map[int32]int, len(s))
	i, ans := 0, 0
	for j, v := range s {
		if _, ok := slidingWindow[v]; ok {
			// 出现重复字符，更新窗口左下标的值，往右滑动
			if i < slidingWindow[v] {
				i = slidingWindow[v]
			}
		}
		// 检查当前是否是最大答案
		if ans < j-i+1 {
			ans = j - i + 1
		}
		//标记当前字符所在的最新下标,更新到目标字符的下一个位置，直接跳过重复字符
		slidingWindow[v] = j + 1
	}
	return ans
}
