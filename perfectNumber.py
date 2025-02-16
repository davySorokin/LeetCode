class Solution:
    def checkPerfectNumber(self, num: int) -> bool:
        # 1 is not a perfect number
        if num <= 1:
            return False
        
        # Start with sum_of_divisors = 1 because 1 is always a divisor (excluding the number itself)
        sum_of_divisors = 1
        i = 2
        
        # Only check up to sqrt(num)
        while i * i <= num:
            if num % i == 0:
                sum_of_divisors += i
                # Add the complement divisor if it's different
                if i * i != num:
                    sum_of_divisors += num // i
            i += 1
        
        return sum_of_divisors == num
