package leetcode

//给定一个 n × n 的二维矩阵表示一个图像。
//
// 将图像顺时针旋转 90 度。
//
// 说明：
//
// 你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。
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
func rotate(matrix [][]int) {
	tr := 0
	tc := 0
	dr := len(matrix) - 1
	dc := len(matrix[0]) - 1
	for tr < dr {
		rotateEdge(matrix, tr, tc, dr, dc)
		tr++
		tc++
		dr--
		dc--
	}
}

func rotateEdge(matrix [][]int, tr int, tc int, dr int, dc int) {
	times := dc - tc
	for i := 0; i < times; i++ {
		tmp := matrix[tr][tc+i]
		matrix[tr][tc+i] = matrix[dr-i][tc]
		matrix[dr-i][tc] = matrix[dr][dc-i]
		matrix[dr][dc-i] = matrix[tr+i][dc]
		matrix[tr+i][dc] = tmp
	}
}
