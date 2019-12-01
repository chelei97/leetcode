package main

import "fmt"

/**
移动0
 */
func main() {
	nums := []int{0,1,0,3,12}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int)  {
	count := 0
	for i := 0; i < len(nums); i ++ {
		if nums[i] == 0 {
			count ++
		}else {
				nums[i - count] = nums[i]
		}
	}
	for i := 0; i < count; i ++ {
		nums[len(nums) - 1 - i] = 0
	}
}
