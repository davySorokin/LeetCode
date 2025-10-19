from collections import deque

class Solution:
    def findLexSmallestString(self, s: str, a: int, b: int) -> str:
        seen = set()
        queue = deque([s])
        res = s

        while queue:
            cur = queue.popleft()
            if cur in seen:
                continue
            seen.add(cur)
            res = min(res, cur)

            # Operation 1: add 'a' to all digits at odd indices
            arr = list(cur)
            for i in range(1, len(arr), 2):
                arr[i] = str((int(arr[i]) + a) % 10)
            added = ''.join(arr)
            if added not in seen:
                queue.append(added)

            # Operation 2: rotate right by b positions
            rotated = cur[-b:] + cur[:-b]
            if rotated not in seen:
                queue.append(rotated)

        return res
