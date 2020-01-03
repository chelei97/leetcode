package main

import "fmt"

func main() {
	nums := []int{4,5,6,7,0,1,2}
	fmt.Println(findMin(nums))
}

func findMin(nums []int) int {
	left, right := 0, len(nums) - 1
	if nums[len(nums) - 1] > nums[0] {
		return nums[0]
	}
	for left, right = 0, len(nums) - 1; left < right; {
		mid := (left + right) / 2
		if nums[mid + 1] < nums[mid] {
			return nums[mid + 1]
		}
		if nums[mid] < nums[mid - 1] {
			return nums[mid]
		}
		if nums[mid] < nums[0] {
			right = mid - 1
		}else {
			left = mid + 1
		}
	}
	return nums[left]
}