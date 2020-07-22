package leetcode

import (
	"reflect"
	"strings"
)

//给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。
//
// 注意：
//
//
// num1 和num2 的长度都小于 5100.
// num1 和num2 都只包含数字 0-9.
// num1 和num2 都不包含任何前导零。
// 你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。
//
// Related Topics 字符串

//leetcode submit region begin(Prohibit modification and deletion)
func addStrings(num1 string, num2 string) string {
	builder := new(strings.Builder)

	i := len(num1) - 1
	j := len(num2) - 1
	var carryNum = 0
	println(reflect.ValueOf(carryNum).Type())
	for i >= 0 || j >= 0 {
		if i >= 0 && j < 0 {
			num := int(num1[i] - '0')
			num += carryNum
			if num >= 10 {
				num -= 10
				carryNum = 1
			} else {
				carryNum = 0
			}
			builder.Write([]byte(string(num + '0')))
			i--
			continue
		}
		if j >= 0 && i < 0 {
			num := int(num2[j] - '0')
			num += carryNum
			if num >= 10 {
				num -= 10
				carryNum = 1
			} else {
				carryNum = 0
			}
			builder.Write([]byte(string(num + '0')))
			j--
			continue
		}
		num := int(num1[i] - '0' + num2[j] - '0')
		num += carryNum
		if num >= 10 {
			num -= 10
			carryNum = 1
		} else {
			carryNum = 0
		}
		builder.Write([]byte(string(num + '0')))
		i--
		j--
	}
	if carryNum != 0 {
		builder.Write([]byte(string(carryNum + '0')))
	}

	return Reverse(builder.String())
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

//leetcode submit region end(Prohibit modification and deletion)
