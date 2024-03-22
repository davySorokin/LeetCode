# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def isPalindrome(self, head: ListNode) -> bool:
        # Helper function to reverse a linked list.
        def reverse_list(node):
            prev = None
            while node:
                next_node = node.next
                node.next = prev
                prev = node
                node = next_node
            return prev

        # Find the middle of the linked list.
        slow = fast = head
        while fast and fast.next:
            slow = slow.next
            fast = fast.next.next
        
        # Reverse the second half.
        second_half = reverse_list(slow)
        
        # Compare the first half and the reversed second half.
        first_half = head
        palindrome = True
        while second_half:
            if first_half.val != second_half.val:
                palindrome = False
                break
            first_half = first_half.next
            second_half = second_half.next
        
        # Optional: Restore the list to its original state by reversing the second half back.
        # reverse_list(slow)
        
        return palindrome
