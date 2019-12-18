package main

import (
	"fmt"
	"sort"
)

/**
三数之和
 */
func main(){
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; i < len(nums) - 2; {
		for j := i + 1; j < len(nums) - 1; {
			target := 0 - nums[i]

		}
	}
	return result
}
