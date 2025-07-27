from typing import List

class Solution:
    def countHillValley(self, nums: List[int]) -> int:
        # Remove consecutive duplicates
        new_nums = [nums[0]]
        for i in range(1, len(nums)):
            if nums[i] != nums[i - 1]:
                new_nums.append(nums[i])
        
        # Count hills and valleys
        count = 0
        for i in range(1, len(new_nums) - 1):
            if (new_nums[i] > new_nums[i - 1] and new_nums[i] > new_nums[i + 1]) or \
               (new_nums[i] < new_nums[i - 1] and new_nums[i] < new_nums[i + 1]):
                count += 1
        
        return count
