package main

import "fmt"

func main() {
	price := []int{1,2,3,0,2}
	fmt.Println(maxProfit3(price))
}

func maxProfit3(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	dp_0, dp_1, dp_pre := 0, -prices[0], 0
	for i := 1; i < len(prices); i ++ {
		tmp := dp_0
		dp_0 = max3(dp_0, dp_1 + prices[i])
		dp_1 = max3(dp_1, dp_pre - prices[i])
		dp_pre = tmp
	}
	return dp_0
}

func max3(a, b int) int {
	if a > b {
		return a
	}
	return b
}
