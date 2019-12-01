package main

import (
	"fmt"
)

/**
螺旋矩阵
 */

func main(){
	matrix := [][]int{{1,2,3,4},{5,6,7,8},{9,10,11,12}}
	fmt.Println(spiralOrder(matrix))
}
func spiralOrder(matrix [][]int) []int {
	row, column, index := len(matrix), len(matrix[0]), 0
	result := make([]int,row * column)
	for i := 0; i < row / 2 + 1; i ++ {
		for j := i; j < column - 1 - i; j ++ {
			result[index] = matrix[i][j]
			index ++
		}
		for j := i; j < row - 1 - i; j ++ {
			result[index] = matrix[j][column - 1 - i]
			index ++
		}
		for j := i; j < column - 1 - i; j ++ {
			result[index] = matrix[row - 1 - i][column - 1 - j]
			index ++
		}
		for j := i; j < row - 1 - i; j ++ {
			result[index] = matrix[row - 1 - j][i]
			index ++
		}
	}
	return result
}
