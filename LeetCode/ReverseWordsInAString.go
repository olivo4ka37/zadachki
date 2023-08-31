package LeetCode

func ReverseWords(s string) string {
	result, word := "", ""

	for i := 0; i < len(s); i++ {
		// i == len(s)-1 Checking end of the string, to add last word.
		// string(s[i]) == " " && string(s[i+1]) != " " && word != "" Checking spaces between words and add last word,
		// also ignoring multiple spaces.
		// after all of it add symbol of word if we have one.
		if i == len(s)-1 {
			// Check necessity to add last symbol in var word.
			if string(s[i]) != " " {
				word += string(s[i])
			}
			result = word + result
		} else if string(s[i]) == " " && string(s[i+1]) != " " && word != "" {
			result = " " + word + result
			word = ""
		} else if string(s[i]) != " " {
			word += string(s[i])
		}
	}

	return result
}
