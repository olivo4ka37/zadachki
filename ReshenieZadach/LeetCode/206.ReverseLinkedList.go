package LeetCode

import "fmt"

func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	array := make([]int, 0)

	listing := true
	for listing {
		if head.Next == nil {
			listing = false
		}
		array = append(array, head.Val)
		if head.Next != nil {
			head = head.Next
		}
	}

	result := make([]ListNode, len(array))

	for i := len(array) - 1; i >= 0; i-- {
		result[len(array)-1-i].Val = array[i]
		if i != len(array)-1 {
			result[len(array)-(i+2)].Next = &result[len(array)-1-i]
		}
	}

	for i := 0; i < len(result)-1; i++ {
		fmt.Println(result[i])
	}

	return &result[0]
}

type ListNode struct {
	Val  int
	Next *ListNode
}
