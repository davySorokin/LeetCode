class Solution:
    def smallestEquivalentString(self, s1: str, s2: str, baseStr: str) -> str:
        parent = {}

        def find(x):
            if parent[x] != x:
                parent[x] = find(parent[x])
            return parent[x]

        def union(x, y):
            px, py = find(x), find(y)
            if px == py:
                return
            if px < py:
                parent[py] = px
            else:
                parent[px] = py

        # Initialise each character's parent to itself
        for c in 'abcdefghijklmnopqrstuvwxyz':
            parent[c] = c

        # Union characters in s1 and s2
        for a, b in zip(s1, s2):
            union(a, b)

        # Build result by finding smallest equivalent character
        return ''.join(find(c) for c in baseStr)
