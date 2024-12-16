from typing import Optional, List

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def binaryTreePaths(self, root: Optional[TreeNode]) -> List[str]:
        result = []
        
        def dfs(node, path):
            if not node:
                return
            
            # Append the current node value to the path
            path += str(node.val)
            
            # If it is a leaf node, add the path to result
            if not node.left and not node.right:
                result.append(path)
                return
            
            # Recur to left and right children
            path += "->"
            dfs(node.left, path)
            dfs(node.right, path)
        
        dfs(root, "")
        return result
