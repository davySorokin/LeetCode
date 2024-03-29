class Solution:
    def buildArray(self, target: List[int], n: int) -> List[str]:
        operations = []
        j = 0  # Pointer for target array
        for i in range(1, n + 1):
            if j >= len(target):
                break  # Target array has been completely built
            
            operations.append("Push")  # Push every number
            
            if target[j] == i:
                j += 1  # Move to the next target if pushed number is in target
            else:
                operations.append("Pop")  # Pop if the number is not in target
        
        return operations

# Example use case
# sol = Solution()
# print(sol.buildArray([1,3], 3))  # Output: ["Push", "Push", "Pop", "Push"]
