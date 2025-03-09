# Definition for a Node.
from typing import List, Optional

class Node:
    def __init__(self, val: Optional[int] = None, children: Optional[List['Node']] = None):
        self.val = val
        self.children = children if children is not None else []

class Solution:
    def preorder(self, root: 'Node') -> List[int]:
        result = []
        
        def dfs(node):
            if not node:
                return
            result.append(node.val)  # Visit the node
            for child in node.children:
                dfs(child)  # Recur for each child
        
        dfs(root)
        return result
