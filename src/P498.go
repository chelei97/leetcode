package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{
		{1},
	}
	fmt.Println(findDiagonalOrder(matrix))
}

func findDiagonalOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	var res []int
	i, j := 0, 0
	row, col := len(matrix) - 1, len(matrix[0]) - 1
	addr, addc := 0, 0
	for !(i > row || j > col) {
		res = append(res, matrix[i][j])
		if i == 0 {
			if j == col {
				i ++
			}else {
				j ++
			}
			if i > row || j > col {
				break
			}
			res = append(res, matrix[i][j])
			addr = 1
			addc = -1
		}else if i == row {
			j ++
			if i > row || j > col {
				break
			}
			res = append(res, matrix[i][j])
			addr = -1
			addc = 1
		}else if j == 0 {
			i ++
			if i > row || j > col {
				break
			}
			res = append(res, matrix[i][j])
			addr = -1
			addc = 1
		}else if j == col {
			i ++
			if i > row || j > col {
				break
			}
			res = append(res, matrix[i][j])
			addr = 1
			addc = -1
		}
		i += addr
		j += addc
		//fmt.Println(res)
	}
	return res
}
