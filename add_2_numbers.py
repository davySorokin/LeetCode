# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        dummy_head = ListNode(0)  # A dummy head to ease the handling of the linked list
        current = dummy_head
        carry = 0  # This will store the carry of the addition

        # Loop through lists l1 and l2 until you reach the end of both
        while l1 is not None or l2 is not None:
            # If we've reached the end of l1 or l2, use 0 as the value for the addition
            x = l1.val if l1 is not None else 0
            y = l2.val if l2 is not None else 0

            # Calculate the sum of the current digits, plus any carry from the previous iteration
            sum = carry + x + y
            carry = sum // 10  # Update carry for next time
            current.next = ListNode(sum % 10)  # Create a new node with the digit value
            current = current.next  # Move to the next node

            # Move to the next nodes in the lists, if possible
            if l1 is not None: l1 = l1.next
            if l2 is not None: l2 = l2.next

        # If there's a carry left after the last addition, create a new node with this value
        if carry > 0:
            current.next = ListNode(carry)

        return dummy_head.next  # Return the head of the new list, excluding the dummy head
