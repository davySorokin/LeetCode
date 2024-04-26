from typing import List

class Solution:
    def maxSubarraySumCircular(self, nums: List[int]) -> int:
        def kadane(gen):
            # Traditional Kadane's algorithm
            max_sum = cur_sum = float('-inf')
            for x in gen:
                cur_sum = x + max(cur_sum, 0)
                max_sum = max(max_sum, cur_sum)
            return max_sum

        total_sum = sum(nums)
        max_sum_subarray = kadane(nums)
        min_sum_subarray = kadane(-x for x in nums)

        # The max wrapped subarray = total sum - min sum subarray
        # Don't consider the case where no element is chosen, hence the max() with total_sum
        max_sum_wrapped = total_sum + min_sum_subarray
        
        # If all numbers are negative, max_sum_subarray is the maximum single element
        # In this case, max_sum_wrapped will be 0, so we shouldn't use it
        if max_sum_wrapped > 0:
            return max(max_sum_subarray, max_sum_wrapped)
        else:
            return max_sum_subarray
