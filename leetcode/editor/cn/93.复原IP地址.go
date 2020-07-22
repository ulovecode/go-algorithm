package leetcode

import "strings"

//给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。
//
// 示例:
//
// 输入: "25525511135"
//输出: ["255.255.11.135", "255.255.111.35"]
// Related Topics 字符串 回溯算法

//leetcode submit region begin(Prohibit modification and deletion)
func restoreIpAddresses(s string) []string {
	result := make([]string, 0)
	backtrack(make([]string, 0), s, &result)
	return result
}

func backtrack(putStr []string, s string, result *[]string) {
	resultLen := len(putStr)
	if resultLen == 4 {
		if len(s) == 0 {
			*result = append(*result, putStr[0]+"."+putStr[1]+"."+putStr[2]+"."+putStr[3])
		}
		return
	}
	for i := 1; i <= 3; i++ {
		if len(s) < i {
			return
		}
		str := s[0:i]
		strLen := len(str)
		if strLen == 3 && strings.Compare(str, "255") > 0 {
			return
		}
		if strLen > 1 && s[0] == '0' {
			return
		}
		putStr = append(putStr, str)
		backtrack(putStr, s[i:], result)
		putStr = putStr[:len(putStr)-1]

	}
}

//leetcode submit region end(Prohibit modification and deletion)
