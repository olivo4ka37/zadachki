package LeetCode

func OddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	queueOdd, queueEven := make([]int, 0, 10001), make([]int, 0, 10001)

	for i := 0; i < 10001; i++ {
		if i%2 == 0 {
			queueOdd = append(queueOdd, head.Val)
		} else {
			queueEven = append(queueEven, head.Val)
		}

		if head.Next == nil {
			i = 10001
		}
		head = head.Next
	}

	queueOdd = append(queueOdd, queueEven...)
	result := make([]ListNode, len(queueOdd))

	for i := 0; i < len(queueOdd); i++ {
		result[i].Val = queueOdd[i]
		if i != 0 {
			result[i-1].Next = &result[i]
		}
	}

	return &result[0]
}
