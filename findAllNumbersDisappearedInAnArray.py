from typing import List

class Solution:
    def findDisappearedNumbers(self, nums: List[int]) -> List[int]:
        # Mark the numbers seen in the array by making the indices negative
        for i in range(len(nums)):
            index = abs(nums[i]) - 1  # Get the index corresponding to the value
            if nums[index] > 0:
                nums[index] = -nums[index]  # Mark it as visited (negative)
        
        # Collect the missing numbers
        result = []
        for i in range(len(nums)):
            if nums[i] > 0:  # If the value is positive, the index+1 is missing
                result.append(i + 1)
        
        return result
