package main

import (
	"fmt"
	"sort"
)

func main() {
	candidates := []int{1,1,1,3,3,5}
	fmt.Println(combinationSum2(candidates, 8))
}

func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}
	sort.Ints(candidates)
	res := make([][]int, 0)
	backtrack2(candidates, target, 0, &res, []int{})
	return res
}

func backtrack2(candidates []int, target, index int, res *[][]int, cur []int) {
	if target == 0 {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}
	for i := index; i < len(candidates); i ++ {
		if target - candidates[i] < 0 {
			break
		}
		if i > index && candidates[i] == candidates[i - 1] {
			continue
		}
		cur = append(cur, candidates[i])
		backtrack2(candidates, target - candidates[i], i + 1, res, cur)
		cur = cur[ : len(cur) - 1]
	}
}