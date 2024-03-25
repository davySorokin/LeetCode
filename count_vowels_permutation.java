class Solution {
    private static final int MOD = 1000000007;

    public int countVowelPermutation(int n) {
        long[] dp = new long[5];
        Arrays.fill(dp, 1); // Each vowel can be the start
        
        for (int i = 2; i <= n; i++) {
            long[] dpNew = new long[5];
            dpNew[0] = (dp[1] + dp[2] + dp[4]) % MOD; // 'a' can be followed by 'e', 'i', or 'u'
            dpNew[1] = (dp[0] + dp[2]) % MOD; // 'e' can be followed by 'a' or 'i'
            dpNew[2] = (dp[1] + dp[3]) % MOD; // 'i' can be followed by 'e' or 'o'
            dpNew[3] = dp[2]; // 'o' can be followed by 'i'
            dpNew[4] = (dp[2] + dp[3]) % MOD; // 'u' can be followed by 'i' or 'o'
            dp = dpNew;
        }
        
        long result = 0;
        for (long count : dp) {
            result = (result + count) % MOD;
        }
        return (int) result;
    }
}
