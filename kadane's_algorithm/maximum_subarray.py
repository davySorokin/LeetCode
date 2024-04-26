from typing import List

class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        def findMaxCrossingSum(nums, left, mid, right):
            # Include elements on left of mid.
            sum_left = float('-inf')
            sum_temp = 0
            for i in range(mid, left - 1, -1):
                sum_temp += nums[i]
                if sum_temp > sum_left:
                    sum_left = sum_temp

            # Include elements on right of mid
            sum_right = float('-inf')
            sum_temp = 0
            for i in range(mid + 1, right + 1):
                sum_temp += nums[i]
                if sum_temp > sum_right:
                    sum_right = sum_temp

            # Return sum of elements on left and right of mid
            return sum_left + sum_right

        def findMaximumSubarray(nums, left, right):
            # Base Case: Only one element
            if left == right:
                return nums[left]

            # Find middle point
            mid = (left + right) // 2

            return max(
                findMaximumSubarray(nums, left, mid),
                findMaximumSubarray(nums, mid + 1, right),
                findMaxCrossingSum(nums, left, mid, right)
            )

        # Call the recursive function to find the maximum subarray sum
        return findMaximumSubarray(nums, 0, len(nums) - 1)
