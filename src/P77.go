package main

import "fmt"

func main() {
	fmt.Println(combine(4, 2))
}

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	helper3([]int{}, &res, 0, k, n)
	return res
}

func helper3(cur []int, res *[][]int, index, k, n int) {
	if len(cur) == k {
		tmp := make([]int, k)
		copy(tmp, cur)
		*res = append(*res, tmp)
	}
	for i := index + 1; i < n + 1; i ++ {
		cur = append(cur, i)
		helper3(cur, res, i, k, n)
		cur = cur[ : len(cur) - 1]
	}
}