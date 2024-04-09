from typing import List

class Solution:
    def timeRequiredToBuy(self, tickets: List[int], k: int) -> int:
        time = 0
        # Loop until the person at index k has no tickets left to buy
        while tickets[k] > 0:
            for i in range(len(tickets)):
                # Sell ticket if the person at position i has tickets left
                if tickets[i] > 0:
                    tickets[i] -= 1
                    time += 1
                # If we reach the person at position k and they have no tickets left, break the loop
                if i == k and tickets[k] == 0:
                    break
        return time

# Example run:
# sol = Solution()
# print(sol.timeRequiredToBuy([2, 3, 2], k = 2)) # Output: 6
# print(sol.timeRequiredToBuy([5, 1, 1, 1], k = 0)) # Output: 8
