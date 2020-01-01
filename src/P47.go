package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{0,1,0,0,9}
	res := permuteUnique(nums)
	fmt.Println(res)
	fmt.Println(len(res))
}

var used []bool

func permuteUnique(nums []int) [][]int {
	used =make([]bool, len(nums))
	sort.Ints(nums)
	res := make([][]int, 0)
	helper2(nums, []int{}, 0, &res)
	return res
}

func helper2(nums, cur []int, index int, res *[][]int) {
	if index == len(nums) {
		tmp := make([]int, index)
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}
	for i := 0; i < len(nums); i ++ {
		if !used[i] {
			if i > 0 && nums[i] == nums[i - 1] && !used[i - 1] {
				continue
			}
			used[i] = true
			cur = append(cur, nums[i])
			helper2(nums, cur, index + 1, res)
			cur = cur[ : len(cur) - 1]
			used[i] = false
		}
	}
}
