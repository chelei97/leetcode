package main

import "fmt"

/**
未出现的最小正整数
 */
func main() {
	nums := []int{1,1}
	fmt.Println(firstMissingPositive(nums))
}
func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i ++{
		for nums[i] > 0 && nums[i] < len(nums) + 1 && nums[i] != i + 1 && nums[nums[i] - 1] != nums[i]{
			temp := nums[i]
			nums[i] = nums[temp - 1]
			nums[temp - 1] = temp
		}
	}
	for i := 0; i < len(nums); i ++ {
		if nums[i] != i + 1 {
			return i + 1
		}
	}
	return len(nums) + 1
}
