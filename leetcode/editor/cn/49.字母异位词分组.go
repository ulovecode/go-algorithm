package leetcode

//给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
//
// 示例:
//
// 输入: ["eat", "tea", "tan", "ate", "nat", "bat"],
//输出:
//[
//  ["ate","eat","tea"],
//  ["nat","tan"],
//  ["bat"]
//]
//
// 说明：
//
//
// 所有输入均为小写字母。
// 不考虑答案输出的顺序。
//
// Related Topics 哈希表 字符串

//leetcode submit region begin(Prohibit modification and deletion)
func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)
	strSum := make(map[int32]int, 0)
	for _, s := range strs {
		var sum int32 = 1
		var addSum int32 = 0
		for _, v := range s {
			sum *= v
			addSum += v
		}
		sum += addSum
		if index, ok := strSum[sum]; ok {
			result[index] = append(result[index], s)
		} else {
			i := len(result)
			strSum[sum] = i
			str := make([]string, 1)
			str[0] = s
			result = append(result, str)
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
