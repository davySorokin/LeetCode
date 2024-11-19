class Solution {
    public boolean[] canEat(int[] candiesCount, int[][] queries) {
        int n = candiesCount.length;
        int m = queries.length;

        // Create a prefix sum array to compute total candies up to a specific type
        long[] prefixSum = new long[n];
        prefixSum[0] = candiesCount[0];
        for (int i = 1; i < n; i++) {
            prefixSum[i] = prefixSum[i - 1] + candiesCount[i];
        }

        boolean[] answer = new boolean[m];

        for (int i = 0; i < m; i++) {
            int favoriteType = queries[i][0];
            int favoriteDay = queries[i][1];
            int dailyCap = queries[i][2];

            // Compute the range of days where candies of the favoriteType can be eaten
            long minCandiesNeeded = favoriteDay + 1; // Minimum candies needed to reach favoriteDay
            long maxCandiesPossible = (long) (favoriteDay + 1) * dailyCap; // Maximum candies one can eat by favoriteDay

            // Compute the range of candies available for favoriteType
            long candiesBefore = favoriteType > 0 ? prefixSum[favoriteType - 1] : 0; // Candies before the favorite type
            long candiesAvailable = prefixSum[favoriteType]; // Total candies up to the favorite type

            // Check if the ranges overlap
            answer[i] = !(minCandiesNeeded > candiesAvailable || maxCandiesPossible <= candiesBefore);
        }

        return answer;
    }
}
