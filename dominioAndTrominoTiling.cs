public class Solution {
    public int NumTilings(int n) {
        int MOD = 1_000_000_007;
        if (n == 1) return 1;

        int[] dp = new int[n + 1];
        int[] dp2 = new int[n + 1];
        dp[0] = 1;
        dp[1] = 1;
        dp2[0] = 0;

        for (int i = 2; i <= n; i++) {
            dp[i] = (int)(((long)dp[i - 1] + dp[i - 2] + 2L * dp2[i - 1]) % MOD);
            dp2[i] = (int)(((long)dp2[i - 1] + dp[i - 2]) % MOD);
        }

        return dp[n];
    }
}

/*public class Solution {
    public int NumTilings(int n) {
        int MOD = 1_000_000_007;
        if (n == 1) return 1;
        if (n == 2) return 2;
        if (n == 3) return 5;

        int[] dp = new int[n + 1];
        dp[0] = 1;
        dp[1] = 1;
        dp[2] = 2;
        dp[3] = 5;

        for (int i = 4; i <= n; i++) {
            dp[i] = ((dp[i - 1] + dp[i - 2]) % MOD + (2L * dp[i - 3]) % MOD) % MOD;
        }

        return dp[n];
    }
}

public class Solution {
    public int NumTilings(int n) {
        int MOD = 1_000_000_007;
        if (n == 1) return 1;
        if (n == 2) return 2;
        if (n == 3) return 5;

        int[] dp = new int[n + 1];
        dp[0] = 1;
        dp[1] = 1;
        dp[2] = 2;
        dp[3] = 5;

        for (int i = 4; i <= n; i++) {
            dp[i] = (int)((dp[i - 1] + dp[i - 2] + 2L * dp[i - 3]) % MOD);
        }

        return dp[n];
    }
}*/
