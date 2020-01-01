package main

import "fmt"

func main() {
	matrix := [][]int{
		{1,1,1},
		{1,0,1},
		{1,1,1},
	}
	setZeroes(matrix)
	fmt.Println(matrix)
}

func setZeroes(matrix [][]int)  {
	isRowZero, isColZero := false, false
	rowNums := len(matrix)
	if rowNums == 0 {
		return
	}
	colNums := len(matrix[0])
	for i := 0; i < rowNums; i ++ {
		if matrix[i][0] == 0 {
			isRowZero = true
			break
		}
	}
	for i := 0; i < colNums; i ++ {
		if matrix[0][i] == 0 {
			isColZero = true
			break
		}
	}
	for i := 1; i < rowNums; i ++ {
		for j := 1; j < colNums; j ++ {
			if matrix[i][j] == 0 {
				matrix[i][0], matrix[0][j] = 0, 0
			}
		}
	}
	for i := 1; i < rowNums; i ++ {
		if matrix[i][0] == 0 {
			for j := 1; j < colNums; j ++ {
				matrix[i][j] = 0
			}
		}
	}
	for i := 1; i < colNums; i ++ {
		if matrix[0][i] == 0 {
			for j := 1; j < rowNums; j ++ {
				matrix[j][i] = 0
			}
		}
	}
	if isRowZero {
		for i := 0; i < rowNums; i ++ {
			matrix[i][0] = 0
		}
	}
	if isColZero {
		for i := 0; i < colNums; i ++ {
			matrix[0][i] = 0
		}
	}
}