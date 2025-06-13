/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val  int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	// Empty list or single node can’t form a cycle.
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next

	for fast != nil && fast.Next != nil {
		if slow == fast {          // pointers meet → cycle exists
			return true
		}
		slow = slow.Next          // move 1 step
		fast = fast.Next.Next     // move 2 steps
	}

	// fast hit the end → no cycle
	return false
}
