class Solution:
    def canJump(self, nums: List[int]) -> bool:
        max_reachable = 0
        for i, jump in enumerate(nums):
            # If the current position is greater than the maximum reachable,
            # then we cannot proceed further, so return False
            if i > max_reachable:
                return False
            # Update the maximum reachable index
            max_reachable = max(max_reachable, i + jump)
            # If we've already reached or passed the last index, return True
            if max_reachable >= len(nums) - 1:
                return True
        return max_reachable >= len(nums) - 1
