//给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。
//
// 示例 1:
//
// 输入:
//11110
//11010
//11000
//00000
//
//输出: 1
//
//
// 示例 2:
//
// 输入:
//11000
//11000
//00100
//00011
//
//输出: 3
//
//

package leetcode

func numIslands(grid [][]byte) int {
	result := 0
	for i := range grid {
		for j, v := range grid[i] {
			if v == 1 {
				landSearch(grid, i, j)
				result++
			}
		}
	}
	return result
}

func landSearch(grid [][]byte, i int, j int) {

}
