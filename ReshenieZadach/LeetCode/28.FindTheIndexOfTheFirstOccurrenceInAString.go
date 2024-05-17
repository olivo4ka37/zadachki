package LeetCode

func StrStr(haystack string, needle string) int {
	result := -1
	if len(needle) > len(haystack) || len(haystack) == 0 {
		return result
	}

	i, j := 0, 0
	for j < len(needle) && i < len(haystack) {
		if haystack[i] != needle[j] {
			i++
			continue
		} else {
			for j < len(needle) && i+j < len(haystack) && haystack[i+j] == needle[j] {
				j++
			}
			if j == len(needle) {
				result = i
				break
			}
			i++
			j = 0
		}
	}

	return result
}
