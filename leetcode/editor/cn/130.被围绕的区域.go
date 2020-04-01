package cn

//给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。
//
// 找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
//
// 示例:
//
// X X X X
//X O O X
//X X O X
//X O X X
//
//
// 运行你的函数后，矩阵变为：
//
// X X X X
//X X X X
//X X X X
//X O X X
//
//
// [["O","X","X","O","X"]
//["X","O","O","X","O"]
//["X","O","X","O","X"]
//["O","X","O","O","O"]
//["X","X","O","X","O"]]

//[["O","X","X","O","X"]
//["X","X","X","X","O"]
//["X","X","X","X","X"]
//["O","X","O","O","O"]
//["X","X","O","X","O"]]

//[["O","X","X","O","X"]
//["X","X","X","X","O"]
//["X","X","X","O","X"]
//["O","X","O","O","O"]
//["X","X","O","X","O"]]
// 解释:
//
// 被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被
//填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
// Related Topics 深度优先搜索 广度优先搜索 并查集

//leetcode submit region begin(Prohibit modification and deletion)
type position struct {
	x int
	y int
}

func solve(board [][]byte) {
	mm := map[position]struct{}{}
	for i := range board {
		for j := range board[i] {
			p := position{x: i, y: j}
			if _, ok := mm[p]; board[i][j] == 'X' && ok {
				continue
			}
			m := map[position]struct{}{}
			if !isInject(board, i, j, m) {
				for k := range m {
					board[k.x][k.y] = 'O'
					mm[k] = struct{}{}
				}
			}
		}
	}
}

func isInject(board [][]byte, i, j int, m map[position]struct{}) bool {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[i]) {
		return false
	}
	p := position{
		x: i,
		y: j,
	}
	if _, ok := m[p]; ok {
		return true
	}
	if board[i][j] == 'X' {
		return true
	}
	board[i][j] = 'X'

	m[p] = struct{}{}
	if isInject(board, i-1, j, m) && isInject(board, i+1, j, m) && isInject(board, i, j-1, m) && isInject(board, i, j+1, m) {
		return true
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
