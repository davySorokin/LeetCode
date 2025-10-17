from bisect import bisect_right, insort
from typing import List

class Solution:
    def avoidFlood(self, rains: List[int]) -> List[int]:
        n = len(rains)
        full = {}        # lake -> last rain day
        dry_days = []    # sorted list of indices where rains[i] == 0
        ans = [-1] * n

        for i, lake in enumerate(rains):
            if lake == 0:
                insort(dry_days, i)
                ans[i] = 1  # placeholder; will adjust later
            else:
                if lake in full:
                    prev = full[lake]
                    j = bisect_right(dry_days, prev)
                    if j == len(dry_days):
                        return []  # no dry day available
                    dry_day = dry_days[j]
                    ans[dry_day] = lake
                    dry_days.pop(j)
                full[lake] = i
                ans[i] = -1

        return ans
