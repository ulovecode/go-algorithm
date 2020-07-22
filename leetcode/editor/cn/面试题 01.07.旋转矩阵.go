package leetcode

//给你一幅由 N × N 矩阵表示的图像，其中每个像素的大小为 4 字节。请你设计一种算法，将图像旋转 90 度。
//
// 不占用额外内存空间能否做到？
//
//
//
// 示例 1:
//
// 给定 matrix =
//[
//  [1,2,3],
//  [4,5,6],
//  [7,8,9]
//],
//
//原地旋转输入矩阵，使其变为:
//[
//  [7,4,1],
//  [8,5,2],
//  [9,6,3]
//]
//
//
// 示例 2:
//
// 给定 matrix =
//[
//  [ 5, 1, 9,11],
//  [ 2, 4, 8,10],
//  [13, 3, 6, 7],
//  [15,14,12,16]
//],
//
//原地旋转输入矩阵，使其变为:
//[
//  [15,13, 2, 5],
//  [14, 3, 4, 1],
//  [12, 6, 8, 9],
//  [16, 7,10,11]
//]
//
// Related Topics 数组

//leetcode submit region begin(Prohibit modification and deletion)
func rotateArr(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	sr, sc, er, ec := 0, 0, len(matrix)-1, len(matrix[0])-1
	for sc < ec {
		rotateArrEdge(matrix, sr, er, sc, ec)
		sr++
		sc++
		er--
		ec--
	}
}

func rotateArrEdge(matrix [][]int, sr int, er int, sc int, ec int) {
	for i := 0; i < ec-sc; i++ {
		temp := matrix[sr][sc+i]
		matrix[sr][sc+i] = matrix[er-i][sc]
		matrix[er-i][sc] = matrix[er][ec-i]
		matrix[er][ec-i] = matrix[sr+i][ec]
		matrix[sr+i][ec] = temp
	}
}

//leetcode submit region end(Prohibit modification and deletion)
