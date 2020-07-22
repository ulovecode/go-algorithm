package leetcode

//给定一个二维网格和一个单词，找出该单词是否存在于网格中。
//
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
//
//
//
// 示例:
//
// board =
//[
//  ['A','B','C','E'],
//  ['S','F','C','S'],
//  ['A','D','E','E']
//]
//
//给定 word = "ABCCED", 返回 true
//给定 word = "SEE", 返回 true
//给定 word = "ABCB", 返回 false
//
//
//
// 提示：
//
//
// board 和 word 中只包含大写和小写英文字母。
// 1 <= board.length <= 200
// 1 <= board[i].length <= 200
// 1 <= word.length <= 10^3
//
// Related Topics 数组 回溯算法
// 👍 483 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func exist(board [][]byte, word string) bool {
	for i := range board {
		for j := range board[i] {
			if isExistWord(i, j, board, word) {
				return true
			}
		}
	}
	return false
}

var positions = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func isExistWord(i int, j int, board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	if (j < 0 || i < 0 || i >= len(board) || j >= len(board[0])) || board[i][j] != word[0] {
		return false
	}
	temp := board[i][j]
	board[i][j] = 0
	var flag bool
	if len(word) != 1 {
		for _, p := range positions {
			x := j + p[0]
			y := i + p[1]
			if isExistWord(y, x, board, word[1:]) {
				flag = true
				break
			}
		}
	} else {
		flag = true
	}
	board[i][j] = temp
	return flag
}

//leetcode submit region end(Prohibit modification and deletion)
