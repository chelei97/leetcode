package main

import (
	"fmt"
	"math"
)

func main() {
	matrix := [][]int{
		{0},
	}
	fmt.Println(updateMatrix(matrix))
}

func updateMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return nil
	}
	lenRow, lenCol := len(matrix), len(matrix[0])
	var res [][]int
	for i := 0; i < lenRow; i ++ {
		tmp := make([]int, lenCol)
		for j := 0; j < lenCol; j ++ {
			if matrix[i][j] != 0 {
				tmp[j] = math.MaxInt32 - 10000
			}
		}
		res = append(res, tmp)
	}
	for i := 1; i < lenCol; i ++ {
		if matrix[0][i] == 0 {
			res[0][i] = 0
		}else {
			res[0][i] = res[0][i - 1] + 1
		}
	}
	for i := 1; i < lenRow; i ++ {
		for j := 0; j < lenCol; j ++ {
			if matrix[i][j] != 0 {
				if j == 0 {
					res[i][j] = res[i - 1][j] + 1
				}else {
					res[i][j] = min12(res[i][j - 1], res[i - 1][j]) + 1
				}
			}else {
				res[i][j] = 0
			}
		}
	}
	for i := lenCol - 2; i >= 0; i -- {
		if matrix[lenRow - 1][i] != 0 {
			res[lenRow - 1][i] = min12(res[lenRow - 1][i + 1] + 1, res[lenRow - 1][i])
		}
	}
	for i := lenRow - 2; i >= 0; i -- {
		for j := lenCol - 1; j >= 0; j -- {
			if matrix[i][j] == 0 {
				res[i][j] = 0
			}else {
				if j == lenCol - 1 {
					res[i][j] = min12(res[i + 1][j] + 1, res[i][j])
				}else {
					res[i][j] = min12(min12(res[i + 1][j], res[i][j + 1]) + 1, res[i][j])
				}
			}
		}
	}
	return res
}

func min12(a, b int) int {
	if a < b {
		return a
	}
	return b
}
