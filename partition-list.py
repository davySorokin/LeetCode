class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def partition(self, head: Optional[ListNode], x: int) -> Optional[ListNode]:
        # Create two dummy heads for less and greater linked lists
        less_head = ListNode(0)
        greater_head = ListNode(0)
        
        # Pointers to the current position in the less and greater linked lists
        less = less_head
        greater = greater_head
        
        # Iterate over the original list and partition it
        while head:
            if head.val < x:
                less.next = head
                less = less.next
            else:
                greater.next = head
                greater = greater.next
            head = head.next
        
        # Close the greater list
        greater.next = None
        
        # Attach the greater list to the end of the less list
        less.next = greater_head.next
        
        return less_head.next
