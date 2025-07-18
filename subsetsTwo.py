class Solution:
    def subsetsWithDup(self, nums: List[int]) -> List[List[int]]:
        # Sort the input to handle duplicates
        nums.sort()
        res = []

        def backtrack(start, path):
            # Add the current subset to the result
            res.append(path)
            
            # Start exploring subsets starting from the next index
            for i in range(start, len(nums)):
                # Skip duplicates
                if i > start and nums[i] == nums[i - 1]:
                    continue
                # Include current element and move to next index
                backtrack(i + 1, path + [nums[i]])

        # Start backtracking from index 0
        backtrack(0, [])
        return res
