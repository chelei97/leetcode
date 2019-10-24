package main

import "fmt"

/**
买卖股票的最佳时机
 */
func main() {
	prices := []int{7,1,5,3,6,4}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	if len(prices) == 0{
		return 0
	}
	min, max, result := prices[0], prices[0], 0
	for i := 1; i < len(prices); i ++ {
		if prices[i] > max {
			max = prices[i]
			if (max - min) > result{
				result = max - min
			}
		}else if prices[i] < min {
			min = prices[i]
			max = prices[i]
		}
	}
	return result
}
