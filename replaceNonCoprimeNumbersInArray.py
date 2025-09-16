from math import gcd
from typing import List

class Solution:
    def replaceNonCoprimes(self, nums: List[int]) -> List[int]:
        st = []
        for x in nums:
            cur = x
            # keep merging while the top and cur are non-coprime
            while st:
                g = gcd(st[-1], cur)
                if g == 1:
                    break
                cur = st.pop() // g * cur  # lcm(top, cur) = top/g * cur
            st.append(cur)
        return st
