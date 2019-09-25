package cn

import (
	"strings"
)

//给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
//
// 示例 1:
//
//
//输入: "Let's take LeetCode contest"
//输出: "s'teL ekat edoCteeL tsetnoc"
//
//
// 注意：在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。
// Related Topics 字符串

//leetcode submit region begin(Prohibit modification and deletion)
func reverseWords(s string) string {
	var Reverse = func(s string) []byte {
		r := []byte(s)
		for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		return r
	}
	words := strings.Split(s, " ")
	builder := new(strings.Builder)
	for i := 0; i < len(words); i++ {
		builder.Write(Reverse(words[i]))
		if i != len(words)-1 {
			builder.Write([]byte(" "))
		}
	}
	return builder.String()
}

//leetcode submit region end(Prohibit modification and deletion)
