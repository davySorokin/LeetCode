# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def isValidBST(self, root: TreeNode) -> bool:
        def validate(node, low=float('-inf'), high=float('inf')):
            # Empty trees are valid BSTs
            if not node:
                return True
            
            # The current node's value must be between low and high
            if node.val <= low or node.val >= high:
                return False
            
            # The left subtree must be < node.val and the right subtree > node.val
            return (validate(node.left, low, node.val) and
                    validate(node.right, node.val, high))
        
        return validate(root)
