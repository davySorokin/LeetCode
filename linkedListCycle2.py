# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

from typing import Optional

class Solution:
    def detectCycle(self, head: Optional["ListNode"]) -> Optional["ListNode"]:
        slow, fast = head, head
        # phase 1: find meeting point if a cycle exists
        while fast and fast.next:
            slow = slow.next
            fast = fast.next.next
            if slow is fast:
                # phase 2: move one pointer to head; advance both 1-step to find cycle start
                slow = head
                while slow is not fast:
                    slow = slow.next
                    fast = fast.next
                return slow
        return None  # no cycle
