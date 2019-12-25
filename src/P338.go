package main

import "fmt"

func main() {
	num := 4
	fmt.Println(countBits(num))
}

func countBits(num int) []int {
	result := make([]int, num + 1)
	result[0] = 0
	for i := 1; i < len(result); i ++ {
		if i % 2 == 0 {
			result[i] = result[i / 2]
		}else {
			result[i] = result[i - 1] + 1
		}
	}
	return result
}
