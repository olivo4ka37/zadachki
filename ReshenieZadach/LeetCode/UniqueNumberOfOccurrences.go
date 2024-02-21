package LeetCode

import "sort"

func UniqueOccurrences(arr []int) bool {
	setHuet := make([]int, 2001)

	for i := 0; i < len(arr); i++ {
		setHuet[arr[i]+1000]++
	}

	sort.Slice(setHuet, func(i, j int) bool {
		return setHuet[i] > setHuet[j]
	})

	for i := 0; i < len(setHuet); i++ {
		if setHuet[i] == setHuet[i+1] && setHuet[i] != 0 {
			return false
		} else if setHuet[i] == 0 {
			break
		}
	}

	return true
}
