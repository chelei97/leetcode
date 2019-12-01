package main

import (
	"fmt"
)

func main(){
	nums := []int{1,3}
	target := 3
	fmt.Println(searchRange(nums,target))
}

func searchRange(nums []int, target int) []int {
	left, right, begin := 0, len(nums) - 1, -1
	result := []int{-1, -1}
	for left <= right {
		mid := (left + right) / 2
		fmt.Println(mid)
		if nums[mid] == target {
			begin = mid
		}
		if nums[mid] < target {
			left = mid + 1
		}else {
			right = mid - 1
		}
	}
	fmt.Println(begin)
	if begin == -1 {
		return result
	}else {
		right = begin
		left = begin
		for right < len(nums) - 1 && nums[right + 1] == nums [begin] {
			right ++
		}
		for left > 0 && nums[left - 1] == nums [begin] {
			left --
		}
		result[0] = left
		result[1] = right
		return result
	}
	return result
}
