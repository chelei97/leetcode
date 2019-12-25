package main

import (
	"fmt"
)

func main() {
	nums := []int{10,9,2,5,3,7,101,18}
	fmt.Println(lengthOfLIS(nums))
}

func lengthOfLIS(nums []int) int {
	record := make([]int, len(nums))
	result := 0
	for i := 0; i < len(nums); i ++ {
		record[i] = 1
	}
	for i := 1; i < len(nums); i ++ {
		for j := 0; j < i; j ++ {
			if nums[i] > nums[j] {
				record[i] = max1(record[j] + 1, record[i])
			}
		}
	}
	for i := 0; i < len(nums); i ++ {
		result = max1(result, record[i])
	}
	return result
}

func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}


