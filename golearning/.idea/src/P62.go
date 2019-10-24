package main

import "fmt"

/**
不同路径
 */
func main() {
	m, n := 7, 3
	fmt.Println(uniquePaths(m,n))
}

func uniquePaths(m int, n int) int {
	rec := make([][]int,m)
	for i := 0; i < m; i ++ {
		rec[i] = make([]int,n)
	}
	for i := 0; i < m; i ++ {
		rec[i][0] = 1
	}
	for i := 0; i < n; i ++ {
		rec[0][i] = 1
	}
	for i := 1; i < m; i ++ {
		for j := 1; j < n; j ++ {
			rec[i][j] = rec[i - 1][j] + rec[i][j - 1]
		}
	}
	return rec[m - 1][n - 1]
}