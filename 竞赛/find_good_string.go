package main

import (
	"strings"
)

func findGoodStrings(n int, s1 string, s2 string, evil string) int {
	strBuilder := strings.Builder{}
	ans := 0
	for i := 0; i < n; i++ {
		var originchar = s1[i]
		for ; originchar < s2[i]; originchar = originchar + 1 {
			strBuilder.WriteString(s1[:i])
			strBuilder.WriteString(string(originchar))
			strBuilder.WriteString(s1[i+1:])
			s := strBuilder.String()
			println(s)
			if !strings.Contains(s, evil) {
				ans %= 10 ^ 9 + 7
			}
			strBuilder.Reset()
		}

	}
	return ans
}

func main() {
	findGoodStrings(2, "aa", "da", "b")
}
