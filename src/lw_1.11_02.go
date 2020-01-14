package main

import "fmt"

func main() {
	mat := [][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	fmt.Println(matrixBlockSum(mat, 2))
}

func matrixBlockSum(mat [][]int, K int) [][]int {
	lena := len(mat)
	if lena == 0 {
		return nil
	}
	lenb := len(mat[0])
	var res [][]int
	for i := 0; i < lena; i ++ {
		tmp := make([]int, lenb)
		res = append(res, tmp)
	}
	for i := 0; i < lena; i ++ {
		for j := 0; j < lenb; j ++ {
			res[i][j] = cal(mat, i, j, K, lena, lenb)
		}
	}
	return res
}

func cal(mat [][]int, i, j, k, lena, lenb int) int {
	res := 0
	for m := i - k; m <= i + k; m ++ {
		if m >= lena {
			break
		}
		if m < 0 {
			continue
		}
		for n := j - k; n <= j + k; n ++ {
			if n >= lenb {
				break
			}
			if n < 0 {
				continue
			}
			res = res + mat[m][n]
		}
	}
	return res
}
