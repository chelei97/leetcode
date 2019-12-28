package main

import "fmt"

func main() {
	candidates := []int{10,1,2,7,6,1,5}
	fmt.Println(combinationSum2(candidates, 8))
}

func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
}