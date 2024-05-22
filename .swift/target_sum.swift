class Solution {
    func findTargetSumWays(_ nums: [Int], _ target: Int) -> Int {
        let sum = nums.reduce(0, +)
        if sum < target || (sum + target) % 2 != 0 {
            return 0
        }

        let s = (sum + target) / 2
        var dp = [Int](repeating: 0, count: s + 1)
        dp[0] = 1

        for num in nums {
            for j in stride(from: s, through: num, by: -1) {
                dp[j] += dp[j - num]
            }
        }

        return dp[s]
    }
}
