package LeetCode

import (
	"sort"
)

func CloseStrings(word1 string, word2 string) bool {
	alphabetArray1, alphabetArray2 := make([]int, 26), make([]int, 26)
	alphabetMap1, alphabetMap2 := make([]bool, 26), make([]bool, 26)

	for _, x := range word1 {
		alphabetArray1[x-97]++
		alphabetMap1[x-97] = true
	}

	for _, x := range word2 {
		alphabetArray2[x-97]++
		alphabetMap2[x-97] = true
	}

	sort.Slice(alphabetArray1, func(i, j int) bool {
		return alphabetArray1[i] > alphabetArray1[j]
	})

	sort.Slice(alphabetArray2, func(i, j int) bool {
		return alphabetArray2[i] > alphabetArray2[j]
	})

	for i := 0; i < 26; i++ {
		if alphabetArray1[i] != alphabetArray2[i] {
			return false
		} else if alphabetMap1[i] != alphabetMap2[i] {
			return false
		}
	}

	return true
}
