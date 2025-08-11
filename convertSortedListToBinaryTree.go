func sortedListToBST(head *ListNode) *TreeNode {
    var findMiddle func(start, end *ListNode) *ListNode
    findMiddle = func(start, end *ListNode) *ListNode {
        slow, fast := start, start
        for fast != end && fast.Next != end {
            slow = slow.Next
            fast = fast.Next.Next
        }
        return slow
    }

    var convert func(start, end *ListNode) *TreeNode
    convert = func(start, end *ListNode) *TreeNode {
        if start == end {
            return nil
        }
        mid := findMiddle(start, end)
        node := &TreeNode{Val: mid.Val}
        node.Left = convert(start, mid)
        node.Right = convert(mid.Next, end)
        return node
    }

    return convert(head, nil)
}
