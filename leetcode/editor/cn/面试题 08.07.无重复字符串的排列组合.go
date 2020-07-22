package leetcode

//无重复字符串的排列组合。编写一种方法，计算某字符串的所有排列组合，字符串每个字符均不相同。
//
// 示例1:
//
//
// 输入：S = "qwe"
// 输出：["qwe", "qew", "wqe", "weq", "ewq", "eqw"]
//
//
// 示例2:
//
//
// 输入：S = "ab"
// 输出：["ab", "ba"]
//
//
// 提示:
//
//
// 字符都是英文字母。
// 字符串长度在[1, 9]之间。
//
// Related Topics 回溯算法

//leetcode submit region begin(Prohibit modification and deletion)
func permutation(S string) []string {
	res := make([]string, 0)
	permutationNext([]rune(S), 0, len(S)-1, &res)
	return res
}

func permutationNext(s []rune, i int, end int, res *[]string) {
	if i > end {
		return
	}
	*res = append(*res, string(s))
	for j := i; j <= end; j++ {
		for k := j + 1; k <= end; k++ {
			s[j], s[k] = s[k], s[j]
			permutationNext(s, j+1, end, res)
			s[j], s[k] = s[k], s[j]
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
