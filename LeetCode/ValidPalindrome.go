package LeetCode

import (
	"regexp"
	"strings"
)

func IsPalindrome(s string) bool {
	s = stripRegex(s)
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToLower(s)

	if len(s) == 0 {
		return true
	}

	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func stripRegex(str string) string {
	reg := regexp.MustCompile("[^a-zA-Z0-9 ]+")
	return reg.ReplaceAllString(str, "")
}
