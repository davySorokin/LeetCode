from heapq import heappop, heappush
from typing import List

class Solution:
    def minTimeToReach(self, moveTime: List[List[int]]) -> int:
        n, m = len(moveTime), len(moveTime[0])
        directions = [(0, 1), (1, 0), (0, -1), (-1, 0)]
        heap = [(0, 0, 0, 0)]
        visited = [[[False] * 2 for _ in range(m)] for _ in range(n)]

        while heap:
            time, x, y, parity = heappop(heap)
            if visited[x][y][parity]:
                continue
            visited[x][y][parity] = True
            if x == n - 1 and y == m - 1:
                return time
            for dx, dy in directions:
                nx, ny = x + dx, y + dy
                if 0 <= nx < n and 0 <= ny < m:
                    wait_time = max(0, moveTime[nx][ny] - time)
                    step_cost = 1 if parity == 0 else 2
                    next_time = time + wait_time + step_cost
                    next_parity = 1 - parity
                    if not visited[nx][ny][next_parity]:
                        heappush(heap, (next_time, nx, ny, next_parity))
        return -1
