class Solution:
    def maxBottlesDrunk(self, numBottles: int, numExchange: int) -> int:
        bottles_drunk = 0
        full_bottles = numBottles
        empty_bottles = 0
        
        while full_bottles > 0:
            # Drink all full bottles
            bottles_drunk += full_bottles
            empty_bottles += full_bottles
            
            # Convert empty bottles to full bottles
            full_bottles = empty_bottles // numExchange
            empty_bottles = empty_bottles % numExchange
            
            # Increase numExchange by 1 after each exchange
            numExchange += 1
            
        return bottles_drunk
