package manacher

import "math"

func manacherSring(s string) []rune {
	runes := []rune(s)
	res := make([]rune, len(s)*2+1)
	for index, i := 0, 0; i < len(res); i++ {
		if (i & 1) == 0 {
			res[i] = '#'
		} else {
			res[i] = runes[index]
			index++
		}
	}
	return res
}

func maxLcpsLength(s string) int {
	if s == "" {
		return 0
	}
	runes := manacherSring(s)
	pArr := make([]int, 0, len(runes))
	var index, pR, max = -1, -1, math.MinInt32
	for i := 0; i < len(runes); i++ {
		pArr[i] = 1
		if pR > i {
			pArr[i] = pArr[2*index-i]
			if pArr[2*index-i] > pR-i {
				pArr[i] = pR - i
			}
		}
		for ; i+pArr[i] < len(runes) && i-pArr[i] > -1 && runes[i+pArr[i]] == runes[i-pArr[i]]; pArr[i]++ {
		}
		if i+pArr[i] > pR {
			pR = i + pArr[i]
			index = i
		}
		max = int(math.Max(float64(max), float64(pArr[i])))
	}
	return max - 1
}
