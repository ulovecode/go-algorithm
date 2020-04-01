package kmp

func kmp(matchString, templateString string) string {
	// 首先对字符串进行预处理,得到next数组
	nextArray := getNextArray(templateString)

}

func getNextArray(templateString string) []int {

	tLen := len(templateString)
	nextArray := make([]int, tLen)
	nextArray[0] = -1
	k, j := -1, 0
	// j代表后指针，k代表前指针
	for j < tLen-1 {
		// 在开始的位置或者字符匹配就向后移动继续匹配，并且把k的位置赋值给j位置
		if k == -1 || templateString[k] == templateString[j] {
			j++
			k++
			nextArray[j] = k
		} else {
			// 如果匹配失败k位置就回当上一个位置进行匹配
			k = nextArray[k]
		}
	}
	return nextArray
}
