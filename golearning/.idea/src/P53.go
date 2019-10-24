package main

import "fmt"

/**
最大子序和
 */
func main() {
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray(nums))
}


func maxSubArray(nums []int) int {
	result := nums[0]
	max := result
	for i := 1; i < len(nums); i ++ {
		result = result + nums[i]
		//fmt.Printf("result = %v\n", result)
		if result < nums[i] {
			result = nums[i]
		}
		if max < result {
			max = result
		}
	}
	return max
}