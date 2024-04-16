from typing import List
from collections import deque

class Solution:
    def minMutation(self, start: str, end: str, bank: List[str]) -> int:
        if end not in bank:
            return -1
        
        bank_set = set(bank)
        queue = deque([(start, 0)])
        while queue:
            current, mutations = queue.popleft()
            if current == end:
                return mutations
            
            for i in range(len(current)):
                for c in "ACGT":
                    mutation = current[:i] + c + current[i+1:]
                    if mutation in bank_set:
                        bank_set.remove(mutation)
                        queue.append((mutation, mutations + 1))
        return -1
