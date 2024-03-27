class Solution:
    def minSubArrayLen(self, target: int, nums: List[int]) -> int:
        min_length = float('inf')
        current_sum = 0
        start = 0
        
        for end in range(len(nums)):
            current_sum += nums[end]
            
            # Once current_sum exceeds the target, move the start pointer to reduce the window size
            while current_sum >= target:
                min_length = min(min_length, end - start + 1)
                current_sum -= nums[start]
                start += 1
        
        return min_length if min_length != float('inf') else 0
