# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        dummy = ListNode(0, head)  # Create a dummy node to handle edge cases
        first = dummy
        second = dummy

        # Advance first n+1 steps ahead of second
        for _ in range(n + 1):
            first = first.next
        
        # Move first to the end, maintaining the gap
        while first:
            first = first.next
            second = second.next
        
        # Skip the nth node from the end
        second.next = second.next.next
        
        return dummy.next
