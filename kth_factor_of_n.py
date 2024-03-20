class Solution:
    def kthFactor(self, n: int, k: int) -> int:
        # First half of factors
        for i in range(1, int(n**0.5) + 1):
            if n % i == 0:
                k -= 1
                if k == 0:
                    return i
        
        # Second half of factors
        # If n is a perfect square, we need to skip the square root
        # because it has already been counted.
        if (int(n**0.5))**2 == n:
            start = int(n**0.5) - 1
        else:
            start = int(n**0.5)

        for i in range(start, 0, -1):
            if n % i == 0:
                k -= 1
                if k == 0:
                    return n // i
        
        return -1
