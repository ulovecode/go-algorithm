package main

func findLucky(arr []int) int {
	m := make(map[int]int, 0)
	for _, n := range arr {
		if count, ok := m[n]; ok {
			m[n] = count + 1
		} else {
			m[n] = 1
		}
	}
	max := -1
	for k, v := range m {
		if k == v && k > max {
			max = k
		}
	}
	return max
}
