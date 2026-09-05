/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 // 0 1 2 3
 // 1 0 2 3

 // iter 1
 // node = 1 -> 0
 // stepNode = 2
 // head = 0


func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	
	var tail = &ListNode{Val: head.Val, Next: nil}

	newHead := head
	for {
		if newHead.Next == nil {
			break
		}

		newTail := &ListNode{Val: newHead.Next.Val, Next: tail}
		tail = newTail
		newHead = newHead.Next
	}

	return tail
}
