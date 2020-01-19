package main

import "fmt"

func main() {
	fmt.Println(nthUglyNumber(10))
}

func nthUglyNumber(n int) int {
	dp, newdp := 1, 1
	for i := 1; i < n; i ++ {
		newdp = min11(dp * 2, dp * 3, dp * 5)
		fmt.Println(newdp)
		dp = newdp
	}
	return dp
}

func min11(a, b, c int) int {
	if a > b {
		a = b
	}
	if a > c {
		return c
	}
	return a
}
