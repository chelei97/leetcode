package main

import "fmt"

/**
买卖股票的最佳时机
 */
func main() {
	prices := []int{7,1,5,3,6,4}
	fmt.Println(maxProfit1(prices))
}

func maxProfit1(prices []int) int {
	if len(prices) == 0{
		return 0
	}
	dp_0, dp_1 := 0, -prices[0]
	for i := 0; i < len(prices); i ++ {
		dp_0 = Max1(dp_0, dp_1 + prices[i])
		dp_1 = Max1(dp_1, -prices[i])
	}
	return dp_0
}

func Max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}