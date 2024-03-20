# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def mergeInBetween(self, list1: ListNode, a: int, b: int, list2: ListNode) -> ListNode:
        # Find the (a-1)th node
        current = list1
        for i in range(a - 1):
            current = current.next
        nodeBeforeA = current
        
        # Find the bth node
        for i in range(b - a + 1):
            current = current.next
        nodeAfterB = current.next
        
        # Connect the (a-1)th node to the head of list2
        nodeBeforeA.next = list2
        
        # Find the last node of list2
        while list2.next:
            list2 = list2.next
        
        # Connect the last node of list2 to the node after bth
        list2.next = nodeAfterB
        
        return list1
