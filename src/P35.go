package main

import "fmt"

func main()  {
	var nums = [] int {1,3,5,6}
	var target int = 5
	fmt.Print(searchInsert(nums, target))
}

func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] > target {
			return i - 1
		}
	}
	return 0
}