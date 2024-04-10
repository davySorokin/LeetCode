class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def maxPathSum(self, root: TreeNode) -> int:
        self.max_sum = float('-inf')

        def max_gain(node):
            if not node:
                return 0

            # Get max sum of paths in left and right subtrees
            left_gain = max(max_gain(node.left), 0)
            right_gain = max(max_gain(node.right), 0)

            # Current path sum including the current node
            price_newpath = node.val + left_gain + right_gain

            # Update max_sum if it's better to start a new path
            self.max_sum = max(self.max_sum, price_newpath)

            # For recursion, return the max gain if continue the same path
            return node.val + max(left_gain, right_gain)

        max_gain(root)
        return self.max_sum
