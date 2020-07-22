package leetcode

//实现一个基本的计算器来计算一个简单的字符串表达式的值。
//
// 字符串表达式仅包含非负整数，+， - ，*，/ 四种运算符和空格 。 整数除法仅保留整数部分。
//
// 示例 1:
//
// 输入: "3+2*2"
//输出: 7
//
//
// 示例 2:
//
// 输入: " 3/2 "
//输出: 1
//
// 示例 3:
//
// 输入: " 3+5 / 2 "
//输出: 5
//
//
// 说明：
//
//
// 你可以假设所给定的表达式都是有效的。
// 请不要使用内置的库函数 eval。
//
// Related Topics 字符串

//leetcode submit region begin(Prohibit modification and deletion)
func calculate(s string) int {
	symbolStack := make([]int32, 0)
	numberStack := make([]int32, 0)
	for _, v := range s {
		if v == ' ' {
			continue
		}
		if v == '+' || v == '-' || v == '*' || v == '/' {
			symbolStack = append(symbolStack, v)
		} else {
			if len(symbolStack) != 0 {
				symbol := symbolStack[len(symbolStack)-1]
				if symbol == '*' {
					v = numberStack[len(numberStack)-1] * (v - '0')
					numberStack = numberStack[:len(numberStack)-1]
					symbolStack = symbolStack[:len(symbolStack)-1]
					v += '0'
				} else if symbol == '/' {
					v = numberStack[len(numberStack)-1] / (v - '0')
					numberStack = numberStack[:len(numberStack)-1]
					symbolStack = symbolStack[:len(symbolStack)-1]
					v += '0'
				}
			}
			numberStack = append(numberStack, v-'0')
		}
	}
	if len(symbolStack) != 0 {
		for i := len(symbolStack) - 1; i >= 0; i-- {
			numberStackLength := len(numberStack)
			if symbolStack[i] == '+' {
				numberStack[numberStackLength-2] = numberStack[numberStackLength-2] + numberStack[numberStackLength-1]
				numberStack = numberStack[:len(numberStack)-1]
			}
			if symbolStack[i] == '-' {
				numberStack[numberStackLength-2] = numberStack[numberStackLength-2] - numberStack[numberStackLength-1]
				numberStack = numberStack[:len(numberStack)-1]
			}

		}
	}
	return int(numberStack[0])
}

//leetcode submit region end(Prohibit modification and deletion)
