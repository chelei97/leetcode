package main

import "fmt"

func main() {
	fmt.Println(minSteps(9))
}

func minSteps(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n + 1)
	for i := 2; i <= n; i ++ {
		tmp := findMax(i)
		dp[i] = dp[tmp] + i / tmp
	}
	return dp[n]
}

func findMax(n int) int {
	for i := n / 2; i >= 1; i -- {
		if n % i == 0 {
			return i
		}
	}
	return 1
}


