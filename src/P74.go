package main

import "fmt"

func main() {
	matrix := [][]int{
		{1,2,3,4,5},
		{6,7,8,9,10},
		{11,12,13,14,15},
		{16,17,18,19,20},
		{21,22,23,24,25},
	}
	fmt.Println(searchMatrix1(matrix, 24))
}

func searchMatrix1(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	width, length := len(matrix), len(matrix[0])
	i, j := 0, length - 1
	for i < width && i >= 0 && j >= 0 && j < length {
		if matrix[i][j] > target {
			j --
		}else if matrix[i][j] < target {
			i ++
		}else {
			return true
		}
	}
	return false
}
