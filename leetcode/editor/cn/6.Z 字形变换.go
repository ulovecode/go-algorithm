//将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。
//
// 比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：
//
// L   C   I   R
//E T O E S I I G
//E   D   H   N
//
//
// 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。
//
// 请你实现这个将字符串进行指定行数变换的函数：
//
// string convert(string s, int numRows);
//
// 示例 1:
//
// 输入: s = "LEETCODEISHIRING", numRows = 3
//输出: "LCIRETOESIIGEDHN"
//
//
// 示例 2:
//
// 输入: s = "LEETCODEISHIRING", numRows = 4
//输出: "LDREOEIIECIHNTSG"
//解释:
//
//L     D     R
//E   O E   I I
//E C   I H   N
//T     S     G
//

package leetcode

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	ints := make([][]uint8, numRows)
	for i := 0; i < numRows; i++ {
		ints = append(ints, make([]uint8, 0))
	}
	var findUp bool
	numRow := 0
	for i := 0; i < len(s); i++ {
		ints[numRow] = append(ints[numRow], s[i])
		if findUp {
			numRow--
			if 0 == numRow {
				findUp = false
				numRow = 0
			}
			if numRow < 0 {
				numRow = numRows - 1
			}
		} else {
			numRow++
			if numRows == numRow {
				numRow = numRows - 2
				findUp = true
			}
		}

	}
	result := ""
	for _, s1 := range ints {
		for _, s2 := range s1 {
			result += string(s2)
		}
	}

	return result
}
