class Solution {
    public int maxProfit(int[] prices) {
        int maxProfit = 0;
        // Iterate through the array of prices
        for (int i = 1; i < prices.length; i++) {
            // If the current price is higher than the previous one
            if (prices[i] > prices[i - 1]) {
                // Add the difference to the max profit
                maxProfit += prices[i] - prices[i - 1];
            }
        }
        // Return the total profit
        return maxProfit;
    }
}
