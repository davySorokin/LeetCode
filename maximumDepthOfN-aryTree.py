# Definition for a Node.
class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children if children is not None else []

class Solution:
    def maxDepth(self, root: 'Node') -> int:
        if not root:
            return 0
        return 1 + max((self.maxDepth(child) for child in root.children), default=0)
