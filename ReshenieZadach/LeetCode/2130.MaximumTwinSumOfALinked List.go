package LeetCode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func PairSum(head *ListNode) int {
	if head.Next.Next == nil {
		return head.Val + head.Next.Val
	}

	result := 0
	xArray := make([]int, 0, 100002)
	for i := 0; i < 100002; i++ {
		xArray = append(xArray, head.Val)

		if head.Next == nil {
			i = 100002
		}
		head = head.Next
	}
	for i := 0; i < len(xArray)-1; i++ {
		if result < xArray[i]+xArray[len(xArray)-1-i] {
			result = xArray[i] + xArray[len(xArray)-1-i]
		}
		if result < xArray[i+1]+xArray[len(xArray)-2-i] {
			result = xArray[i+1] + xArray[len(xArray)-2-i]
		}
		i++
	}

	return result
}
