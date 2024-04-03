# Definition for a Node.
class Node:
    def __init__(self, x: int, next: 'Node' = None, random: 'Node' = None):
        self.val = int(x)
        self.next = next
        self.random = random

class Solution:
    def copyRandomList(self, head: 'Optional[Node]') -> 'Optional[Node]':
        if not head:
            return None

        # First pass: make a copy of the nodes and insert them next to their originals
        current = head
        while current:
            new_node = Node(current.val, current.next)
            current.next = new_node
            current = new_node.next
        
        # Second pass: set the random pointers for the copied nodes
        current = head
        while current:
            if current.random:
                current.next.random = current.random.next
            current = current.next.next
        
        # Third pass: restore the original list and extract the copy
        current = head
        copy_head = head.next
        while current:
            temp = current.next
            current.next = temp.next
            current = current.next
            if temp.next:
                temp.next = temp.next.next
        
        return copy_head
