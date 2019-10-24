package main

import "fmt"

func main(){
	nums := []int{1,3}
	target := 0
	fmt.Println(search(nums,target))
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if target == nums[0]{
			return 0
		}else {
			return -1
		}
	}
	if target >= nums[0]{
		for i := 0; ; i ++ {
			if nums[i] == target {
				return i
			}
			if nums[i] > target  {
				return -1
			}
			if i == len(nums) - 1 || nums[i + 1] < nums[i] {
				return -1
			}
		}
	}else {
		for i := len(nums) - 1; ; i -- {
			if nums[i] == target{
				return i
			}
			if nums[i] < target  {
				return -1
			}
			if i == 0 || nums[i - 1] > nums[i]{
				return -1
			}
		}
	}
	return -1
}
