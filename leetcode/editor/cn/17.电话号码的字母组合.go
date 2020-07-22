package leetcode

//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
//
//
// 示例:
//
// 输入："23"
//输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
//
//
// 说明:
//尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
// Related Topics 字符串 回溯算法

//leetcode submit region begin(Prohibit modification and deletion)
func letterCombinations(digits string) []string {
	result := make([]string, 0)
	if digits == "" {
		return result
	}
	strMap := [][]string{
		{},
		{},
		{"a", "b", "c"},
		{"d", "e", "f"},
		{"g", "h", "i"},
		{"j", "k", "l"},
		{"m", "n", "o"},
		{"p", "q", "r", "s"},
		{"t", "u", "v"},
		{"w", "x", "y", "z"},
	}
	dfs(strMap, 0, digits, &result, "")
	return result
}

func dfs(strMap [][]string, start int, digits string, result *[]string, ans string) {
	if start == len(digits) {
		*result = append(*result, ans)
		return
	}

	str := strMap[int(digits[start]-'0')]
	for i := range str {
		temp := ans
		ans += str[i]
		dfs(strMap, start+1, digits, result, ans)
		ans = temp
	}
}

//leetcode submit region end(Prohibit modification and
// deletion)
