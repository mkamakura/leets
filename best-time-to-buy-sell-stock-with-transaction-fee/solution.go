func max(a, b int) int {
    if (a > b) {
        return a
    }
    return b
}

func maxProfit(prices []int, fee int) int {
    cash := 0
    hold := -prices[0]

    for i := 1; i < len(prices); i++ {
        cash = max(cash, hold + prices[i] - fee)
        hold = max(hold, cash - prices[i])
    }
    return cash
}
