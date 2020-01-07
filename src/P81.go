package main

import "fmt"

func main() {
	nums := []int{2,5,6,0,0,1,2}
	fmt.Println(search1(nums, 0))
}

func search1(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	for left, right := 0, len(nums) - 1; left < right; {
		mid := left + (right - left) / 2
		if nums[mid] == target {
			return true
		}
		if nums[left] == nums[mid] {
			left ++
			continue
		}
		if nums[left] < nums[mid] {
			if nums[mid] > target && nums[left] <= target {
				right = mid - 1
			} else {  //否则，去后半部分找
				left = mid + 1
			}
		}else {
			if nums[mid] < target && nums[right] >= target {
				left = mid + 1
			} else {  //否则，去后半部分找
				right = mid - 1
			}
		}
	}
	return false
}