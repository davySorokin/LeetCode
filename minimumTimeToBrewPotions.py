class Solution:
    def minTime(self, skill: list[int], mana: list[int]) -> int:
        n = len(skill)
        f = [0] * n

        for x in mana:
            now = f[0]
            for i in range(1, n):
                now = max(now + skill[i - 1] * x, f[i])
            f[-1] = now + skill[-1] * x
            # update f backwards
            for i in range(n - 2, -1, -1):
                f[i] = f[i + 1] - skill[i + 1] * x

        return f[-1]
