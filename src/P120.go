package main

import "fmt"

func main() {
	triangle := [][]int{
		{2},
		{3,4},
		{6,5,7},
		{4,1,8,3},
	}
	fmt.Println(minimumTotal(triangle))
}

func minimumTotal(triangle [][]int) int {
	height := len(triangle)
	record := make([]int, height)
	record[0] = triangle[0][0]
	for i := 1; i < height; i ++ {
		tmp := make([]int, i + 1)
		tmp[0] = record[0] + triangle[i][0]
		tmp[i] = record[i - 1] + triangle[i][i]
		for j := 1; j < i; j ++ {
			tmp[j] = min3(record[j], record[j - 1]) + triangle[i][j]
		}
		for j := 0; j <= i; j ++ {
			record[j] = tmp[j]
		}
	}
	res := record[0]
	for i := 1; i < height; i ++ {
		if res > record[i] {
			res = record[i]
		}
	}
	return res
}

func min3(a, b int) int {
	if a < b {
		return a
	}
	return b
}