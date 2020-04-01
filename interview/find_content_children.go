package interview

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	ans, index := 0, 0
	for i := 0; i < len(g); i++ {
		if index >= len(s) {
			break
		}
		for index < len(s) && g[i] > s[index] {
			index++
		}
		if index < len(s) && g[i] <= s[index] {
			ans++
			index++
		}
	}
	return ans
}
