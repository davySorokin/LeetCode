# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def findMode(self, root: Optional[TreeNode]) -> List[int]:
        if not root:
            return []
        
        self.current_val = None
        self.max_count = 0
        self.current_count = 0
        self.modes = []
        
        def in_order(node):
            if not node:
                return
            in_order(node.left)
            
            # Count the occurrences of the current value
            if node.val != self.current_val:
                self.current_val = node.val
                self.current_count = 1
            else:
                self.current_count += 1
            
            # Update modes if current count equals max count
            if self.current_count == self.max_count:
                self.modes.append(self.current_val)
            # Update max count and reset modes if current count exceeds max count
            elif self.current_count > self.max_count:
                self.max_count = self.current_count
                self.modes = [self.current_val]
            
            in_order(node.right)
        
        # First pass to find the max count
        in_order(root)
        
        # Reset for the second pass
        self.current_val = None
        self.current_count = 0
        modes = self.modes
        self.modes = []
        
        # Second pass to collect all the modes
        in_order(root)
        
        return self.modes
