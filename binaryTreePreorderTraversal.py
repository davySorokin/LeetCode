# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def preorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        # Iterative solution
        if not root:
            return []
        
        stack = [root]  # Use a stack to simulate recursion
        result = []
        
        while stack:
            node = stack.pop()  # Pop the top node from the stack
            if node:
                result.append(node.val)  # Process the current node
                # Push right child first, so the left child is processed first
                if node.right:
                    stack.append(node.right)
                if node.left:
                    stack.append(node.left)
        
        return result
