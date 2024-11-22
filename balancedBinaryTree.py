class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def isBalanced(self, root: TreeNode) -> bool:
        def check_balance_and_height(node):
            if not node:
                return True, 0
            
            left_balanced, left_height = check_balance_and_height(node.left)
            right_balanced, right_height = check_balance_and_height(node.right)
            
            # A tree is balanced if:
            # 1. The left subtree is balanced
            # 2. The right subtree is balanced
            # 3. The height difference between left and right is <= 1
            is_balanced = left_balanced and right_balanced and abs(left_height - right_height) <= 1
            height = max(left_height, right_height) + 1
            
            return is_balanced, height
        
        return check_balance_and_height(root)[0]
