package main

import (
	"fmt"
)

func main() {
	n := 13
	fmt.Println(numSquares(n))
}

func numSquares(n int) int {
	nums := make([]int, sqrt(n))
	for i := 0; i < len(nums); i ++ {
		nums[i] = (i + 1) * (i + 1)
	}
	dp := make([]int, n + 1)
	for i := 1; i < len(dp); i ++ {
		dp[i] = n + 1
	}
	for i := 1; i <= n; i ++ {
		for _, v := range nums {
			if i >= v {
				dp[i] = min(dp[i], dp[i - v] + 1)
			}
		}
	}
	return dp[n]
}

func sqrt(n int) int {
	if n < 4 {
		return 1
	}
	result := 2
	for result * result <= n {
		result ++
	}
	return result - 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
