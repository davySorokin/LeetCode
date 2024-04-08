# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def getMinimumDifference(self, root: TreeNode) -> int:
        self.prev = float('-inf')
        self.min_diff = float('inf')
        
        def in_order_traverse(node):
            if not node:
                return
            in_order_traverse(node.left)
            self.min_diff = min(self.min_diff, node.val - self.prev)
            self.prev = node.val
            in_order_traverse(node.right)
            
        in_order_traverse(root)
        return self.min_diff
