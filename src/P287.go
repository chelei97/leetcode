package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1,3,4,2,2}
	fmt.Println(findDuplicate(nums))
}

func findDuplicate(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i ++ {
		if nums[i] == nums[i + 1]{
			return nums[i]
		}
	}
	return 0
}