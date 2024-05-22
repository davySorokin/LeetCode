class Solution {
    func findTargetSumWays(_ nums: [Int], _ target: Int) -> Int {
        var memo = [String: Int]()
        
        func backtrack(_ index: Int, _ currentSum: Int) -> Int {
            let key = "\(index)_\(currentSum)"
            if let val = memo[key] {
                return val
            }
            if index == nums.count {
                return currentSum == target ? 1 : 0
            }
            
            // Calculate the number of ways by adding and subtracting the current number
            let add = backtrack(index + 1, currentSum + nums[index])
            let subtract = backtrack(index + 1, currentSum - nums[index])
            
            // Memoize the result
            memo[key] = add + subtract
            return memo[key]!
        }
        
        return backtrack(0, 0)
    }
}
