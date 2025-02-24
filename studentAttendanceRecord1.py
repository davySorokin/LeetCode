class Solution:
    def checkRecord(self, s: str) -> bool:
        # Check if 'A' appears fewer than 2 times
        if s.count('A') >= 2:
            return False
        
        # Check if 'L' appears in a sequence of 3 or more
        if "LLL" in s:
            return False
        
        return True
