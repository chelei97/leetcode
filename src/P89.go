package main

import "fmt"

func main() {
	fmt.Println(grayCode(2))
}

func grayCode(n int) []int {
	res := []int{0}
	head := 1
	for i := 0; i < n; i ++ {
		length := len(res)
		tmp := make([]int, length)
		for j := length - 1; j >= 0; j -- {
			tmp[length - 1 - j] = head + res[j]
		}
		head = head << 1
		res = append(res, tmp...)
	}
	return res
}
