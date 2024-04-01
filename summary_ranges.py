from typing import List

class Solution:
    def summaryRanges(self, nums: List[int]) -> List[str]:
        if not nums:
            return []

        ranges = []
        start = nums[0]
        end = nums[0]

        for n in nums[1:]:
            if n == end + 1:
                # Continue the sequence
                end = n
            else:
                # End the current sequence and start a new one
                if start == end:
                    ranges.append(str(start))
                else:
                    ranges.append(f"{start}->{end}")
                start = end = n

        # Add the last range
        if start == end:
            ranges.append(str(start))
        else:
            ranges.append(f"{start}->{end}")

        return ranges

# Example usage:
# sol = Solution()
# print(sol.summaryRanges([0, 1, 2, 4, 5, 7]))  # Output: ["0->2", "4->5", "7"]
# print(sol.summaryRanges([0, 2, 3, 4, 6, 8, 9]))  # Output: ["0", "2->4", "6", "8->9"]
