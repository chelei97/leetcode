package main

import (
	"fmt"
	"sort"
)

func main() {
	candidates := []int{2,3,6,7}
	fmt.Println(combinationSum(candidates, 7))
}

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) < 1 {
		return [][]int{}
	}
	var result [][]int
	sort.Ints(candidates)
	helper1(candidates, target, []int{}, &result, 0)
	return result
}

func helper1(candidates []int, need int, cur []int, result *[][]int, index int) {
	if need < 0 {
		return
	}
	if need == 0 {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*result = append(*result, tmp)
	}
	for i := index; i < len(candidates); i ++ {
		helper1(candidates, need - candidates[i], append(cur, candidates[i]), result, i)
	}
}