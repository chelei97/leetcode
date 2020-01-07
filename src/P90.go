package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1,2,2}
	fmt.Println(subsetsWithDup(nums))
}

func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	helper5(nums, []int{}, 0, &res)
	return res
}

func helper5(nums, cur []int, index int, res *[][]int) {
	tmp := make([]int, len(cur))
	copy(tmp, cur)
	*res = append(*res, tmp)
	for i := index; i < len(nums); i ++ {
		if i > index && nums[i] == nums[i - 1] {
			continue
		}
		cur = append(cur, nums[i])
		helper5(nums, cur, i + 1, res)
		cur = cur[ : len(cur) - 1]
	}
}
