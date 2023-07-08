package LeetCode

func MergeAlternately(word1 string, word2 string) string {
	var resultString string

	if len(word1) >= len(word2) {
		for i := 0; i < len(word1); i++ {
			resultString = resultString + string(word1[i])
			if i < len(word2) {
				resultString = resultString + string(word2[i])
			}
		}
	} else {
		for i := 0; i < len(word2); i++ {
			if i < len(word1) {
				resultString = resultString + string(word1[i])
			}
			resultString = resultString + string(word2[i])
		}
	}

	return resultString
}
