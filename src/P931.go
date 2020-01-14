package main

import "fmt"

func main() {
	A := [][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	fmt.Println(minFallingPathSum(A))
}

func minFallingPathSum(A [][]int) int {
	if len(A) == 0 {
		return 0
	}
	lenrow, lencol := len(A), len(A[0])
	dp := make([]int, lencol)
	for i := 0; i < lencol; i ++ {
		dp[i] = A[0][i]
	}
	fmt.Println(dp)
	for i := 1; i < lenrow; i ++ {
		tmpslice := make([]int, lencol)
		for j := 0; j < lencol; j ++ {
			tmp := 0
			if j == 0 {
				tmp = min8(dp[j], dp[j + 1])
			}else if j == lencol - 1 {
				tmp = min8(dp[j], dp[j - 1])
			}else {
				tmp = min8(dp[j], dp[j - 1], dp[j + 1])
			}
			tmpslice[j] = A[i][j] + tmp
		}
		dp = tmpslice
	}
	res := dp[0]
	for i := 1; i < lencol; i ++ {
		if res > dp[i] {
			res = dp[i]
		}
	}
	return res
}

func min8(a ...int) int {
	tmp := a[0]
	for i := 1; i < len(a); i ++ {
		if tmp > a[i] {
			tmp = a[i]
		}
	}
	return tmp
}