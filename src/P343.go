package main

import "fmt"

func main() {
	fmt.Println(integerBreak(10))
}

func integerBreak(n int) int {
	nums := make([]int, n + 1)
	nums[1] = 1
	nums[2] = 1
	for i := 3; i <= n; i ++ {
		res := i - 1
		for j := 1; j <= i / 2; j ++ {
			tmp := max9(nums[j], j) * max9(nums[i - j], i - j)
			if tmp > res {
				res = tmp
			}
		}
		nums[i] = res
	}
	return nums[n]
}

func max9(a, b int) int {
	if a > b {
		return a
	}
	return b
}