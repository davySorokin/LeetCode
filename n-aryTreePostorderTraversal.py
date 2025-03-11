# Definition for a Node.
from typing import List, Optional

class Node:
    def __init__(self, val: Optional[int] = None, children: Optional[List['Node']] = None):
        self.val = val
        self.children = children if children is not None else []

class Solution:
    def postorder(self, root: 'Node') -> List[int]:
        if not root:
            return []
        
        result = []
        
        def traverse(node):
            for child in node.children:
                traverse(child)
            result.append(node.val)
        
        traverse(root)
        return result
