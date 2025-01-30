from typing import List

class Solution:
    def findContentChildren(self, g: List[int], s: List[int]) -> int:
        # Sort both greed factors and cookie sizes in ascending order
        g.sort()
        s.sort()
        
        child_i = 0  # Pointer for children
        cookie_j = 0  # Pointer for cookies
        
        while child_i < len(g) and cookie_j < len(s):
            if s[cookie_j] >= g[child_i]:  
                child_i += 1  
            
            cookie_j += 1  
        
        return child_i
