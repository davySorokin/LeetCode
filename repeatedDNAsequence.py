from typing import List

class Solution:
    def findRepeatedDnaSequences(self, s: str) -> List[str]:
        n = len(s)
        if n < 10:
            return []
        enc = {'A': 0, 'C': 1, 'G': 2, 'T': 3}
        mask = (1 << 20) - 1  # keep last 20 bits (10 chars * 2 bits)
        x = 0

        seen = set()
        added = set()
        ans: List[str] = []

        for i, ch in enumerate(s):
            x = ((x << 2) & mask) | enc[ch]
            if i >= 9:
                if x in seen:
                    if x not in added:
                        ans.append(s[i - 9 : i + 1])
                        added.add(x)
                else:
                    seen.add(x)
        return ans
