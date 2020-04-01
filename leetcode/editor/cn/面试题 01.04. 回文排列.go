package cn

//给定一个字符串，编写一个函数判定其是否为某个回文串的排列之一。
//回文串是指正反两个方向都一样的单词或短语。排列是指字母的重新排列。
//回文串不一定是字典当中的单词。
//

//
//
//
// 示例1：
//
//输入："tactcoa"
//输出：true（排列有"tacocat"、"atcocta"，等等）
//
//

//leetcode submit region begin(Prohibit modification and deletion)
func canPermutePalindrome(s string) bool {
	amap := [300]int{}
	for _, c := range s {
		amap[c]++
	}
	var oneNumber int
	for _, v := range amap {
		if v%2 == 1 {
			oneNumber++
		}
	}
	return oneNumber < 2
}

//leetcode submit region end(Prohibit modification and deletion)
