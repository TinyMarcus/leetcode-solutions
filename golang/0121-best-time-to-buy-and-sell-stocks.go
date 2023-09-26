func maxProfit(prices []int) int {
    minVal := prices[0]
    maxProfit := 0

    for _, val := range prices {
        if val < minVal {
            minVal = val
        }

        if val - minVal > maxProfit {
            maxProfit = val - minVal
        }
    }

    return maxProfit
}