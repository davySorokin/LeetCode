func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }
    
    // First buy/sell and second buy/sell profits
    firstBuy, secondBuy := -prices[0], -prices[0]
    firstSell, secondSell := 0, 0

    for i := 1; i < len(prices); i++ {
        // First transaction: either do nothing or buy at a lower price
        firstBuy = max(firstBuy, -prices[i])
        // First sell: either do nothing or sell at current price
        firstSell = max(firstSell, firstBuy + prices[i])
        
        // Second transaction: either do nothing or use profit from first sell to buy again
        secondBuy = max(secondBuy, firstSell - prices[i])
        // Second sell: either do nothing or sell after second buy
        secondSell = max(secondSell, secondBuy + prices[i])
    }

    return secondSell
}

// Helper function to get the maximum of two integers
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
