package main

import "fmt"

/**
打家劫舍
 */
func main(){
	nums := []int{2,7,9,3,1}
	fmt.Println(rob(nums))
}

func rob(nums []int) int {
	first, second := nums[0], 0
	if nums[1] > nums[0] {
		second = nums[1]
	}else {
		second = nums[0]
	}
	for i := 2; i < len(nums); i ++ {
		if first + nums[i] > second {
			temp := first
			first = second
			second = temp + nums[i]
		}else {
			first = second
		}
	}
	return second
}

