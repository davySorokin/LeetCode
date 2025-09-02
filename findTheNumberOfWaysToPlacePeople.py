from typing import List

class Solution:
    def numberOfPairs(self, points: List[List[int]]) -> int:
        n = len(points)
        ans = 0
        for i in range(n):
            ax, ay = points[i]
            for j in range(n):
                if i == j:
                    continue
                bx, by = points[j]
                # A must be upper-left of B (lines allowed)
                if ax <= bx and ay >= by and (ax != bx or ay != by):
                    # Check if rectangle (including border) is empty
                    blocked = False
                    minx, maxx = ax, bx
                    miny, maxy = by, ay
                    for k in range(n):
                        if k == i or k == j:
                            continue
                        xk, yk = points[k]
                        if minx <= xk <= maxx and miny <= yk <= maxy:
                            blocked = True
                            break
                    if not blocked:
                        ans += 1
        return ans
