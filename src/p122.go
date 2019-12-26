package main

import "fmt"

/**
买卖股票的最佳时机
 */
func main() {

	prices := []int{1,2,3,4,5}
	fmt.Println(maxProfit(prices))

}

func maxProfit(prices []int) int {
	if len(prices) == 0{
		return 0
	}
	dp_0, dp_1 := 0, -prices[0]
	for i := 0; i < len(prices); i ++ {
		tmp := dp_0
		dp_0 = Max2(dp_0, dp_1 + prices[i])
		dp_1 = Max2(dp_1, tmp - prices[i])
	}
	return dp_0
}

func Max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}