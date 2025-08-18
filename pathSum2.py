# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

from typing import List, Optional

class Solution:
    def pathSum(self, root: Optional['TreeNode'], targetSum: int) -> List[List[int]]:
        res, path = [], []

        def dfs(node: Optional['TreeNode'], remain: int) -> None:
            if not node:
                return
            path.append(node.val)
            remain -= node.val

            if not node.left and not node.right and remain == 0:
                res.append(path[:])  # record a copy

            dfs(node.left, remain)
            dfs(node.right, remain)

            path.pop()  # backtrack

        dfs(root, targetSum)
        return res
