package LeetCode

import "strings"

//For two strings s and t, we say "t divides s" if and only if s = t + ... + t (i.e., t is concatenated with itself one or more times).
//Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.

// Описать функцию которая будет проверять сначала схожесть строк по длине предполагаемого делителя, а потом делимость обоих строк,
// при успешной проверке записывает в результат. Заканчивается когда строки начинают разнится или делимость не прошла успешно.
func GcdOfStrings(str1 string, str2 string) string {
	result := ""
	if len(str1) >= len(str2) {
		for i := 0; i <= len(str2); i++ {
			if str1[:i] == str2[:i] && strings.ReplaceAll(str1, str1[:i], "") == "" && strings.ReplaceAll(str2, str1[:i], "") == "" {
				result = str1[:i]
			}
		}
	} else {
		for i := 0; i <= len(str1); i++ {
			if str1[:i] == str2[:i] && strings.ReplaceAll(str1, str2[:i], "") == "" && strings.ReplaceAll(str2, str2[:i], "") == "" {
				result = str2[:i]
			}
		}
	}

	return result
}
