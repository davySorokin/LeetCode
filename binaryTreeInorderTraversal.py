# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def inorderTraversal(self, root: TreeNode) -> list[int]:
        result = []
        stack = []
        current = root

        while current or stack:
            while current:  # Traverse to the leftmost node
                stack.append(current)
                current = current.left
            current = stack.pop()  # Start backtracking
            result.append(current.val)  # Visit the root
            current = current.right  # Traverse the right subtree

        return result
