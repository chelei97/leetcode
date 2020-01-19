package main

import (
	"fmt"
)

func main() {
	nums := []int{-1, 3, 2, 0}
	fmt.Println(find132pattern(nums))
}

func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	var stack []int
	min := make([]int, len(nums))
	min[0] = nums[0]
	for i := 1; i < len(nums); i ++ {
		if nums[i] < min[i - 1] {
			min[i] = nums[i]
		}else {
			min[i] = min[i - 1]
		}
	}
	for i := len(nums) - 1; i > 0; i -- {
		for len(stack) != 0 && stack[len(stack) - 1] <= min[i]{
			stack = stack[ : len(stack) - 1]
		}
		if len(stack) != 0 && stack[len(stack) - 1] < nums[i] {
			return true
		}
		stack = append(stack, nums[i])
	}
	return false
}
