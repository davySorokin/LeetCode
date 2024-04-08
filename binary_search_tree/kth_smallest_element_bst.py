# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def kthSmallest(self, root: TreeNode, k: int) -> int:
        # In-order traversal: Left -> Node -> Right
        def in_order(r):
            return in_order(r.left) + [r.val] + in_order(r.right) if r else []

        # Perform in-order traversal and return the kth element
        return in_order(root)[k - 1]

# If you want to optimize for a frequently modified BST, you could augment the tree nodes to hold a count of nodes in the left subtree.

class TreeNodeAugmented:
    def __init__(self, val=0, left=None, right=None, left_count=0):
        self.val = val
        self.left = left
        self.right = right
        self.left_count = left_count

class SolutionAugmented:
    def kthSmallest(self, root: TreeNodeAugmented, k: int) -> int:
        def in_order_count(node, k):
            if not node:
                return None
            # Check the left subtree size
            if k <= node.left_count:
                return in_order_count(node.left, k)
            # Check if current node is the kth node
            elif k == node.left_count + 1:
                return node.val
            # If not, proceed to the right subtree
            else:
                return in_order_count(node.right, k - node.left_count - 1)
                
        return in_order_count(root, k)
