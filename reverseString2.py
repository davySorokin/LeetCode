class Solution:
    def reverseStr(self, s: str, k: int) -> str:
        s = list(s) 
        for i in range(0, len(s), 2 * k):
            s[i:i + k] = reversed(s[i:i + k])  # Reverse first k characters in every 2k group
        return "".join(s)
