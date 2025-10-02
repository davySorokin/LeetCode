class Solution:
    def maxBottlesDrunk(self, numBottles: int, numExchange: int) -> int:
        drank, full, empty = 0, numBottles, 0
        while full > 0:
            drank += full
            empty += full
            full = 0
            if empty >= numExchange:
                empty -= numExchange
                full += 1
                numExchange += 1
        return drank
