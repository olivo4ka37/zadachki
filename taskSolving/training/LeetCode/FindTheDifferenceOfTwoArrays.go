package LeetCode

import (
	"sort"
	"strconv"
)

func FindDifference(nums1 []int, nums2 []int) [][]int {
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
	sort.Slice(nums2, func(i, j int) bool {
		return nums2[i] < nums2[j]
	})

	uniqueArray1 := newSet()
	uniqueArray2 := newSet()

	result := make([][]int, 2)

	for i := 0; i < len(nums1); i++ {
		uniqueArray1.add(strconv.Itoa(nums1[i]))
	}

	for i := 0; i < len(nums2); i++ {
		uniqueArray2.add(strconv.Itoa(nums2[i]))
	}

	for i := 0; i < len(nums2); i++ {
		uniqueArray1.remove(strconv.Itoa(nums2[i]))
	}

	for i := 0; i < len(nums1); i++ {
		uniqueArray2.remove(strconv.Itoa(nums1[i]))
	}

	if uniqueArray1.contains(strconv.Itoa(nums1[0])) {
		result[0] = append(result[0], nums1[0])
	}
	if uniqueArray2.contains(strconv.Itoa(nums2[0])) {
		result[1] = append(result[1], nums2[0])
	}

	for i := 1; i < len(nums1); i++ {
		if uniqueArray1.contains(strconv.Itoa(nums1[i])) && nums1[i] != nums1[i-1] {
			result[0] = append(result[0], nums1[i])
		}
	}

	for i := 1; i < len(nums2); i++ {
		if uniqueArray2.contains(strconv.Itoa(nums2[i])) && nums2[i] != nums2[i-1] {
			result[1] = append(result[1], nums2[i])
		}
	}

	return result
}

var exists = struct{}{}

type set struct {
	m map[string]struct{}
}

func newSet() *set {
	s := &set{}
	s.m = make(map[string]struct{})
	return s
}

func (s *set) add(value string) {
	s.m[value] = exists
}

func (s *set) remove(value string) {
	delete(s.m, value)
}

func (s *set) contains(value string) bool {
	_, c := s.m[value]
	return c
}
