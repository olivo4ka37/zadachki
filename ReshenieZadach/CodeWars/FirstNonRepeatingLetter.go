package CodeWars

import (
	"strings"
)

func FirstNonRepeating(str string) string {
	var result string
	if len(str) == 0 {
		return ""
	}

	var proverka bool = true

	for i := 0; i < len(str); i++ {
		for c := 0; c < len(str); c++ {
			if strings.ToLower(string(str[i])) == strings.ToLower(string(str[c])) && c != i {
				proverka = false
			}
		}

		if proverka == true {
			result = string(str[i])
			return result
		}
		proverka = true
	}

	return ""
}
