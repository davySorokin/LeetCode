/* Declared already
// Definition for singly-linked list.
type ListNode struct {
    Val  int
    Next *ListNode
}
*/

func insertionSortList(head *ListNode) *ListNode {
    // Base case: if head is nil or only one node, it's already sorted
    if head == nil || head.Next == nil {
        return head
    }

    // Initialize a new sorted list
    sorted := &ListNode{} // Dummy node for ease of use

    // Iterate through the original list
    current := head
    for current != nil {
        // At each step, take the current node and find its place in the sorted list
        prev := sorted
        next := current.Next

        // Find the place to insert the current node in the sorted list
        for prev.Next != nil && prev.Next.Val < current.Val {
            prev = prev.Next
        }

        // Insert the current node
        current.Next = prev.Next
        prev.Next = current

        // Move to the next node in the original list
        current = next
    }

    // The sorted list is after the dummy node
    return sorted.Next
}
