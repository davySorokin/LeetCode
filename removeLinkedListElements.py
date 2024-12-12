# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def removeElements(self, head: ListNode, val: int) -> ListNode:

        dummy = ListNode(-1)
        dummy.next = head
        current = dummy
        
        # Traverse the list
        while current.next is not None:
            if current.next.val == val:
                # Remove the node with the value
                current.next = current.next.next
            else:
                current = current.next
        
        # Return the new head
        return dummy.next
