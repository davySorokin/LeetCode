# Do NOT redefine TreeNode or ListNode when using in LeetCode.

class Solution:
    def sortedListToBST(self, head: Optional[ListNode]) -> Optional[TreeNode]:
        def find_middle(start, end):
            slow = fast = start
            while fast != end and fast.next != end:
                slow = slow.next
                fast = fast.next.next
            return slow

        def convert(start, end):
            if start == end:
                return None
            mid = find_middle(start, end)
            node = TreeNode(mid.val)
            node.left = convert(start, mid)
            node.right = convert(mid.next, end)
            return node

        return convert(head, None)
