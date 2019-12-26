package main

import "fmt"

func main() {
	price := []int{1, 3, 2, 8, 4, 9}
	fmt.Println(maxProfit4(price, 2))
}

func maxProfit4(prices []int, fee int) int {
	if len(prices) < 2 {
		return 0
	}
	dp_0, dp_1 := 0, -prices[0]
	for i := 1; i < len(prices); i ++ {
		tmp := dp_0
		dp_0 = max4(dp_0, dp_1 + prices[i] - fee)
		dp_1 = max4(dp_1, tmp - prices[i])
	}
	return dp_0
}

func max4(a, b int) int {
	if a > b {
		return a
	}
	return b
}