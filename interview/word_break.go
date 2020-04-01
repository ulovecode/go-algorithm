package interview

import "strings"

func wordBreak(s string, wordDict []string) bool {
	if s == "" {
		return true
	}
	if len(wordDict) == 0 {
		return false
	}
	for i := 0; i < len(wordDict); i++ {
		index := strings.Index(s, wordDict[i])
		if index != -1 {
			if wordBreak(s[:index], wordDict) && wordBreak(s[index+len(wordDict[i]):], wordDict) {
				return true
			}
		} else {
			if wordBreak(s, wordDict[1:]) {
				return true
			}
		}
	}
	return false
}
