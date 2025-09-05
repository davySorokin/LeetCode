class Solution:
    def makeTheIntegerZero(self, num1: int, num2: int) -> int:
        # Try k from 1 to 60 (since i ∈ [0, 60] and that's enough for constraints)
        for k in range(1, 61):
            S = num1 - k * num2
            if S <= 0:
                # If num2 >= 0, S will only get smaller; we can break early.
                if num2 >= 0:
                    break
                # If num2 < 0, S might become positive later; continue.
                continue
            # Need popcount(S) <= k <= S
            if S.bit_count() <= k <= S:
                return k
        return -1
