class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def deleteDuplicates(self, head: Optional[ListNode]) -> Optional[ListNode]:
        dummy = ListNode(0)  # Create a dummy node that precedes the list
        dummy.next = head
        
        prev = dummy  # Set the previous pointer to the dummy node
        while head:
            # If we find a sequence of nodes with the same value...
            if head.next and head.val == head.next.val:
                # Skip all nodes with the current value
                while head.next and head.val == head.next.val:
                    head = head.next
                prev.next = head.next  # Remove the duplicate nodes
            else:
                prev = prev.next  # Move the prev pointer to the next node
            head = head.next  # Continue to the next node

        return dummy.next
