class Solution:
    def sortedArrayToBST(self, nums: List[int]) -> Optional[TreeNode]:
        def helper(left, right):
            if left > right:
                return None

            # Find the middle index
            mid = (left + right) // 2

            # Create the root node with the middle element
            root = TreeNode(nums[mid])

            # Recursively construct the left and right subtrees
            root.left = helper(left, mid - 1)
            root.right = helper(mid + 1, right)

            return root

        return helper(0, len(nums) - 1)
