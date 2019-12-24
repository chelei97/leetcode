package main

import "fmt"

func main() {
	n := 12
	fmt.Println(numSquares(n))
}

func numSquares(n int) int {
	var result int
	for n != 0 {
		n = n - sqrt(n) * sqrt(n)
		fmt.Println(n)
		result ++
	}
	return result
}

func sqrt(n int) int {
	if n < 4 {
		return 1
	}
	result := n / 2
	for result * result > n {
		result --
	}
	return result
}
