package main

import (
	"fmt"
)

func main() {
	matrix := [][]byte{
		{'0', '0', '0'},
		{'0', '0', '0'},
		{'0', '0', '0'},
		{'0', '0', '0'},
	}
	fmt.Println(maximalSquare(matrix))
}

func maximalSquare(matrix [][]byte) int {
	width, length := len(matrix), len(matrix[0])
	var result int
	dp := make([][]int, 0)
	for i := 0; i < width; i ++ {
		tmp := make([]int, length)
		dp = append(dp, tmp)
	}
	for i := 0; i < max2(width, length); i ++ {
		if i < length{
			dp[0][i] = int(matrix[0][i] - '0')
		}
		if i < width {
			dp[i][0] = int(matrix[i][0] - '0')
		}
	}
	for i := 1; i < width; i ++ {
		for j := 1; j < length; j ++ {
			if matrix[i][j] == '1' {
				dp[i][j] = min1(dp[i - 1][j], dp[i - 1][j - 1], dp[i][j - 1]) + 1
			}
		}
	}
	for i := 0; i < width; i ++ {
		for j := 0; j < length; j ++ {
			result = max2(result, dp[i][j])
		}
	}
	return result
}

func min1(a ...int) int {
	if len(a)==0 {
		return 0
	}
	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}


