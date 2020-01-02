package main

import "fmt"

func main() {
	nums := []int{-2,3,-4}
	fmt.Println(maxProduct(nums))
}

func maxProduct(nums []int) int {
	dp_max, dp_min := nums[0], nums[0]
	res := nums[0]
	for i := 1; i < len(nums); i ++ {
		tmp := dp_max
		dp_max = max5(nums[i], dp_min * nums[i], dp_max * nums[i])
		dp_min = min2(nums[i], dp_min * nums[i], tmp * nums[i])
		if res < dp_max {
			res = dp_max
		}
	}
	return res
}

func min2(a... int) int {
	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}

func max5(a... int) int {
	max := a[0]
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	return max
}

