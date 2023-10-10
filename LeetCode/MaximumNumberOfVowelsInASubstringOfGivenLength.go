package LeetCode

func MaxVowels(s string, k int) int {
	result, currentAmountOfVowels := 0, 0

	for i := 0; i < len(s); i++ {
		if i < k {
			if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
				currentAmountOfVowels++
			}
			if result < currentAmountOfVowels {
				result = currentAmountOfVowels
			}
		} else {
			if s[i-k] == 'a' || s[i-k] == 'e' || s[i-k] == 'i' || s[i-k] == 'o' || s[i-k] == 'u' {
				currentAmountOfVowels--
			}
			if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
				currentAmountOfVowels++
			}
			if result < currentAmountOfVowels {
				result = currentAmountOfVowels
			}
		}
	}

	return result
}
