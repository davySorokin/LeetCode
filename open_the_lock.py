from typing import List
from collections import deque

class Solution:
    def openLock(self, deadends: List[str], target: str) -> int:
        def neighbors(node):
            for i in range(4):
                x = int(node[i])
                for d in (-1, 1):
                    y = (x + d) % 10
                    yield node[:i] + str(y) + node[i+1:]

        dead = set(deadends)
        queue = deque([('0000', 0)])
        visited = set('0000')

        while queue:
            node, depth = queue.popleft()
            if node == target: return depth
            if node in dead: continue
            for neighbor in neighbors(node):
                if neighbor not in visited:
                    visited.add(neighbor)
                    queue.append((neighbor, depth + 1))

        return -1
