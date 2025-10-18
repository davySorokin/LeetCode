class Solution:
    def maxDistinctElements(self, nums: List[int], k: int) -> int:
        # Sort the array to handle elements in increasing order
        nums.sort()
        
        # A set to track the distinct elements
        used = set()
        
        for num in nums:
            # Try the number itself first
            if num not in used:
                used.add(num)
            # Try num + 1, num - 1, ... within the range [-k, k]
            else:
                for delta in range(-k, k + 1):
                    new_num = num + delta
                    if new_num not in used:
                        used.add(new_num)
                        break
        
        # The size of the set is the maximum number of distinct elements
        return len(used)
