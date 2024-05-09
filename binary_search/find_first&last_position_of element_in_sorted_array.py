from typing import List

class Solution:
    def searchRange(self, nums: List[int], target: int) -> List[int]:
        def findFirst():
            low, high = 0, len(nums) - 1
            while low <= high:
                mid = (low + high) // 2
                if nums[mid] < target:
                    low = mid + 1
                elif nums[mid] > target:
                    high = mid - 1
                else:
                    if mid == 0 or nums[mid - 1] != target:
                        return mid
                    high = mid - 1
            return -1
        
        def findLast():
            low, high = 0, len(nums) - 1
            while low <= high:
                mid = (low + high) // 2
                if nums[mid] < target:
                    low = mid + 1
                elif nums[mid] > target:
                    high = mid - 1
                else:
                    if mid == len(nums) - 1 or nums[mid + 1] != target:
                        return mid
                    low = mid + 1
            return -1
        
        first = findFirst()
        last = findLast() if first != -1 else -1
        return [first, last]
