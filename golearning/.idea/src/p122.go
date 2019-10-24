package main

import "fmt"

/**
买卖股票的最佳时机
 */
func main() {

	prices := []int{7,6,4,3,1}
	fmt.Println(maxProfit(prices))

}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	result := 0
	price := prices[0]
	for i := 1; i < len(prices); i ++ {
		if prices[i] > price {
			for i + 1 < len(prices) && prices[i + 1] > prices[i]{
				i ++
			}
			result = result + prices[i] - price
			if i + 1 < len(prices) {
				price = prices[i + 1]
			}
		}
		if prices[i] < price {
			price = prices[i]
		}
	}
	return result
}