class Solution:
    def doesAliceWin(self, s: str) -> bool:
        return any(ch in "aeiou" for ch in s)
