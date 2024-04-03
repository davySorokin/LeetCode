# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def reverseBetween(self, head: Optional[ListNode], left: int, right: int) -> Optional[ListNode]:
        if not head or left == right:
            return head
        
        dummy = ListNode(0)
        dummy.next = head
        pre = dummy
        
        # Move pre to the node before the left-th node
        for _ in range(left - 1):
            pre = pre.next
        
        # reverse_start will be the first node to be reversed
        reverse_start = pre.next
        then = reverse_start.next
        
        # Start the reversal process
        for _ in range(right - left):
            reverse_start.next = then.next
            then.next = pre.next
            pre.next = then
            then = reverse_start.next
        
        return dummy.next
