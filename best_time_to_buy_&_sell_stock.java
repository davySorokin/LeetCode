class Solution {
    public int maxProfit(int[] prices) {
        if (prices == null || prices.length <= 1) {
            return 0;
        }

        int minPrice = prices[0]; // Initialize minPrice to the first price
        int maxProfit = 0; // Initialize maxProfit to 0

        for (int i = 1; i < prices.length; i++) {
            if (prices[i] < minPrice) {
                minPrice = prices[i]; // Update minPrice if the current price is lower
            } else {
                maxProfit = Math.max(maxProfit, prices[i] - minPrice); // Update maxProfit if selling gives more profit
            }
        }

        return maxProfit;
    }
}
