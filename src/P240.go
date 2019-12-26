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
	fmt.Println(searchMatrix(matrix, 15))
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0{
		return false
	}
	width, length := len(matrix) - 1, len(matrix[0]) - 1
	w, l := width, 0
	for l <= length && w >= 0 {
		if matrix[w][l] < target {
			l ++
		}else if matrix[w][l] > target {
			w --
		}else {
			return true
		}
	}
	return false
}
