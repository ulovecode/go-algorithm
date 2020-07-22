package leetcode

//给定一个以字符串表示的非负整数 num，移除这个数中的 k 位数字，使得剩下的数字最小。
//
// 注意:
//
//
// num 的长度小于 10002 且 ≥ k。
// num 不会包含任何前导零。
//
//
// 示例 1 :
//
//
//输入: num = "1432219", k = 3
//输出: "1219"
//解释: 移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219。
//
//
// 示例 2 :
//
//
//输入: num = "10200", k = 1
//输出: "200"
//解释: 移掉首位的 1 剩下的数字为 200. 注意输出不能有任何前导零。
//
//
// 示例 3 :
//
//
//输入: num = "10", k = 2
//输出: "0"
//解释: 从原数字移除所有的数字，剩余为空就是0。
//
// Related Topics 栈 贪心算法

//leetcode submit region begin(Prohibit modification and deletion)

//func removeKdigits(num string, k int) string {
//	minStr(num, k)
//	i := 0
//	for ; i < len(minString) && minString[i] == '0'; i++ {
//	}
//	minString = minString[i:]
//	if minString == "" {
//		minString = "0"
//	}
//	return minString
//}
//
//var minString string
//
//func minStr(num string, k int) {
//	if k == 0 {
//		if minString == "" {
//			minString = num
//		} else {
//			if minString > num {
//				minString = num
//			}
//		}
//		return
//	}
//	newNum := num
//	for i := range num {
//		minStr(newNum[:i]+newNum[i+1:], k-1)
//	}
//}

func removeKdigits(num string, k int) string {
	tagetL := len(num) - k
	if tagetL <= 0 {
		return "0"
	}
	i := 1
	for len(num) != tagetL && i < len(num) {
		for i > 0 && len(num) != tagetL && len(num) > 1 && num[i] < num[i-1] {
			num = num[:i-1] + num[i:]
			i--
		}
		i++
	}
	for len(num) > tagetL {
		num = num[:tagetL]
	}
	i = 0
	for ; i < len(num) && num[i] == '0'; i++ {
	}
	num = num[i:]
	if len(num) == 0 {
		num = "0"
	}
	return num
}

//leetcode submit region end(Prohibit modification and deletion)
