# Definition for a binary tree node.
from typing import Optional

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def findTilt(self, root: Optional[TreeNode]) -> int:
        self.total_tilt = 0

        def subtree_sum(node):
            if not node:
                return 0
            left_sum = subtree_sum(node.left)
            right_sum = subtree_sum(node.right)
            self.total_tilt += abs(left_sum - right_sum)
            return node.val + left_sum + right_sum

        subtree_sum(root)
        return self.total_tilt
