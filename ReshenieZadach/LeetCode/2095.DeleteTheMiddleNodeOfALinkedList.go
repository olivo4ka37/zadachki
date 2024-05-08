package LeetCode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func DeleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	} else if head.Next.Next == nil {
		head.Next = nil
		return head
	}

	nodeForSkip := head
	result := nodeForSkip
	skip := false
	head = head.Next
	for head != nil {
		if skip && head.Next != nil {
			nodeForSkip = nodeForSkip.Next
			skip = false
		} else {
			skip = true
		}
		head = head.Next
	}
	nodeForSkip.Next = nodeForSkip.Next.Next

	return result
}
