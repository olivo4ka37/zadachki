package LeetCode

import "fmt"

func OddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	result, queueEven := make([]ListNode, 10001), make([]ListNode, 10001)
	addOdd, firstOdd, firstEven, listing := true, true, true, true
	i, j := 0, 0

	for listing {
		if head.Next == nil {
			listing = false
		}
		if addOdd == true {
			newNode := ListNode{
				Val:  head.Val,
				Next: nil,
			}
			if firstOdd != true {
				result[i].Next = &newNode
				i++
			}
			result[i] = newNode
			addOdd = false
			firstOdd = false
		} else {
			newNode := ListNode{
				Val:  head.Val,
				Next: nil,
			}
			if firstEven != true {
				queueEven[j].Next = &newNode
				j++
			}
			queueEven[j] = newNode
			firstEven = false
			addOdd = true
		}
		if head.Next != nil {
			head = head.Next
		}
	}

	listing = true
	for listing {
		if queueEven[0].Next == nil {
			listing = false
		}
		newNode := ListNode{
			Val:  queueEven[0].Val,
			Next: nil,
		}
		result[i].Next = &newNode
		i++
		result[i] = newNode

		if queueEven[0].Next != nil {
			queueEven[0] = *queueEven[0].Next
		}
	}

	for _, x := range result {
		fmt.Println(&x, "11111", x.Val, x.Next.Next)
	}

	return &result[0]
}
