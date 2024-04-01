from typing import List

class Solution:
    def findMinArrowShots(self, points: List[List[int]]) -> int:
        if not points:
            return 0

        # Sort the points based on the end coordinate
        points.sort(key=lambda x: x[1])
        arrows = 1
        last_end = points[0][1]

        for start, end in points[1:]:
            # If the start of the current interval is greater than the
            # end of the last interval where an arrow was shot,
            # we need a new arrow.
            if start > last_end:
                arrows += 1
                last_end = end

        return arrows

# Example usage:
# sol = Solution()
# print(sol.findMinArrowShots([[10,16],[2,8],[1,6],[7,12]]))  # Output: 2
# print(sol.findMinArrowShots([[1,2],[3,4],[5,6],[7,8]]))    # Output: 4
# print(sol.findMinArrowShots([[1,2],[2,3],[3,4],[4,5]]))    # Output: 2
