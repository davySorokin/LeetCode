from collections import defaultdict

class Solution:
    def maxSubarrayLength(self, nums: List[int], k: int) -> int:
        count = defaultdict(int)
        left = 0
        max_length = 0

        for right, value in enumerate(nums):
            count[value] += 1
            
            # If the count of any number exceeds k, shrink the window from the left
            while count[value] > k:
                count[nums[left]] -= 1
                left += 1
            
            # Update the max length of the window
            max_length = max(max_length, right - left + 1)
        
        return max_length
