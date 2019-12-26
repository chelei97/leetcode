package main

import "fmt"

func main() {
	nums := []int{2,4,6,8}
	fmt.Println(productExceptSelf(nums))
}

func productExceptSelf(nums []int) []int {
	record := make([]int, len(nums))
	k := 1
	for i := 0; i < len(nums); i ++ {
		record[i] = k
		k = k * nums[i]
	}
	k = 1
	for i := len(nums) - 1; i >= 0; i -- {
		record[i] = k * record[i]
		k = k * nums[i]
	}
	return record
}