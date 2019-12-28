package main

import "fmt"

func main() {
	nums := []int{2,3,2}
	fmt.Println(rob2(nums))
}

func rob2(nums []int) int {
	dp_0, dp_1 := 0, nums[0]
	for i := 1; i < len(nums) - 1; i ++ {
		tmp := dp_0
		dp_0 = max7(dp_0, dp_1)
		dp_1 = tmp + nums[i]
	}
	res := max7(dp_0, dp_1)
	dp_0, dp_1 = 0, nums[1]
	for i := 2; i < len(nums); i ++ {
		tmp := dp_0
		dp_0 = max7(dp_0, dp_1)
		dp_1 = tmp + nums[i]
	}
	res = max7(res, max7(dp_1, dp_0))
	return res
}

func max7(a, b int) int {
	if a > b {
		return a
	}
	return b
}
