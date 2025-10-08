from typing import List
import bisect

class Solution:
    def successfulPairs(self, spells: List[int], potions: List[int], success: int) -> List[int]:
        potions.sort()
        m = len(potions)
        ans = []
        for s in spells:
            # minimal potion needed: ceil(success / s)
            need = (success + s - 1) // s
            idx = bisect.bisect_left(potions, need)   # first potion >= need
            ans.append(m - idx)
        return ans
