package leetcode

import "strings"

//给定一个单词列表，我们将这个列表编码成一个索引字符串 S 与一个索引列表 A。
//
// 例如，如果这个列表是 ["time", "me", "bell"]，我们就可以将其表示为 S = "time#bell#" 和 indexes = [0,
// 2, 5]。
//
// 对于每一个索引，我们可以通过从字符串 S 中索引的位置开始读取字符串，直到 "#" 结束，来恢复我们之前的单词列表。
//
// 那么成功对给定单词列表进行编码的最小字符串长度是多少呢？
//
//
//
// 示例：
//
// 输入: words = ["time", "me", "bell"]
//输出: 10
//说明: S = "time#bell#" ， indexes = [0, 2, 5] 。
//
//
//
//
// 提示：
//
//
// 1 <= words.length <= 2000
// 1 <= words[i].length <= 7
// 每个单词都是小写字母 。
//
//

//leetcode submit region begin(Prohibit modification and deletion)
func minimumLengthEncoding(words []string) int {
	m := map[string]struct{}{}
	for _, str := range words {
		backContains(str, m)
	}
	miniLen := 0
	for s := range m {
		miniLen += len(s) + 1
	}
	return miniLen
}
func backContains(word1 string, m map[string]struct{}) {
	var longestStr string
	var smallStr string
	var deleteStr string
	for v := range m {
		longestStr, smallStr = getStrByLength(v, word1, longestStr, smallStr)
		if strings.HasSuffix(longestStr, smallStr) {
			if longestStr == smallStr {
				continue
			}
			m[longestStr] = struct{}{}
			deleteStr = smallStr
		}
	}
	if deleteStr == "" {
		m[word1] = struct{}{}
	} else {
		delete(m, deleteStr)
	}
}

func getStrByLength(v string, word1 string, longestStr string, smallStr string) (string, string) {
	if len(v) > len(word1) {
		longestStr = v
		smallStr = word1
	} else {
		longestStr = word1
		smallStr = v
	}
	return longestStr, smallStr
}

//leetcode submit region end(Prohibit modification and deletion)
