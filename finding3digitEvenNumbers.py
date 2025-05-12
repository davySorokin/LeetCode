from typing import List
from itertools import permutations

class Solution:
    def findEvenNumbers(self, digits: List[int]) -> List[int]:
        result = set()
        
        for perm in permutations(digits, 3):
            if perm[0] == 0:  # No leading zero
                continue
            if perm[2] % 2 != 0:  # Must end with an even digit
                continue
            
            num = perm[0] * 100 + perm[1] * 10 + perm[2]
            result.add(num)
        
        return sorted(result)
